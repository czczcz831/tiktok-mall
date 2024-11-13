package auth

import (
	"context"
	auth "github.com/czczcz831/tiktok-mall/client/auth/kitex_gen/auth"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/czczcz831/tiktok-mall/client/auth/kitex_gen/auth/authservice"
)

type RPCClient interface {
	KitexClient() authservice.Client
	Service() string
	DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error)
	RefeshTokenByRPC(ctx context.Context, req *auth.RefeshTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := authservice.NewClient(dstService, opts...)
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
	kitexClient authservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() authservice.Client {
	return c.kitexClient
}

func (c *clientImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error) {
	return c.kitexClient.DeliverTokenByRPC(ctx, req, callOptions...)
}

func (c *clientImpl) RefeshTokenByRPC(ctx context.Context, req *auth.RefeshTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error) {
	return c.kitexClient.RefeshTokenByRPC(ctx, req, callOptions...)
}
