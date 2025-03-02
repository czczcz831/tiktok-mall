// Code generated by Kitex v0.9.1. DO NOT EDIT.

package einoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	eino "github.com/czczcz831/tiktok-mall/client/eino/kitex_gen/eino"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CallAssistantAgent(ctx context.Context, req *eino.CallAssistantAgentReq, callOptions ...callopt.Option) (r *eino.CallAssistantAgentResp, err error)
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
	return &kEinoServiceClient{
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

type kEinoServiceClient struct {
	*kClient
}

func (p *kEinoServiceClient) CallAssistantAgent(ctx context.Context, req *eino.CallAssistantAgentReq, callOptions ...callopt.Option) (r *eino.CallAssistantAgentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CallAssistantAgent(ctx, req)
}
