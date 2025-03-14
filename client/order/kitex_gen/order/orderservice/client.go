// Code generated by Kitex v0.9.1. DO NOT EDIT.

package orderservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	order "github.com/czczcz831/tiktok-mall/client/order/kitex_gen/order"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateOrder(ctx context.Context, req *order.CreateOrderReq, callOptions ...callopt.Option) (r *order.CreateOrderResp, err error)
	UpdateOrderAddress(ctx context.Context, req *order.UpdateOrderAddressReq, callOptions ...callopt.Option) (r *order.UpdateOrderAddressResp, err error)
	MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (r *order.MarkOrderPaidResp, err error)
	GetOrder(ctx context.Context, req *order.GetOrderReq, callOptions ...callopt.Option) (r *order.GetOrderResp, err error)
	GetUserOrders(ctx context.Context, req *order.GetUserOrdersReq, callOptions ...callopt.Option) (r *order.GetUserOrdersResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kOrderServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kOrderServiceClient struct {
	*kClient
}

func (p *kOrderServiceClient) CreateOrder(ctx context.Context, req *order.CreateOrderReq, callOptions ...callopt.Option) (r *order.CreateOrderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateOrder(ctx, req)
}

func (p *kOrderServiceClient) UpdateOrderAddress(ctx context.Context, req *order.UpdateOrderAddressReq, callOptions ...callopt.Option) (r *order.UpdateOrderAddressResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateOrderAddress(ctx, req)
}

func (p *kOrderServiceClient) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (r *order.MarkOrderPaidResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MarkOrderPaid(ctx, req)
}

func (p *kOrderServiceClient) GetOrder(ctx context.Context, req *order.GetOrderReq, callOptions ...callopt.Option) (r *order.GetOrderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetOrder(ctx, req)
}

func (p *kOrderServiceClient) GetUserOrders(ctx context.Context, req *order.GetUserOrdersReq, callOptions ...callopt.Option) (r *order.GetUserOrdersResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserOrders(ctx, req)
}
