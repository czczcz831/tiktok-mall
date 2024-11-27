package service

import (
	"context"
	"encoding/json"
	"time"

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
var rocketCreateOrderDelayedTag = consts.RocketCreateOrderDelayedTag
var delayedTime = time.Minute * 10

// Run create note info
func (s *CreateOrderService) Run(req *order.CreateOrderReq) (resp *order.CreateOrderResp, err error) {
	// Finish your business logic.

	nodeId := conf.GetConf().NodeID

	orderUUID, err := utils.UUIDGenerate(nodeId)

	if err != nil {
		return nil, err
	}

	createOrder := &model.Order{
		Base:        model.Base{UUID: orderUUID},
		UserUuid:    req.UserUuid,
		AddressUuid: req.AddressUuid,
		Total:       req.Total,
		Status:      model.OrderStatusUnpaid,
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
	rocketTx := rocketmq.CreateOrderTxProducer.BeginTransaction()
	_, err = rocketmq.CreateOrderTxProducer.SendWithTransaction(context.TODO(), createOrderMsg, rocketTx)
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

	//send delayed msgs (cancel order)

	deliveryTimestamp := time.Now().Add(delayedTime)
	cancelOrderUuidBytes := []byte(orderUUID)

	createOrderDelayedMsg := &rocketGolang.Message{
		Topic: conf.GetConf().RocketMQ.Topic,
		Body:  cancelOrderUuidBytes,
		Tag:   &rocketCreateOrderDelayedTag,
	}

	createOrderDelayedMsg.SetDelayTimestamp(deliveryTimestamp)

	_, err = rocketmq.CreateOrderTxProducer.Send(context.TODO(), createOrderDelayedMsg)
	if err != nil {
		mysqlTx.Rollback()
		rocketTx.RollBack()
		return nil, err
	}

	err = rocketTx.Commit()
	if err != nil {
		// Hand over to RocketMQ check-back to handle the message
		klog.Errorf("rocketmq commit failed: %v", err)
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
			Status:    int32(createOrder.Status),
			CreatedAt: createOrder.CreatedAt.Unix(),
			Items:     respOrderItems,
		},
	}, nil
}
