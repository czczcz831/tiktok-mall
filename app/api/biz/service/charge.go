package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

type ChargeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewChargeService(Context context.Context, RequestContext *app.RequestContext) *ChargeService {
	return &ChargeService{RequestContext: RequestContext, Context: Context}
}

func (h *ChargeService) Run(req *api.ChargeReq) (resp *api.ChargeResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
