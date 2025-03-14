package auth

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	auth "github.com/czczcz831/tiktok-mall/client/auth/kitex_gen/auth"
)

func DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq, callOptions ...callopt.Option) (resp *auth.DeliveryResp, err error) {
	resp, err = defaultClient.DeliverTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeliverTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RefeshTokenByRPC(ctx context.Context, req *auth.RefeshTokenReq, callOptions ...callopt.Option) (resp *auth.DeliveryResp, err error) {
	resp, err = defaultClient.RefeshTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RefeshTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
