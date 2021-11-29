package main

import (
	"errors"
	goflag "flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/cheggaaa/pb/v3"
	log "github.com/golang/glog"
	"github.com/spf13/cobra"

	instagramclient "github.com/footprintai/shrimping/components/grpc/client/instagram"
	_ "github.com/footprintai/shrimping/components/grpc/flags"
	protopb "github.com/footprintai/shrimping/components/grpc/proto/pb"
	"github.com/footprintai/shrimping/components/grpc/version"
)

var (
	serverAddr   string
	timeout      int
	apiKey       string
	outputfolder string
	cachePolicy  string
	// Shritagram = Shrimping + instagram

	rootCmd = &cobra.Command{
		Use:   "Shritagram",
		Short: "A EC data platform client",
		Long:  `Shritagram is a client program to access EC data platform with gRPC method`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// For cobra + glog flags. Available to all subcommands.
			goflag.Parse()
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "show server/client version",
		RunE: func(cmd *cobra.Command, args []string) error {
			run := mustNewRunCmd(serverAddr, apiKey, cachePolicy)
			if err := run.doVersion(); err != nil {
				return err
			}
			if err := run.Close(); err != nil {
				return err
			}
			return nil
		},
	}

	profileCmd = &cobra.Command{
		Use:   "profile",
		Short: "get instagram profile with username",
		RunE: func(cmd *cobra.Command, args []string) error {
			run := mustNewRunCmd(serverAddr, apiKey, cachePolicy)
			if err := run.getProfile(args, outputfolder); err != nil {
				return err
			}
			if err := run.Close(); err != nil {
				return err
			}
			return nil
		},
	}

	postsCmd = &cobra.Command{
		Use:   "posts",
		Short: "get instagram posts with username",
		RunE: func(cmd *cobra.Command, args []string) error {
			run := mustNewRunCmd(serverAddr, apiKey, cachePolicy)
			if err := run.getPosts(args, outputfolder); err != nil {
				return err
			}
			if err := run.Close(); err != nil {
				return err
			}
			return nil
		},
	}

	topsearchCmd = &cobra.Command{
		Use:   "topsearch",
		Short: "search instagram with keyword or hashtag",
		RunE: func(cmd *cobra.Command, args []string) error {
			run := mustNewRunCmd(serverAddr, apiKey, cachePolicy)
			if err := run.topSearch(args, outputfolder); err != nil {
				return err
			}
			if err := run.Close(); err != nil {
				return err
			}
			return nil
		},
	}
)

func mustNewRunCmd(serverAddr, token, cachePolicy string) *runCmd {
	cmd, err := newRunCmd(serverAddr, token, cachePolicy)
	if err != nil {
		panic(err)
	}
	return cmd
}

func newRunCmd(serverAddr, apiKey, cachePolicy string) (*runCmd, error) {
	if serverAddr == "" {
		return nil, errors.New("invalid server address, use --svr_address to specify")
	}
	p := newProgressBar()
	p.Start()

	clientService, err := instagramclient.NewClient(
		serverAddr,
		apiKey,
		time.Duration(timeoutInSec())*time.Second,
		cachePolicy,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect server :%s", serverAddr)
	}
	r := &runCmd{
		clientService: clientService,
		progressbar:   p,
		startTime:     time.Now(),
	}
	return r, nil
}

type runCmd struct {
	clientService *instagramclient.Client
	progressbar   *ProgressBar
	startTime     time.Time
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

func (r *runCmd) getProfile(usernames []string, outputfolder string) error {
	out, err := r.clientService.Profile(usernames)
	if err != nil {
		return err
	}
	return writeInstagramObjectToFolder(out, outputfolder)
}

func (r *runCmd) getPosts(usernames []string, outputfolder string) error {
	out, err := r.clientService.Posts(usernames)
	if err != nil {
		return err
	}
	return writeInstagramObjectToFolder(out, outputfolder)
}

func writeInstagramObjectToFolder(outs []*protopb.RawInstgramObject, folder string) error {
	for _, obj := range outs {
		folderPath := filepath.Join(folder, obj.Username)
		os.MkdirAll(folderPath, os.ModePerm)
		if obj.RawProfile != nil {
			if err := writeFile(filepath.Join(folderPath, "profile.html"), obj.RawProfile.RawBytes); err != nil {
				log.Errorf("failed to write file, err:%+v\n", err)
				return err
			}
		}
		for _, post := range obj.RawPosts {
			if post != nil {
				if err := writeFile(filepath.Join(folderPath, post.Shortcode), post.RawBytes); err != nil {
					log.Errorf("failed to write file, err:%+v\n", err)
					return err
				}
			}
		}
	}
	return nil
}

func writeFile(filepath string, content []byte) error {
	log.Infof("write file:%s, %d bytes\n", filepath, len(content))
	return ioutil.WriteFile(filepath, content, os.ModePerm)
}

func (r *runCmd) topSearch(keywords []string, outputfolder string) error {
	out, err := r.clientService.TopSearch(keywords)
	if err != nil {
		return err
	}
	return writeTopSearchObject(out, outputfolder)
}

func writeTopSearchObject(outs []*protopb.RawInstgramTopSearchObject, folder string) error {
	for _, out := range outs {
		if err := writeFile(filepath.Join(folder, out.Keyword), out.RawBytes); err != nil {
			log.Errorf("failed to write file, err:%+v\n", err)
			return err
		}
	}
	return nil
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
	rootCmd.AddCommand(profileCmd)
	rootCmd.AddCommand(postsCmd)
	rootCmd.AddCommand(topsearchCmd)

	rootCmd.PersistentFlags().StringVar(&serverAddr, "svr_address", "172.17.0.1:50050", "data platform server address (default is 172.17.0.1:50050)")
	rootCmd.PersistentFlags().IntVar(&timeout, "timeout", 180, "timeout duration in second. (default: 180)")
	rootCmd.PersistentFlags().StringVar(&apiKey, "apikey", "", "apikey for access apis")
	rootCmd.PersistentFlags().StringVar(&cachePolicy, "cachepolicy", "", "cache policy, (default: max-age:86400). use no-cache to ignore cache check")
	profileCmd.Flags().StringVar(&outputfolder, "output_dir", "./output", "output dir")
	postsCmd.Flags().StringVar(&outputfolder, "output_dir", "./output", "output dir")
}
