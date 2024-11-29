package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

type DeleteCreditCardService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteCreditCardService(Context context.Context, RequestContext *app.RequestContext) *DeleteCreditCardService {
	return &DeleteCreditCardService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteCreditCardService) Run(req *api.DeleteCreditCardReq) (resp *api.DeleteCreditCardResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
