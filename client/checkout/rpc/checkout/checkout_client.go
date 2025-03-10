package checkout

import (
	"context"
	checkout "github.com/czczcz831/tiktok-mall/client/checkout/kitex_gen/checkout"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/czczcz831/tiktok-mall/client/checkout/kitex_gen/checkout/checkoutservice"
)

type RPCClient interface {
	KitexClient() checkoutservice.Client
	Service() string
	CreateAddress(ctx context.Context, req *checkout.CreateAddressReq, callOptions ...callopt.Option) (r *checkout.CreateAddressResp, err error)
	UpdateAddress(ctx context.Context, req *checkout.UpdateAddressReq, callOptions ...callopt.Option) (r *checkout.UpdateAddressResp, err error)
	DeleteAddress(ctx context.Context, req *checkout.DeleteAddressReq, callOptions ...callopt.Option) (r *checkout.DeleteAddressResp, err error)
	GetAddress(ctx context.Context, req *checkout.GetAddressReq, callOptions ...callopt.Option) (r *checkout.GetAddressResp, err error)
	Checkout(ctx context.Context, req *checkout.CheckoutReq, callOptions ...callopt.Option) (r *checkout.CheckoutResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := checkoutservice.NewClient(dstService, opts...)
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
	kitexClient checkoutservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() checkoutservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateAddress(ctx context.Context, req *checkout.CreateAddressReq, callOptions ...callopt.Option) (r *checkout.CreateAddressResp, err error) {
	return c.kitexClient.CreateAddress(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateAddress(ctx context.Context, req *checkout.UpdateAddressReq, callOptions ...callopt.Option) (r *checkout.UpdateAddressResp, err error) {
	return c.kitexClient.UpdateAddress(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteAddress(ctx context.Context, req *checkout.DeleteAddressReq, callOptions ...callopt.Option) (r *checkout.DeleteAddressResp, err error) {
	return c.kitexClient.DeleteAddress(ctx, req, callOptions...)
}

func (c *clientImpl) GetAddress(ctx context.Context, req *checkout.GetAddressReq, callOptions ...callopt.Option) (r *checkout.GetAddressResp, err error) {
	return c.kitexClient.GetAddress(ctx, req, callOptions...)
}

func (c *clientImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq, callOptions ...callopt.Option) (r *checkout.CheckoutResp, err error) {
	return c.kitexClient.Checkout(ctx, req, callOptions...)
}
