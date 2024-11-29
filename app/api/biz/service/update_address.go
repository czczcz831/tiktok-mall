package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

type UpdateAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateAddressService(Context context.Context, RequestContext *app.RequestContext) *UpdateAddressService {
	return &UpdateAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateAddressService) Run(req *api.UpdateAddressReq) (resp *api.UpdateAddressResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
