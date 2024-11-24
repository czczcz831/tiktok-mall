package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	order "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
)

type GetOrderService struct {
	ctx context.Context
} // NewGetOrderService new GetOrderService
func NewGetOrderService(ctx context.Context) *GetOrderService {
	return &GetOrderService{ctx: ctx}
}

// Run create note info
func (s *GetOrderService) Run(req *order.GetOrderReq) (resp *order.GetOrderResp, err error) {
	// Finish your business logic.

	var createOrder model.Order

	orderResp := mysql.DB.Where("uuid = ?", req.Uuid).First(createOrder)
	if orderResp.Error != nil {
		return nil, orderResp.Error
	}

	var orderItems []*model.OrderItem

	orderItemsResp := mysql.DB.Where("order_uuid = ?", req.Uuid).Find(&orderItems)
	if orderItemsResp.Error != nil {
		return nil, orderItemsResp.Error
	}

	getOrderItemsResp := make([]*order.OrderItem, 0, len(orderItems))
	for _, item := range orderItems {
		getOrderItemsResp = append(getOrderItemsResp, &order.OrderItem{
			ProductUuid: item.ProductUuid,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}

	getOrderResp := &order.GetOrderResp{
		Order: &order.Order{
			Uuid:      createOrder.UUID,
			UserUuid:  createOrder.UserUuid,
			Total:     createOrder.Total,
			IsPaid:    createOrder.IsPaid,
			CreatedAt: createOrder.CreatedAt.Unix(),
			Items:     getOrderItemsResp,
		},
	}

	return getOrderResp, nil
}
