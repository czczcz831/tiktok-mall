package cart

import (
	"context"
	cart "github.com/czczcz831/tiktok-mall/client/cart/kitex_gen/cart"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/czczcz831/tiktok-mall/client/cart/kitex_gen/cart/cartservice"
)

type RPCClient interface {
	KitexClient() cartservice.Client
	Service() string
	AddProductToCart(ctx context.Context, req *cart.AddProductToCartReq, callOptions ...callopt.Option) (r *cart.AddProductToCartResp, err error)
	ClearCart(ctx context.Context, req *cart.ClearCartReq, callOptions ...callopt.Option) (r *cart.ClearCartResp, err error)
	GetCart(ctx context.Context, req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := cartservice.NewClient(dstService, opts...)
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
	kitexClient cartservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() cartservice.Client {
	return c.kitexClient
}

func (c *clientImpl) AddProductToCart(ctx context.Context, req *cart.AddProductToCartReq, callOptions ...callopt.Option) (r *cart.AddProductToCartResp, err error) {
	return c.kitexClient.AddProductToCart(ctx, req, callOptions...)
}

func (c *clientImpl) ClearCart(ctx context.Context, req *cart.ClearCartReq, callOptions ...callopt.Option) (r *cart.ClearCartResp, err error) {
	return c.kitexClient.ClearCart(ctx, req, callOptions...)
}

func (c *clientImpl) GetCart(ctx context.Context, req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error) {
	return c.kitexClient.GetCart(ctx, req, callOptions...)
}
