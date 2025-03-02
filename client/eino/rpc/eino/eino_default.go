package eino

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	eino "github.com/czczcz831/tiktok-mall/client/eino/kitex_gen/eino"
)

func QueryUserOrders(ctx context.Context, req *eino.QueryUserOrdersReq, callOptions ...callopt.Option) (resp *eino.QueryUserOrdersResp, err error) {
	resp, err = defaultClient.QueryUserOrders(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "QueryUserOrders call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
