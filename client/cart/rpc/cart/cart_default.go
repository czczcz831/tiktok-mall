package cart

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	cart "github.com/czczcz831/tiktok-mall/client/cart/kitex_gen/cart"
)

func AddProductToCart(ctx context.Context, req *cart.AddProductToCartReq, callOptions ...callopt.Option) (resp *cart.AddProductToCartResp, err error) {
	resp, err = defaultClient.AddProductToCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddProductToCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ClearCart(ctx context.Context, req *cart.ClearCartReq, callOptions ...callopt.Option) (resp *cart.ClearCartResp, err error) {
	resp, err = defaultClient.ClearCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ClearCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetCart(ctx context.Context, req *cart.GetCartReq, callOptions ...callopt.Option) (resp *cart.GetCartResp, err error) {
	resp, err = defaultClient.GetCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
