package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	rocketGolang "github.com/apache/rocketmq-clients/golang"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/rocketmq/producer"
	"github.com/czczcz831/tiktok-mall/app/order/biz/model"

	"github.com/czczcz831/tiktok-mall/app/order/conf"
	order "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
	product "github.com/czczcz831/tiktok-mall/client/product/kitex_gen/product"
	productAgent "github.com/czczcz831/tiktok-mall/client/product/rpc/product"
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
var delayedTime = consts.RocketOrderDelayedTime

// Run create note info
func (s *CreateOrderService) Run(req *order.CreateOrderReq) (resp *order.CreateOrderResp, err error) {
	// Finish your business logic.

	//1.Pre-process
	nodeId := conf.GetConf().NodeID

	orderUUID, err := utils.UUIDGenerate(nodeId)

	if err != nil {
		return nil, err
	}
	//1.1 Pre-decr stock
	err = preDecrStock(req)
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

	//2. Transaction Begin,To ensure the atomicity of the order creation
	//Send Half-Message to RocketMQ
	orderProducerMsg := &producer.OrderProducerMsg{
		OrderUuid: orderUUID,
		UserUuid:  req.UserUuid,
		Items:     req.Items,
	}
	createOrderBytes, err := json.Marshal(orderProducerMsg)
	if err != nil {
		return nil, err
	}

	createOrderMsg := &rocketGolang.Message{
		Topic: consts.RocketOrderTransactionTopic,
		Body:  createOrderBytes,
		Tag:   &rocketCreateOrderTag,
	}
	// RocketMQ Transaction Begin
	rocketTx := producer.OrderProducer.BeginTransaction()
	_, err = producer.OrderProducer.SendWithTransaction(context.TODO(), createOrderMsg, rocketTx)
	if err != nil {
		return nil, err
	}

	// MysqlTransaction Begin
	mysqlTx := mysql.DB.Begin()

	res := mysqlTx.Create(&createOrder)

	if res.Error != nil {
		mysqlTx.Rollback()
		rocketTx.RollBack()
		stockRollback(orderItems)
		return nil, res.Error
	}

	res = mysqlTx.Create(&orderItems)

	if res.Error != nil {
		mysqlTx.Rollback()
		rocketTx.RollBack()
		stockRollback(orderItems)
		return nil, res.Error
	}

	err = mysqlTx.Commit().Error
	if err != nil {
		mysqlTx.Rollback()
		rocketTx.RollBack()
		stockRollback(orderItems)
		return nil, err
	}

	//send delayed msgs (cancel order)

	deliveryTimestamp := time.Now().Add(delayedTime)
	cancelOrderUuidBytes := []byte(orderUUID)

	createOrderDelayedMsg := &rocketGolang.Message{
		Topic: consts.RocketOrderDelayedTopic,
		Body:  cancelOrderUuidBytes,
		Tag:   &rocketCreateOrderDelayedTag,
	}

	createOrderDelayedMsg.SetDelayTimestamp(deliveryTimestamp)

	_, err = producer.OrderProducer.Send(context.TODO(), createOrderDelayedMsg)
	if err != nil {
		mysqlTx.Rollback()
		rocketTx.RollBack()
		stockRollback(orderItems)
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

func preDecrStock(req *order.CreateOrderReq) error {
	preDecrReqItems := make([]*product.OrderItem, 0)
	for _, item := range req.Items {
		preDecrReqItems = append(preDecrReqItems, &product.OrderItem{
			Uuid:     item.ProductUuid,
			Quantity: item.Quantity,
		})
	}

	preDecrStockResp, err := productAgent.PreDecrStock(context.Background(), &product.PreDecrStockReq{
		Items: preDecrReqItems,
	})

	if err != nil {
		return err
	}

	if !preDecrStockResp.Ok {
		return errors.New("preDecrStock failed")
	}

	return nil

}

func stockRollback(orderItems []*model.OrderItem) error {
	chargeStockReqItems := make([]*product.OrderItem, 0)
	for _, item := range orderItems {
		chargeStockReqItems = append(chargeStockReqItems, &product.OrderItem{
			Uuid:     item.ProductUuid,
			Quantity: item.Quantity,
		})
	}

	chargeStockReq := &product.ChargeStockReq{
		Items: chargeStockReqItems,
	}

	chargeStockResp, err := productAgent.ChargeStock(context.Background(), chargeStockReq)
	if err != nil {
		return err
	}

	if !chargeStockResp.Ok {
		return errors.New("charge stock failed")
	}

	return nil
}
