package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	cart "github.com/czczcz831/tiktok-mall/client/cart/kitex_gen/cart"
	cartAgent "github.com/czczcz831/tiktok-mall/client/cart/rpc/cart"
)

type ClearCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewClearCartService(Context context.Context, RequestContext *app.RequestContext) *ClearCartService {
	return &ClearCartService{RequestContext: RequestContext, Context: Context}
}

func (h *ClearCartService) Run(req *api.ClearCartReq) (resp *api.ClearCartResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	clearResp, err := cartAgent.ClearCart(h.Context, &cart.ClearCartReq{
		UserUuid: req.UserUUID,
	})

	if err != nil {
		return nil, err
	}

	return &api.ClearCartResp{
		UserUUID: clearResp.UserUuid,
	}, nil
}
