// Code generated by hertz generator.

package checkout_service

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
	CreateAddress(context context.Context, req *api.CreateAddressReq, reqOpt ...config.RequestOption) (resp *api.CreateAddressResp, rawResponse *protocol.Response, err error)

	UpdateAddress(context context.Context, req *api.UpdateAddressReq, reqOpt ...config.RequestOption) (resp *api.UpdateAddressResp, rawResponse *protocol.Response, err error)

	DeleteAddress(context context.Context, req *api.DeleteAddressReq, reqOpt ...config.RequestOption) (resp *api.DeleteAddressResp, rawResponse *protocol.Response, err error)

	GetAddress(context context.Context, req *api.GetAddressReq, reqOpt ...config.RequestOption) (resp *api.GetAddressResp, rawResponse *protocol.Response, err error)

	Checkout(context context.Context, req *api.CheckoutReq, reqOpt ...config.RequestOption) (resp *api.CheckoutResp, rawResponse *protocol.Response, err error)
}

type CheckoutServiceClient struct {
	client *cli
}

func NewCheckoutServiceClient(hostUrl string, ops ...Option) (Client, error) {
	opts := getOptions(append(ops, withHostUrl(hostUrl))...)
	cli, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	return &CheckoutServiceClient{
		client: cli,
	}, nil
}

func (s *CheckoutServiceClient) CreateAddress(context context.Context, req *api.CreateAddressReq, reqOpt ...config.RequestOption) (resp *api.CreateAddressResp, rawResponse *protocol.Response, err error) {
	httpResp := &api.CreateAddressResp{}
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
		execute("POST", "/checkout/address")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *CheckoutServiceClient) UpdateAddress(context context.Context, req *api.UpdateAddressReq, reqOpt ...config.RequestOption) (resp *api.UpdateAddressResp, rawResponse *protocol.Response, err error) {
	httpResp := &api.UpdateAddressResp{}
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
		execute("PUT", "/checkout/address")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *CheckoutServiceClient) DeleteAddress(context context.Context, req *api.DeleteAddressReq, reqOpt ...config.RequestOption) (resp *api.DeleteAddressResp, rawResponse *protocol.Response, err error) {
	httpResp := &api.DeleteAddressResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{
			"uuid": req.GetUUID(),
		}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("DELETE", "/checkout/address/:uuid")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *CheckoutServiceClient) GetAddress(context context.Context, req *api.GetAddressReq, reqOpt ...config.RequestOption) (resp *api.GetAddressResp, rawResponse *protocol.Response, err error) {
	httpResp := &api.GetAddressResp{}
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
		execute("GET", "/checkout/address")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *CheckoutServiceClient) Checkout(context context.Context, req *api.CheckoutReq, reqOpt ...config.RequestOption) (resp *api.CheckoutResp, rawResponse *protocol.Response, err error) {
	httpResp := &api.CheckoutResp{}
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
		execute("POST", "/checkout")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

var defaultClient, _ = NewCheckoutServiceClient("")

func ConfigDefaultClient(ops ...Option) (err error) {
	defaultClient, err = NewCheckoutServiceClient("", ops...)
	return
}

func CreateAddress(context context.Context, req *api.CreateAddressReq, reqOpt ...config.RequestOption) (resp *api.CreateAddressResp, rawResponse *protocol.Response, err error) {
	return defaultClient.CreateAddress(context, req, reqOpt...)
}

func UpdateAddress(context context.Context, req *api.UpdateAddressReq, reqOpt ...config.RequestOption) (resp *api.UpdateAddressResp, rawResponse *protocol.Response, err error) {
	return defaultClient.UpdateAddress(context, req, reqOpt...)
}

func DeleteAddress(context context.Context, req *api.DeleteAddressReq, reqOpt ...config.RequestOption) (resp *api.DeleteAddressResp, rawResponse *protocol.Response, err error) {
	return defaultClient.DeleteAddress(context, req, reqOpt...)
}

func GetAddress(context context.Context, req *api.GetAddressReq, reqOpt ...config.RequestOption) (resp *api.GetAddressResp, rawResponse *protocol.Response, err error) {
	return defaultClient.GetAddress(context, req, reqOpt...)
}

func Checkout(context context.Context, req *api.CheckoutReq, reqOpt ...config.RequestOption) (resp *api.CheckoutResp, rawResponse *protocol.Response, err error) {
	return defaultClient.Checkout(context, req, reqOpt...)
}
