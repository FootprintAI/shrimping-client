package main

import (
	"errors"
	goflag "flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	//log "github.com/golang/glog"
	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/cobra"

	"github.com/footprintai/shrimping/components/grpc/client"
	shopeeclient "github.com/footprintai/shrimping/components/grpc/client/shopee"
	shopeetypes "github.com/footprintai/shrimping/components/grpc/types/shopee"
	"github.com/footprintai/shrimping/components/grpc/version"
)

var (
	cfg          string
	serverAddr   string
	format       string
	timeout      int
	limit        int
	dateCategory string
	email        string
	password     string

	rootCmd = &cobra.Command{
		Use:   "shrimping",
		Short: "A EC data platform client",
		Long:  `Shrimping is a client program to access EC data platform with gRPC method`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// For cobra + glog flags. Available to all subcommands.
			goflag.Parse()
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "show server/client version",
		RunE: func(cmd *cobra.Command, args []string) error {
			run := mustNewRunCmd(serverAddr, cfg)
			if err := run.doVersion(); err != nil {
				return err
			}
			if err := run.Close(); err != nil {
				return err
			}
			return nil
		},
	}

	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "login with email & password",
		Long:  "login with email & password, ex: ./shrimping login --email=<email> --password=<password>",
		RunE: func(cmd *cobra.Command, args []string) error {
			run := mustNewRunCmd(serverAddr, cfg)
			if err := run.doLogin(); err != nil {
				return err
			}
			if err := run.Close(); err != nil {
				return err
			}
			return nil
		},
	}

	findItemByIdCmd = &cobra.Command{
		Use:   "itemids",
		Short: "lookup items by item ids",
		Long:  "lookup items by item ids, ex: ./shrimping itemid 9519447478",
		RunE: func(cmd *cobra.Command, args []string) error {
			run := mustNewRunCmd(serverAddr, cfg)
			if err := run.doItemIds(args); err != nil {
				return err
			}
			if err := run.Close(); err != nil {
				return err
			}
			return nil
		},
	}
	findItemByURLCmd = &cobra.Command{
		Use:   "itemurls",
		Short: "lookup items by item urls",
		Long:  "lookup items by item urls, ex: ./shrimping itemurl https://shopee.tw/x-i.297780283.9519447478",
		RunE: func(cmd *cobra.Command, args []string) error {
			run := mustNewRunCmd(serverAddr, cfg)
			if err := run.doItemURLs(args); err != nil {
				return err
			}
			if err := run.Close(); err != nil {
				return err
			}
			return nil
		},
	}
	findItemsByCategoryNameCmd = &cobra.Command{
		Use:   "category",
		Short: "lookup items by category name",
		Long:  "lookup items by category name, ex: ./shrimping category 男運動休閒鞋 --limit 1000 --sortedby day1",
		RunE: func(cmd *cobra.Command, args []string) error {
			run := mustNewRunCmd(serverAddr, cfg)
			if err := run.doCategory(args); err != nil {
				return err
			}
			if err := run.Close(); err != nil {
				return err
			}
			return nil
		},
	}
)

func mustNewRunCmd(serverAddr, cfg string) *runCmd {
	cmd, err := newRunCmd(serverAddr, cfg)
	if err != nil {
		panic(err)
	}
	return cmd
}

func newRunCmd(serverAddr, cfg string) (*runCmd, error) {
	if serverAddr == "" {
		return nil, errors.New("invalid server address, use --svr_address to specify")
	}
	p := newProgressBar()
	p.Start()

	clientService, err := shopeeclient.NewClient(
		serverAddr,
		cfg,
		time.Duration(timeoutInSec())*time.Second,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect server :%s", serverAddr)
	}
	r := &runCmd{
		clientService: clientService,
		progressbar:   p,
		startTime:     time.Now(),
		stdout:        client.NewFormatWriter(os.Stdout, client.MustParseFormat(format)),
	}
	return r, nil
}

type runCmd struct {
	clientService *shopeeclient.Client
	progressbar   *ProgressBar
	startTime     time.Time
	stdout        *client.FormatWriter
}

const defaultTimeoutInSecond = 180

func timeoutInSec() int {
	if timeout < 0 {
		return defaultTimeoutInSecond
	}
	return timeout
}

func newProgressBar() *ProgressBar {
	return &ProgressBar{
		closeCh: make(chan struct{}),
	}
}

type ProgressBar struct {
	closeCh chan struct{}
}

func (p *ProgressBar) Start() {
	go p.reportLoop()
}

func (p *ProgressBar) Stop() {
	close(p.closeCh)
}

