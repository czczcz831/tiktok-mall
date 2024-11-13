// Code generated by Kitex v0.9.1. DO NOT EDIT.

package authservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	auth "github.com/czczcz831/tiktok-mall/client/auth/kitex_gen/auth"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error)
	RefeshTokenByRPC(ctx context.Context, req *auth.RefeshTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error)
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
	return &kAuthServiceClient{
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

type kAuthServiceClient struct {
	*kClient
}

func (p *kAuthServiceClient) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeliverTokenByRPC(ctx, req)
}

func (p *kAuthServiceClient) RefeshTokenByRPC(ctx context.Context, req *auth.RefeshTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RefeshTokenByRPC(ctx, req)
}