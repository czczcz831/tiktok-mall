package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	order "github.com/czczcz831/tiktok-mall/client/order/kitex_gen/order"
	orderAgent "github.com/czczcz831/tiktok-mall/client/order/rpc/order"
	payment "github.com/czczcz831/tiktok-mall/client/payment/kitex_gen/payment"
	paymentAgent "github.com/czczcz831/tiktok-mall/client/payment/rpc/payment"
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

	userUUID, ok := h.RequestContext.Get("uuid")
	if !ok {
		return nil, &packer.MyError{
			Code: packer.UNKNOWN_SERVER_ERROR,
		}
	}

	userUUIDStr, _ := userUUID.(string)

	order, err := orderAgent.GetOrder(h.Context, &order.GetOrderReq{
		Uuid: req.OrderUUID,
	})
	if err != nil {
		return nil, err
	}

	payRes, err := paymentAgent.Charge(h.Context, &payment.ChargeReq{
		UserUuid:  userUUIDStr,
		OrderUuid: order.Order.Uuid,
		Amount:    order.Order.Total,
		CreditCard: &payment.CreditCard{
			CreditCardNumber:   req.CreditCard.CreditCardNumber,
			CreditCardCvv:      req.CreditCard.CreditCardCvv,
			CreditCardExpMonth: req.CreditCard.CreditCardExpMonth,
			CreditCardExpYear:  req.CreditCard.CreditCardExpYear,
		},
	})

	if err != nil {
		return nil, err
	}

	if !payRes.Success {
		return nil, &packer.MyError{
			Code: packer.CHARGE_FAILED_ERROR,
		}
	}

	return &api.ChargeResp{
		TransactionUUID: payRes.TransactionUuid,
	}, nil
}
