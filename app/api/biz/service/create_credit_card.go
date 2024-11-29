package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

type CreateCreditCardService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateCreditCardService(Context context.Context, RequestContext *app.RequestContext) *CreateCreditCardService {
	return &CreateCreditCardService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateCreditCardService) Run(req *api.CreateCreditCardReq) (resp *api.CreateCreditCardResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
