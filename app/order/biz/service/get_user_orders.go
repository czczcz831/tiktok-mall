package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/order/biz/model"
	order "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
)

type GetUserOrdersService struct {
	ctx context.Context
} // NewGetUserOrdersService new GetUserOrdersService
func NewGetUserOrdersService(ctx context.Context) *GetUserOrdersService {
	return &GetUserOrdersService{ctx: ctx}
}

// Run create note info
func (s *GetUserOrdersService) Run(req *order.GetUserOrdersReq) (resp *order.GetUserOrdersResp, err error) {
	// Finish your business logic.

	orders := make([]*model.Order, 0)
	res := mysql.DB.Where("user_uuid = ?", req.UserUuid).Find(&orders)

	if res.Error != nil {
		return nil, res.Error
	}

	if len(orders) == 0 {
		return &order.GetUserOrdersResp{
			Orders: []*order.Order{},
		}, nil
	}

	orderUUIDs := make([]string, 0, len(orders))
	for _, o := range orders {
		orderUUIDs = append(orderUUIDs, o.UUID)
	}

	orderItems := make([]*model.OrderItem, 0)
	res = mysql.DB.Where("order_uuid IN ?", orderUUIDs).Find(&orderItems)

	if res.Error != nil {
		return nil, res.Error
	}

	orderItemsMap := make(map[string][]*model.OrderItem)
	for _, item := range orderItems {
		orderItemsMap[item.OrderUUID] = append(orderItemsMap[item.OrderUUID], item)
	}

	ordersResp := make([]*order.Order, 0, len(orders))
	for _, o := range orders {
		items := orderItemsMap[o.UUID]

		orderItemsResp := make([]*order.OrderItem, 0, len(items))
		for _, item := range items {
			orderItemsResp = append(orderItemsResp, &order.OrderItem{
				ProductUuid: item.ProductUuid,
				Price:       item.Price,
				Quantity:    item.Quantity,
			})
		}

		ordersResp = append(ordersResp, &order.Order{
			Uuid:        o.UUID,
			UserUuid:    o.UserUuid,
			AddressUuid: o.AddressUuid,
			Total:       o.Total,
			Status:      int32(o.Status),
			CreatedAt:   o.CreatedAt.Unix(),
			Items:       orderItemsResp,
		})
	}

	return &order.GetUserOrdersResp{
		Orders: ordersResp,
	}, nil
}
