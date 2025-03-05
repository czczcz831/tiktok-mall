package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

type AddUserBlacklistService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddUserBlacklistService(Context context.Context, RequestContext *app.RequestContext) *AddUserBlacklistService {
	return &AddUserBlacklistService{RequestContext: RequestContext, Context: Context}
}

func (h *AddUserBlacklistService) Run(req *api.AddProductToCartReq) (resp *api.AddUserBlacklistResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
