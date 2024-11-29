package service

import (
	"context"

	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
	order "github.com/czczcz831/tiktok-mall/client/order/kitex_gen/order"
	orderAgent "github.com/czczcz831/tiktok-mall/client/order/rpc/order"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.

	total := int64(0)
	items := []*order.OrderItem{}
	for _, item := range req.Items {
		total += item.Price * int64(item.Quantity)
		items = append(items, &order.OrderItem{
			ProductUuid: item.ProductUuid,
			Quantity:    item.Quantity,
		})
	}

	//Rocketmq Transaction

	createResp, err := orderAgent.CreateOrder(s.ctx, &order.CreateOrderReq{
		UserUuid: req.UserUuid,
		Total:    total,
		Items:    items,
	})

	//Async call product service to update stock

	if err != nil {
		return nil, err
	}

	return &checkout.CheckoutResp{
		OrderUuid: createResp.Order.Uuid,
	}, nil
}
