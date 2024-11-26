package service

import (
	"context"
	"encoding/json"

	rocketGolang "github.com/apache/rocketmq-clients/golang"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/rocketmq"
	"github.com/czczcz831/tiktok-mall/app/order/conf"
	order "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
	consts "github.com/czczcz831/tiktok-mall/common/consts"
	"github.com/czczcz831/tiktok-mall/common/utils"
)

type CreateOrderService struct {
	ctx context.Context
} // NewCreateOrderService new CreateOrderService
func NewCreateOrderService(ctx context.Context) *CreateOrderService {
	return &CreateOrderService{ctx: ctx}
}

var rocketCreateOrderTag = consts.RocketCreateOrderTag

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
		orderItemUuid, err := utils.UUIDGenerate(nodeId)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, &model.OrderItem{
			Base: model.Base{
				UUID: orderItemUuid,
			},
			OrderUUID:   orderUUID,
			ProductUuid: item.ProductUuid,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}

	//Send Half-Message to RocketMQ
	createOrderBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	createOrderMsg := &rocketGolang.Message{
		Topic: conf.GetConf().RocketMQ.Topic,
		Body:  createOrderBytes,
		Tag:   &rocketCreateOrderTag,
	}
	// RocketMQ Transaction Begin
	rocketTx := rocketmq.CheckoutProducer.BeginTransaction()
	_, err = rocketmq.CheckoutProducer.SendWithTransaction(context.TODO(), createOrderMsg, rocketTx)
	if err != nil {
		return nil, err
	}

	// LocalTransaction Begin
	mysqlTx := mysql.DB.Begin()

	res := mysqlTx.Create(&createOrder)

	if res.Error != nil {
		mysqlTx.Rollback()
		rocketTx.RollBack()
		return nil, res.Error
	}

	res = mysqlTx.Create(&orderItems)

	if res.Error != nil {
		mysqlTx.Rollback()
		rocketTx.RollBack()
		return nil, res.Error
	}

	err = mysqlTx.Commit().Error
	if err != nil {
		mysqlTx.Rollback()
		rocketTx.RollBack()
		return nil, err
	}

	err = rocketTx.Commit()
	if err != nil {
		klog.Errorf("rocketmq commit failed: %v", err)
		// Hand over to RocketMQ Back-Query to handle the message
		return nil, err
	}

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
