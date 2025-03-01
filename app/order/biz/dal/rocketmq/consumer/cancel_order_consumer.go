package consumer

import (
	"context"
	"errors"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/order/conf"
	"github.com/czczcz831/tiktok-mall/common/consts"

	"github.com/czczcz831/tiktok-mall/app/order/biz/model"
	product "github.com/czczcz831/tiktok-mall/client/product/kitex_gen/product"
	productAgent "github.com/czczcz831/tiktok-mall/client/product/rpc/product"
)

var (
	delayedCancelOrderConsumer golang.SimpleConsumer
)

func delayedCancelOrderConsumerInit() error {
	var err error

	delayedCancelOrderConsumer, err = golang.NewSimpleConsumer(&golang.Config{
		Endpoint:      conf.GetConf().RocketMQ.Endpoint,
		ConsumerGroup: consts.RocketDelayOrderCancelOrderConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    conf.GetConf().RocketMQ.AccessKey,
			AccessSecret: conf.GetConf().RocketMQ.AccessKey,
		},
	},
		golang.WithAwaitDuration(awaitDuration),
		golang.WithSubscriptionExpressions(map[string]*golang.FilterExpression{
			consts.RocketOrderDelayedTopic: golang.NewFilterExpressionWithType(consts.RocketCreateOrderDelayedTag, golang.TAG),
		}),
	)

	klog.Infof("endpoint: %s", conf.GetConf().RocketMQ.Endpoint)

	if err != nil {
		klog.Fatalf("new simple consumer failed: %v", err)
	}

	err = delayedCancelOrderConsumer.Start()
	if err != nil {
		klog.Fatalf("start simple consumer failed: %v", err)
	}

	// Start handlers
	{
		go delayedCancelOrderConsumerHandler()
	}

	return nil

}

func delayedCancelOrderConsumerHandler() {
	defer delayedCancelOrderConsumer.GracefulStop()
	for {
		klog.Info("start recevie message")
		mvs, err := delayedCancelOrderConsumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
		if err != nil {
			klog.Errorf("receive message failed: %v", err)
		}
		// ack message
		for _, mv := range mvs {
			err = cancelOrderBiz(mv)
			if err != nil {
				klog.Errorf("clear cart failed: %v", err)
			}
			err = delayedCancelOrderConsumer.Ack(context.TODO(), mv)
			if err != nil {
				klog.Errorf("ack message failed: %v", err)
			}
		}
	}
}

func cancelOrderBiz(mv *golang.MessageView) error {
	//Unmarshal message

	orderUuid := string(mv.GetBody())
	//delete order

	klog.Infof("orderUuid %s", &orderUuid)

	res := mysql.DB.Model(&model.Order{}).Where("uuid = ?", orderUuid).Where("status != ?", model.OrderStatusPaid).Update("status", model.OrderStatusCancelled)
	if res.Error != nil {
		klog.Errorf("cancel order failed: %v", res.Error)
		return res.Error
	}

	//Unpaid order
	if res.RowsAffected != 0 {
		//Recharge stock
		err := rechargeStockBiz(orderUuid)
		if err != nil {
			klog.Errorf("recharge stock failed: %v", err)
			return err
		}
	}

	return nil
}

func rechargeStockBiz(orderUuid string) error {
	orderItems := make([]*model.OrderItem, 0)

	res := mysql.DB.Model(&model.OrderItem{}).Where("order_uuid = ?", orderUuid).Find(&orderItems)
	if res.Error != nil {
		klog.Errorf("find order items failed: %v", res.Error)
		return res.Error
	}

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
		klog.Errorf("charge stock failed: %v", err)
		return err
	}

	if !chargeStockResp.Ok {
		klog.Error("charge stock failed")
		return errors.New("charge stock failed")
	}
	return nil
}
