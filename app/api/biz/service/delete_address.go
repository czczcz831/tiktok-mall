package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

type DeleteAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteAddressService(Context context.Context, RequestContext *app.RequestContext) *DeleteAddressService {
	return &DeleteAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteAddressService) Run(req *api.DeleteAddressReq) (resp *api.DeleteAddressResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
