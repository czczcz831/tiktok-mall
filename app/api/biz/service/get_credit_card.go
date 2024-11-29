package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

type GetCreditCardService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCreditCardService(Context context.Context, RequestContext *app.RequestContext) *GetCreditCardService {
	return &GetCreditCardService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCreditCardService) Run(req *api.GetCreditCardReq) (resp *api.GetCreditCardResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
