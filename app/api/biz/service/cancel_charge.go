package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	payment "github.com/czczcz831/tiktok-mall/client/payment/kitex_gen/payment"
	paymentAgent "github.com/czczcz831/tiktok-mall/client/payment/rpc/payment"
)

type CancelChargeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCancelChargeService(Context context.Context, RequestContext *app.RequestContext) *CancelChargeService {
	return &CancelChargeService{RequestContext: RequestContext, Context: Context}
}

func (h *CancelChargeService) Run(req *api.CancelChargeReq) (resp *api.CancelChargeResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userUUID, ok := h.RequestContext.Get("uuid")
	if !ok {
		return nil, &packer.MyError{
			Code: packer.UNKNOWN_SERVER_ERROR,
		}
	}

	res, err := paymentAgent.CancelCharge(h.Context, &payment.CancelChargeReq{
		UserUuid:        userUUID.(string),
		TransactionUuid: req.TransactionUUID,
	})

	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, &packer.MyError{
			Code: packer.TRANSACTION_NOT_FOUND_ERROR,
		}
	}

	return &api.CancelChargeResp{
		Ok: true,
	}, nil
}
