// Code generated by hertz generator.

package payment_service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type Client interface {
	Charge(context context.Context, req *api.ChargeReq, reqOpt ...config.RequestOption) (resp *api.ChargeResp, rawResponse *protocol.Response, err error)

	CancelCharge(context context.Context, req *api.CancelChargeReq, reqOpt ...config.RequestOption) (resp *api.CancelChargeResp, rawResponse *protocol.Response, err error)
}

type PaymentServiceClient struct {
	client *cli
}

func NewPaymentServiceClient(hostUrl string, ops ...Option) (Client, error) {
	opts := getOptions(append(ops, withHostUrl(hostUrl))...)
	cli, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	return &PaymentServiceClient{
		client: cli,
	}, nil
}

func (s *PaymentServiceClient) Charge(context context.Context, req *api.ChargeReq, reqOpt ...config.RequestOption) (resp *api.ChargeResp, rawResponse *protocol.Response, err error) {
	httpResp := &api.ChargeResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/payment/charge")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *PaymentServiceClient) CancelCharge(context context.Context, req *api.CancelChargeReq, reqOpt ...config.RequestOption) (resp *api.CancelChargeResp, rawResponse *protocol.Response, err error) {
	httpResp := &api.CancelChargeResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/payment/cancel")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

var defaultClient, _ = NewPaymentServiceClient("")

func ConfigDefaultClient(ops ...Option) (err error) {
	defaultClient, err = NewPaymentServiceClient("", ops...)
	return
}

func Charge(context context.Context, req *api.ChargeReq, reqOpt ...config.RequestOption) (resp *api.ChargeResp, rawResponse *protocol.Response, err error) {
	return defaultClient.Charge(context, req, reqOpt...)
}

func CancelCharge(context context.Context, req *api.CancelChargeReq, reqOpt ...config.RequestOption) (resp *api.CancelChargeResp, rawResponse *protocol.Response, err error) {
	return defaultClient.CancelCharge(context, req, reqOpt...)
}
