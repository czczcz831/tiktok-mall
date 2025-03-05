package payment

import (
	"context"
	payment "github.com/czczcz831/tiktok-mall/client/payment/kitex_gen/payment"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/czczcz831/tiktok-mall/client/payment/kitex_gen/payment/paymentservice"
)

type RPCClient interface {
	KitexClient() paymentservice.Client
	Service() string
	Charge(ctx context.Context, req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error)
	CancelCharge(ctx context.Context, req *payment.CancelChargeReq, callOptions ...callopt.Option) (r *payment.CancelChargeResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := paymentservice.NewClient(dstService, opts...)
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
	kitexClient paymentservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() paymentservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Charge(ctx context.Context, req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error) {
	return c.kitexClient.Charge(ctx, req, callOptions...)
}

func (c *clientImpl) CancelCharge(ctx context.Context, req *payment.CancelChargeReq, callOptions ...callopt.Option) (r *payment.CancelChargeResp, err error) {
	return c.kitexClient.CancelCharge(ctx, req, callOptions...)
}
