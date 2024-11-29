package checkout

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	checkout "github.com/czczcz831/tiktok-mall/client/checkout/kitex_gen/checkout"
)

func CreateAddress(ctx context.Context, req *checkout.CreateAddressReq, callOptions ...callopt.Option) (resp *checkout.CreateAddressResp, err error) {
	resp, err = defaultClient.CreateAddress(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateAddress call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateAddress(ctx context.Context, req *checkout.UpdateAddressReq, callOptions ...callopt.Option) (resp *checkout.UpdateAddressResp, err error) {
	resp, err = defaultClient.UpdateAddress(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateAddress call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteAddress(ctx context.Context, req *checkout.DeleteAddressReq, callOptions ...callopt.Option) (resp *checkout.DeleteAddressResp, err error) {
	resp, err = defaultClient.DeleteAddress(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteAddress call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetAddress(ctx context.Context, req *checkout.GetAddressReq, callOptions ...callopt.Option) (resp *checkout.GetAddressResp, err error) {
	resp, err = defaultClient.GetAddress(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetAddress call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Checkout(ctx context.Context, req *checkout.CheckoutReq, callOptions ...callopt.Option) (resp *checkout.CheckoutResp, err error) {
	resp, err = defaultClient.Checkout(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Checkout call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
