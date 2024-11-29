package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

type CreateAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateAddressService(Context context.Context, RequestContext *app.RequestContext) *CreateAddressService {
	return &CreateAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateAddressService) Run(req *api.CreateAddressReq) (resp *api.CreateAddressResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
