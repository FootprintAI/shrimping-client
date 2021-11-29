package shopeeclient

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	//log "github.com/golang/glog"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/metadata"

	"github.com/footprintai/shrimping/components/grpc/client"
	errors "github.com/footprintai/shrimping/components/grpc/errors"
	pb "github.com/footprintai/shrimping/components/grpc/proto/pb"
	"github.com/footprintai/shrimping/components/grpc/types"
	"github.com/footprintai/shrimping/components/grpc/version"
)

func NewClient(serverAddr string, localConfigureFilePath string, timeout time.Duration) (*Client, error) {
	clientConn, err := client.NewGrpcClient(serverAddr)
	if err != nil {
		return nil, err
	}
	lfc := client.NewLocalFileConfigure(localConfigureFilePath)
	lfc.Load()

	return &Client{
		pbClient:  pb.NewShrimpingClient(clientConn),
		verClient: pb.NewVersioningClient(clientConn),
		conn:      clientConn,
		lfc:       lfc,
		timeout:   timeout,
	}, nil
}

type Client struct {
	pbClient  pb.ShrimpingClient
	verClient pb.VersioningClient
	lfc       *client.LocalFileConfigure
	conn      *client.GrpcClient
	timeout   time.Duration
}

func (c *Client) Login(email, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.pbClient.Login(ctx, &pb.LoginRequest{
		ServiceType: "local",
		Email:       email,
		Password:    password,
	})
	if err != nil {
		return errors.ParseError(err)
	}
	fmt.Printf("login succeeded!\n")

	// update to configure
	c.lfc.Data.Token = resp.Token
	c.lfc.Data.LastLoginTime = time.Now()
	return nil
}

func (c *Client) LookupItemByIDs(itemIds []uint64) ([]ClientItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.pbClient.LookupItemByIDs(c.ctxWithToken(ctx), &pb.LookupItemByIDsRequest{
		Ids: itemIds,
	})
	if err != nil {
		return nil, errors.ParseError(err)
	}
	return convertRawResponse2Items(resp)
}

func (c *Client) ctxWithToken(ctx context.Context) context.Context {
	md := metadata.Pairs("authorization", fmt.Sprintf("bearer %v", c.lfc.Data.Token))
	nCtx := metautils.NiceMD(md).ToOutgoing(ctx)
	return nCtx
}

func convertRawResponse2Items(resp *pb.RawResponse) ([]ClientItem, error) {
	var items []ClientItem
	for _, rawObject := range resp.RawObjects {
		if len(rawObject.Item) != 0 {
			itm := ClientItem{}
			if err := json.Unmarshal(rawObject.Item, &itm); err != nil {
				return nil, err
			}
			items = append(items, itm)
		}
	}
	return items, nil
}

type ClientItem struct {
	ItemId    uint64 `json:"itemId"`
	ShopId    uint64 `json:"shopId"`
	CatId1    uint64 `json:"catId1"`
	CatId2    uint64 `json:"catId2"`
	CatId3    uint64 `json:"catId3"`
	CatName1  string `json:"catName1"`
	CatName2  string `json:"catName2"`
	CatName3  string `json:"catName3"`
	SoldDay1  int64  `json:"s1"`
	SoldDay7  int64  `json:"s7"`
	SoldDay30 int64  `json:"s30"`

	Details ItemDetails `json:"details"`
}

type ItemDetails struct {
	Snapshots []ItemSoldSnapshot `json:"s"`
}

type ItemSoldSnapshot struct {
	CrawledTs int64 `json:"cts"`
	Sold      int64 `json:"s"`
}

func (c ClientItem) ToCSVHeader() []string {
	return []string{
		"ItemID",
		"ShopID",
		"Category",
		"Sale D1",
		"Sale D7",
		"Sale D30",
		"Url",
		"LastCheckedSale",
		"LastCheckedTime",
	}
}

func (c ClientItem) ToCSVValue() []string {
	return []string{
		fmt.Sprintf("%d", c.ItemId),
		fmt.Sprintf("%d", c.ShopId),
		fmt.Sprintf("%s(%d),%s(%d),%s(%d)", c.CatName1, c.CatId1, c.CatName2, c.CatId2, c.CatName3, c.CatId3),
		fmt.Sprintf("%d", c.SoldDay1),
		fmt.Sprintf("%d", c.SoldDay7),
		fmt.Sprintf("%d", c.SoldDay30),
		fmt.Sprintf("https://shopee.tw/xx-i.%d.%d", c.ShopId, c.ItemId),
		fmt.Sprintf("%d", c.Details.Snapshots[len(c.Details.Snapshots)-1].Sold),
		time.Unix(c.Details.Snapshots[len(c.Details.Snapshots)-1].CrawledTs, 0).Format("20060102"),
	}
}

func (c *Client) LookupItemBySales(dayCategory types.DateCategory, limit int32, categoryId *int32, categoryName *string) ([]ClientItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.pbClient.LookupItemBySales(c.ctxWithToken(ctx), &pb.LookupItemBySalesRequest{
		DayCategory:  dayCategory.String(),
		CategoryId:   categoryId,
		CategoryName: categoryName,
		Limit:        &limit,
	})
	if err != nil {
		return nil, errors.ParseError(err)
	}
	return convertRawResponse2Items(resp)
}

func (c *Client) Version() (serverVersion string, isVersionCompatible bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ver := version.GetVersion()
	resp, err := c.verClient.Version(
		c.ctxWithToken(ctx),
		&pb.VersionRequest{
			ClientVersion: ver,
		},
	)
	if err != nil {
		return "", false, err
	}
	return resp.ServerVersion, resp.Compatible, nil
}

func (c *Client) Close() error {
	c.lfc.Save()
	return c.conn.Close()
}
