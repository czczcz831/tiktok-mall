package main

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/order/biz/service"
	"github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order.CreateOrderReq) (resp *order.CreateOrderResp, err error) {
	resp, err = service.NewCreateOrderService(ctx).Run(req)

	return resp, err
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	resp, err = service.NewMarkOrderPaidService(ctx).Run(req)

	return resp, err
}

// GetOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrder(ctx context.Context, req *order.GetOrderReq) (resp *order.GetOrderResp, err error) {
	resp, err = service.NewGetOrderService(ctx).Run(req)

	return resp, err
}

// UpdateOrderAddress implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrderAddress(ctx context.Context, req *order.UpdateOrderAddressReq) (resp *order.UpdateOrderAddressResp, err error) {
	resp, err = service.NewUpdateOrderAddressService(ctx).Run(req)

	return resp, err
}

// GetUserOrders implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetUserOrders(ctx context.Context, req *order.GetUserOrdersReq) (resp *order.GetUserOrdersResp, err error) {
	resp, err = service.NewGetUserOrdersService(ctx).Run(req)

	return resp, err
}
