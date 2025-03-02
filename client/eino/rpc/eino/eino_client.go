package eino

import (
	"context"
	eino "github.com/czczcz831/tiktok-mall/client/eino/kitex_gen/eino"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/czczcz831/tiktok-mall/client/eino/kitex_gen/eino/einoservice"
)

type RPCClient interface {
	KitexClient() einoservice.Client
	Service() string
	QueryUserOrders(ctx context.Context, req *eino.QueryUserOrdersReq, callOptions ...callopt.Option) (r *eino.QueryUserOrdersResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := einoservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient einoservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() einoservice.Client {
	return c.kitexClient
}

func (c *clientImpl) QueryUserOrders(ctx context.Context, req *eino.QueryUserOrdersReq, callOptions ...callopt.Option) (r *eino.QueryUserOrdersResp, err error) {
	return c.kitexClient.QueryUserOrders(ctx, req, callOptions...)
}
