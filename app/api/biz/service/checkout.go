package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	checkout "github.com/czczcz831/tiktok-mall/client/checkout/kitex_gen/checkout"
	checkoutAgent "github.com/czczcz831/tiktok-mall/client/checkout/rpc/checkout"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *api.CheckoutReq) (resp *api.CheckoutResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	reqItems := make([]*checkout.OrderItem, 0)
	for _, item := range req.Items {
		reqItems = append(reqItems, &checkout.OrderItem{
			ProductUuid: item.ProductUUID,
			Quantity:    item.Quantity,
		})
	}

	checkoutResp, err := checkoutAgent.Checkout(h.Context, &checkout.CheckoutReq{
		UserUuid:       req.UserUUID,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		AddressUuid:    req.AddressUUID,
		CreditCardUuid: req.CreditCardUUID,
		Items:          reqItems,
	})

	if err != nil {
		return nil, err
	}

	return &api.CheckoutResp{
		OrderUUID: checkoutResp.OrderUuid,
	}, nil
}
