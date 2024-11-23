package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/order/conf"
	order "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
	"github.com/czczcz831/tiktok-mall/common/utils"
)

type CreateOrderService struct {
	ctx context.Context
} // NewCreateOrderService new CreateOrderService
func NewCreateOrderService(ctx context.Context) *CreateOrderService {
	return &CreateOrderService{ctx: ctx}
}

// Run create note info
func (s *CreateOrderService) Run(req *order.CreateOrderReq) (resp *order.CreateOrderResp, err error) {
	// Finish your business logic.

	nodeId := conf.GetConf().NodeID

	orderUUID, err := utils.UUIDGenerate(nodeId)

	if err != nil {
		return nil, err
	}

	createOrder := &model.Order{
		Base: model.Base{
			UUID: orderUUID,
		},
		UserUuid: req.UserUuid,
		Total:    req.Total,
	}

	orderItems := make([]*model.OrderItem, 0)

	for _, item := range req.Items {
		orderItems = append(orderItems, &model.OrderItem{
			OrderUUID:   orderUUID,
			ProductUuid: item.ProductUuid,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}

	// LocalTransaction Begin
	tx := mysql.DB.Begin()

	res := tx.Create(&createOrder)

	if res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}

	res = tx.Create(&orderItems)

	if res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}

	tx.Commit()

	//Transaction Commit

	respOrderItems := make([]*order.OrderItem, 0)

	for _, item := range orderItems {
		respOrderItems = append(respOrderItems, &order.OrderItem{
			ProductUuid: item.ProductUuid,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}

	return &order.CreateOrderResp{
		Order: &order.Order{
			Uuid:      orderUUID,
			UserUuid:  createOrder.UserUuid,
			Total:     createOrder.Total,
			IsPaid:    false,
			CreatedAt: createOrder.CreatedAt.Unix(),
			Items:     respOrderItems,
		},
	}, nil
}
