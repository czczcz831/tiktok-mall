package service

import (
	"context"

	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
	order "github.com/czczcz831/tiktok-mall/client/order/kitex_gen/order"
	orderAgent "github.com/czczcz831/tiktok-mall/client/order/rpc/order"
	product "github.com/czczcz831/tiktok-mall/client/product/kitex_gen/product"
	productAgent "github.com/czczcz831/tiktok-mall/client/product/rpc/product"
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
		//Call product service to get product info
		product, err := productAgent.GetProduct(s.ctx, &product.GetProductReq{
			Uuid: item.ProductUuid,
		})

		if err != nil {
			return nil, err
		}

		total += product.Product.Price * int64(item.Quantity)
		items = append(items, &order.OrderItem{
			ProductUuid: item.ProductUuid,
			Price:       product.Product.Price,
			Quantity:    item.Quantity,
		})
	}

	createResp, err := orderAgent.CreateOrder(s.ctx, &order.CreateOrderReq{
		UserUuid:    req.UserUuid,
		AddressUuid: req.AddressUuid,
		Total:       total,
		Items:       items,
	})

	if err != nil {
		return nil, err
	}

	return &checkout.CheckoutResp{
		OrderUuid: createResp.Order.Uuid,
	}, nil
}
