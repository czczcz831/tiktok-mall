// Code generated by hertz generator.

package user_service

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
	Login(context context.Context, req *api.LoginReq, reqOpt ...config.RequestOption) (resp *api.LoginResp, rawResponse *protocol.Response, err error)

	RefreshToken(context context.Context, req *api.RefreshTokenReq, reqOpt ...config.RequestOption) (resp *api.LoginResp, rawResponse *protocol.Response, err error)

	Register(context context.Context, req *api.RegisterReq, reqOpt ...config.RequestOption) (resp *api.RegisterResp, rawResponse *protocol.Response, err error)
}

type UserServiceClient struct {
	client *cli
}

func NewUserServiceClient(hostUrl string, ops ...Option) (Client, error) {
	opts := getOptions(append(ops, withHostUrl(hostUrl))...)
	cli, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	return &UserServiceClient{
		client: cli,
	}, nil
}

func (s *UserServiceClient) Login(context context.Context, req *api.LoginReq, reqOpt ...config.RequestOption) (resp *api.LoginResp, rawResponse *protocol.Response, err error) {
	httpResp := &api.LoginResp{}
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
		execute("POST", "/user/login")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *UserServiceClient) RefreshToken(context context.Context, req *api.RefreshTokenReq, reqOpt ...config.RequestOption) (resp *api.LoginResp, rawResponse *protocol.Response, err error) {
	httpResp := &api.LoginResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{
			"Refresh-Token": req.GetRefreshToken(),
		}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/user/refresh_token")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *UserServiceClient) Register(context context.Context, req *api.RegisterReq, reqOpt ...config.RequestOption) (resp *api.RegisterResp, rawResponse *protocol.Response, err error) {
	httpResp := &api.RegisterResp{}
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
		execute("POST", "/user/register")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

var defaultClient, _ = NewUserServiceClient("")

func ConfigDefaultClient(ops ...Option) (err error) {
	defaultClient, err = NewUserServiceClient("", ops...)
	return
}

func Login(context context.Context, req *api.LoginReq, reqOpt ...config.RequestOption) (resp *api.LoginResp, rawResponse *protocol.Response, err error) {
	return defaultClient.Login(context, req, reqOpt...)
}

func RefreshToken(context context.Context, req *api.RefreshTokenReq, reqOpt ...config.RequestOption) (resp *api.LoginResp, rawResponse *protocol.Response, err error) {
	return defaultClient.RefreshToken(context, req, reqOpt...)
}

func Register(context context.Context, req *api.RegisterReq, reqOpt ...config.RequestOption) (resp *api.RegisterResp, rawResponse *protocol.Response, err error) {
	return defaultClient.Register(context, req, reqOpt...)
}