func (p *ProgressBar) reportLoop() {
	var tmpl pb.ProgressBarTemplate = `{{string . "prefix"}}{{counters .  }} {{bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) }} {{string . "suffix"}}`
	bar := tmpl.Start(timeoutInSec())
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			bar.Increment()
		case <-p.closeCh:
			bar.Finish()
			return
		}
	}
}

func (r *runCmd) doLogin() error {
	return r.clientService.Login(email, password)
}

func (r *runCmd) doItemIds(args []string) error {
	var ids []uint64
	for _, arg := range args {
		val, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}
		ids = append(ids, uint64(val))
	}
	fmt.Printf("lookup items: %+v\n", ids)
	items, err := r.clientService.LookupItemByIDs(ids)
	if err != nil {
		return err
	}
	return r.generateOutput(items)
}

func (r *runCmd) generateOutput(items []shopeeclient.ClientItem) error {
	headers := shopeeclient.ClientItem{}.ToCSVHeader()
	values := [][]string{}
	for _, item := range items {
		values = append(values, item.ToCSVValue())
	}

	return r.stdout.WriteAndClose(headers, values)
}

func (r *runCmd) doItemURLs(args []string) error {
	var ids []uint64
	re, _ := regexp.Compile("-i.([0-9]*).([0-9]*)")
	for _, arg := range args {
		allSubmatches := re.FindAllStringSubmatch(arg, -1)
		if len(allSubmatches[0]) != 3 {
			return fmt.Errorf("failed to parse item url:%s", arg)
		}
		val, _ := strconv.Atoi(allSubmatches[0][2])
		ids = append(ids, uint64(val))
	}
	fmt.Printf("lookup items: %+v\n", ids)
	items, err := r.clientService.LookupItemByIDs(ids)
	if err != nil {
		return err
	}
	return r.generateOutput(items)
}

func (r *runCmd) doCategory(args []string) error {
	dateCategory := shopeetypes.ParseDateCategory(dateCategory)
	if dateCategory == shopeetypes.UnknownDateCategory {
		return fmt.Errorf("invalid sortedby value:%s\n", dateCategory)
	}
	fmt.Printf("lookup categories: %+v\n", args)
	var items []shopeeclient.ClientItem
	for _, arg := range args {
		partItems, err := r.clientService.LookupItemBySales(dateCategory, int32(limit), nil, &arg)
		if err != nil {
			return err
		}
		items = append(items, partItems...)
	}
	return r.generateOutput(items)
}

func (r *runCmd) doVersion() error {
	cver := version.GetVersion()
	buildt := version.GetBuildTime()
	hashId := version.GetCommitHash()
	ver, compatible, err := r.clientService.Version()
	if err != nil {
		return err
	}
	fmt.Printf("client build time:%s, hash id:%s\n", buildt, hashId)
	fmt.Printf("client version:%s, server version:%s, compatibility:%v\n", cver, ver, compatible)
	if !compatible {
		fmt.Printf("this client is out-of-date, the result may not be guaranteed.\n")
		fmt.Printf("please go to https://github.com/footprintai/shrimping-client to download latest version\n")
	}
	return nil
}

func (r *runCmd) Close() error {
	r.progressbar.Stop()

	dur := time.Now().Sub(r.startTime)
	fmt.Printf("this cmd tooks %s\n", dur)

	return r.clientService.Close()
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(findItemByIdCmd)
	rootCmd.AddCommand(findItemByURLCmd)
	rootCmd.AddCommand(findItemsByCategoryNameCmd)

	rootCmd.PersistentFlags().StringVar(&cfg, "cfg", "./shrimpingcfg", "config file (default is ./shrimpingcfg)")
	rootCmd.PersistentFlags().StringVar(&serverAddr, "svr_address", "172.17.0.1:50050", "data platform server address (default is 172.17.0.1:50050)")
	rootCmd.PersistentFlags().StringVar(&format, "format", "table", "output format (default: table), possible value: csv / table")
	rootCmd.PersistentFlags().IntVar(&timeout, "timeout", 60, "timeout duration in second. (default: 60)")
	loginCmd.Flags().StringVar(&email, "email", "", "email for login")
	loginCmd.Flags().StringVar(&password, "password", "", "password for login")
	findItemsByCategoryNameCmd.Flags().IntVar(&limit, "limit", 100, "number of items maximum")
	findItemsByCategoryNameCmd.Flags().StringVar(&dateCategory, "sortedby", "day1", "sorted by day1 (sale yesterday) / day7 (last 7 days) / day30 (last 30 days)")
}
