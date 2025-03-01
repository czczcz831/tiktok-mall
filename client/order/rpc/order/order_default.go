package order

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	order "github.com/czczcz831/tiktok-mall/client/order/kitex_gen/order"
)

func CreateOrder(ctx context.Context, req *order.CreateOrderReq, callOptions ...callopt.Option) (resp *order.CreateOrderResp, err error) {
	resp, err = defaultClient.CreateOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateOrderAddress(ctx context.Context, req *order.UpdateOrderAddressReq, callOptions ...callopt.Option) (resp *order.UpdateOrderAddressResp, err error) {
	resp, err = defaultClient.UpdateOrderAddress(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateOrderAddress call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (resp *order.MarkOrderPaidResp, err error) {
	resp, err = defaultClient.MarkOrderPaid(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MarkOrderPaid call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetOrder(ctx context.Context, req *order.GetOrderReq, callOptions ...callopt.Option) (resp *order.GetOrderResp, err error) {
	resp, err = defaultClient.GetOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetUserOrders(ctx context.Context, req *order.GetUserOrdersReq, callOptions ...callopt.Option) (resp *order.GetUserOrdersResp, err error) {
	resp, err = defaultClient.GetUserOrders(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetUserOrders call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
