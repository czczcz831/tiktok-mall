// Code generated by Kitex v0.9.1. DO NOT EDIT.

package einoservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	eino "github.com/czczcz831/tiktok-mall/app/eino/kitex_gen/eino"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"QueryUserOrders": kitex.NewMethodInfo(
		queryUserOrdersHandler,
		newEinoServiceQueryUserOrdersArgs,
		newEinoServiceQueryUserOrdersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	einoServiceServiceInfo                = NewServiceInfo()
	einoServiceServiceInfoForClient       = NewServiceInfoForClient()
	einoServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return einoServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return einoServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return einoServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "EinoService"
	handlerType := (*eino.EinoService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "eino",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func queryUserOrdersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*eino.EinoServiceQueryUserOrdersArgs)
	realResult := result.(*eino.EinoServiceQueryUserOrdersResult)
	success, err := handler.(eino.EinoService).QueryUserOrders(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newEinoServiceQueryUserOrdersArgs() interface{} {
	return eino.NewEinoServiceQueryUserOrdersArgs()
}

func newEinoServiceQueryUserOrdersResult() interface{} {
	return eino.NewEinoServiceQueryUserOrdersResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) QueryUserOrders(ctx context.Context, req *eino.QueryUserOrdersReq) (r *eino.QueryUserOrdersResp, err error) {
	var _args eino.EinoServiceQueryUserOrdersArgs
	_args.Req = req
	var _result eino.EinoServiceQueryUserOrdersResult
	if err = p.c.Call(ctx, "QueryUserOrders", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
