// Code generated by Kitex v0.9.1. DO NOT EDIT.

package checkoutservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateAddress": kitex.NewMethodInfo(
		createAddressHandler,
		newCheckoutServiceCreateAddressArgs,
		newCheckoutServiceCreateAddressResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateAddress": kitex.NewMethodInfo(
		updateAddressHandler,
		newCheckoutServiceUpdateAddressArgs,
		newCheckoutServiceUpdateAddressResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteAddress": kitex.NewMethodInfo(
		deleteAddressHandler,
		newCheckoutServiceDeleteAddressArgs,
		newCheckoutServiceDeleteAddressResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetAddress": kitex.NewMethodInfo(
		getAddressHandler,
		newCheckoutServiceGetAddressArgs,
		newCheckoutServiceGetAddressResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Checkout": kitex.NewMethodInfo(
		checkoutHandler,
		newCheckoutServiceCheckoutArgs,
		newCheckoutServiceCheckoutResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	checkoutServiceServiceInfo                = NewServiceInfo()
	checkoutServiceServiceInfoForClient       = NewServiceInfoForClient()
	checkoutServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return checkoutServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return checkoutServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return checkoutServiceServiceInfoForClient
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
	serviceName := "CheckoutService"
	handlerType := (*checkout.CheckoutService)(nil)
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
		"PackageName": "checkout",
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

func createAddressHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*checkout.CheckoutServiceCreateAddressArgs)
	realResult := result.(*checkout.CheckoutServiceCreateAddressResult)
	success, err := handler.(checkout.CheckoutService).CreateAddress(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCheckoutServiceCreateAddressArgs() interface{} {
	return checkout.NewCheckoutServiceCreateAddressArgs()
}

func newCheckoutServiceCreateAddressResult() interface{} {
	return checkout.NewCheckoutServiceCreateAddressResult()
}

func updateAddressHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*checkout.CheckoutServiceUpdateAddressArgs)
	realResult := result.(*checkout.CheckoutServiceUpdateAddressResult)
	success, err := handler.(checkout.CheckoutService).UpdateAddress(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCheckoutServiceUpdateAddressArgs() interface{} {
	return checkout.NewCheckoutServiceUpdateAddressArgs()
}

func newCheckoutServiceUpdateAddressResult() interface{} {
	return checkout.NewCheckoutServiceUpdateAddressResult()
}

func deleteAddressHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*checkout.CheckoutServiceDeleteAddressArgs)
	realResult := result.(*checkout.CheckoutServiceDeleteAddressResult)
	success, err := handler.(checkout.CheckoutService).DeleteAddress(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCheckoutServiceDeleteAddressArgs() interface{} {
	return checkout.NewCheckoutServiceDeleteAddressArgs()
}

func newCheckoutServiceDeleteAddressResult() interface{} {
	return checkout.NewCheckoutServiceDeleteAddressResult()
}

func getAddressHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*checkout.CheckoutServiceGetAddressArgs)
	realResult := result.(*checkout.CheckoutServiceGetAddressResult)
	success, err := handler.(checkout.CheckoutService).GetAddress(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCheckoutServiceGetAddressArgs() interface{} {
	return checkout.NewCheckoutServiceGetAddressArgs()
}

func newCheckoutServiceGetAddressResult() interface{} {
	return checkout.NewCheckoutServiceGetAddressResult()
}

func checkoutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*checkout.CheckoutServiceCheckoutArgs)
	realResult := result.(*checkout.CheckoutServiceCheckoutResult)
	success, err := handler.(checkout.CheckoutService).Checkout(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCheckoutServiceCheckoutArgs() interface{} {
	return checkout.NewCheckoutServiceCheckoutArgs()
}

func newCheckoutServiceCheckoutResult() interface{} {
	return checkout.NewCheckoutServiceCheckoutResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateAddress(ctx context.Context, req *checkout.CreateAddressReq) (r *checkout.CreateAddressResp, err error) {
	var _args checkout.CheckoutServiceCreateAddressArgs
	_args.Req = req
	var _result checkout.CheckoutServiceCreateAddressResult
	if err = p.c.Call(ctx, "CreateAddress", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateAddress(ctx context.Context, req *checkout.UpdateAddressReq) (r *checkout.UpdateAddressResp, err error) {
	var _args checkout.CheckoutServiceUpdateAddressArgs
	_args.Req = req
	var _result checkout.CheckoutServiceUpdateAddressResult
	if err = p.c.Call(ctx, "UpdateAddress", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteAddress(ctx context.Context, req *checkout.DeleteAddressReq) (r *checkout.DeleteAddressResp, err error) {
	var _args checkout.CheckoutServiceDeleteAddressArgs
	_args.Req = req
	var _result checkout.CheckoutServiceDeleteAddressResult
	if err = p.c.Call(ctx, "DeleteAddress", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAddress(ctx context.Context, req *checkout.GetAddressReq) (r *checkout.GetAddressResp, err error) {
	var _args checkout.CheckoutServiceGetAddressArgs
	_args.Req = req
	var _result checkout.CheckoutServiceGetAddressResult
	if err = p.c.Call(ctx, "GetAddress", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Checkout(ctx context.Context, req *checkout.CheckoutReq) (r *checkout.CheckoutResp, err error) {
	var _args checkout.CheckoutServiceCheckoutArgs
	_args.Req = req
	var _result checkout.CheckoutServiceCheckoutResult
	if err = p.c.Call(ctx, "Checkout", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
