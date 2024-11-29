// Code generated by Kitex v0.9.1. DO NOT EDIT.

package cartservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	cart "github.com/czczcz831/tiktok-mall/app/cart/kitex_gen/cart"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddProductToCart(ctx context.Context, req *cart.AddProductToCartReq, callOptions ...callopt.Option) (r *cart.AddProductToCartResp, err error)
	ClearCart(ctx context.Context, req *cart.ClearCartReq, callOptions ...callopt.Option) (r *cart.ClearCartResp, err error)
	GetCart(ctx context.Context, req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error)
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
	return &kCartServiceClient{
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

type kCartServiceClient struct {
	*kClient
}

func (p *kCartServiceClient) AddProductToCart(ctx context.Context, req *cart.AddProductToCartReq, callOptions ...callopt.Option) (r *cart.AddProductToCartResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddProductToCart(ctx, req)
}

func (p *kCartServiceClient) ClearCart(ctx context.Context, req *cart.ClearCartReq, callOptions ...callopt.Option) (r *cart.ClearCartResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ClearCart(ctx, req)
}

func (p *kCartServiceClient) GetCart(ctx context.Context, req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCart(ctx, req)
}
