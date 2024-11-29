package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

type GetAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetAddressService(Context context.Context, RequestContext *app.RequestContext) *GetAddressService {
	return &GetAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *GetAddressService) Run(req *api.GetAddressReq) (resp *api.GetAddressResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
