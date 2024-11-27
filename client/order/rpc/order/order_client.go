package order

import (
	"context"
	order "github.com/czczcz831/tiktok-mall/client/order/kitex_gen/order"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/czczcz831/tiktok-mall/client/order/kitex_gen/order/orderservice"
)

type RPCClient interface {
	KitexClient() orderservice.Client
	Service() string
	CreateOrder(ctx context.Context, req *order.CreateOrderReq, callOptions ...callopt.Option) (r *order.CreateOrderResp, err error)
	UpdateOrderAddress(ctx context.Context, req *order.UpdateOrderAddressReq, callOptions ...callopt.Option) (r *order.UpdateOrderAddressResp, err error)
	MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (r *order.MarkOrderPaidResp, err error)
	GetOrder(ctx context.Context, req *order.GetOrderReq, callOptions ...callopt.Option) (r *order.GetOrderResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := orderservice.NewClient(dstService, opts...)
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
	kitexClient orderservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() orderservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateOrder(ctx context.Context, req *order.CreateOrderReq, callOptions ...callopt.Option) (r *order.CreateOrderResp, err error) {
	return c.kitexClient.CreateOrder(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateOrderAddress(ctx context.Context, req *order.UpdateOrderAddressReq, callOptions ...callopt.Option) (r *order.UpdateOrderAddressResp, err error) {
	return c.kitexClient.UpdateOrderAddress(ctx, req, callOptions...)
}

func (c *clientImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (r *order.MarkOrderPaidResp, err error) {
	return c.kitexClient.MarkOrderPaid(ctx, req, callOptions...)
}

func (c *clientImpl) GetOrder(ctx context.Context, req *order.GetOrderReq, callOptions ...callopt.Option) (r *order.GetOrderResp, err error) {
	return c.kitexClient.GetOrder(ctx, req, callOptions...)
}
