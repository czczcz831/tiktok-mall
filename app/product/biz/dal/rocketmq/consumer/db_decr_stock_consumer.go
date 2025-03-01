package consumer

import (
	"context"
	"encoding/json"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/product/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/product/biz/model"
	"github.com/czczcz831/tiktok-mall/app/product/conf"
	order "github.com/czczcz831/tiktok-mall/client/order/kitex_gen/order"
	orderAgent "github.com/czczcz831/tiktok-mall/client/order/rpc/order"
	"github.com/czczcz831/tiktok-mall/common/consts"
	"gorm.io/gorm"
)

var dbDecreaseStockConsumer golang.SimpleConsumer

func dbDecreaseStockConsumerInit() error {
	var err error

	dbDecreaseStockConsumer, err = golang.NewSimpleConsumer(&golang.Config{
		Endpoint:      conf.GetConf().RocketMQ.Endpoint,
		ConsumerGroup: consts.RocketDBDecreaseStockConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    conf.GetConf().RocketMQ.AccessKey,
			AccessSecret: conf.GetConf().RocketMQ.AccessKey,
		},
	},
		golang.WithAwaitDuration(awaitDuration),
		golang.WithSubscriptionExpressions(map[string]*golang.FilterExpression{
			consts.RocketPaymentTransactionTopic: golang.NewFilterExpressionWithType(consts.RocketPaidSuccessTag, golang.TAG),
		}),
	)

	klog.Infof("endpoint: %s", conf.GetConf().RocketMQ.Endpoint)

	if err != nil {
		klog.Fatalf("new simple consumer failed: %v", err)
	}

	err = dbDecreaseStockConsumer.Start()
	if err != nil {
		klog.Fatalf("start simple consumer failed: %v", err)
	}

	// Start handlers
	{
		go dbDecreaseStockConsumerHandler()
	}

	return nil
}

func dbDecreaseStockConsumerHandler() {
	defer dbDecreaseStockConsumer.GracefulStop()
	for {
		mvs, err := dbDecreaseStockConsumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
		if err != nil {
			klog.Errorf("receive message failed: %v", err)
		}
		// ack message
		for _, mv := range mvs {
			err = dbDecreaseStockBiz(mv)
			if err != nil {
				klog.Errorf("cancel payment failed: %v", err)
			}
			err = dbDecreaseStockConsumer.Ack(context.TODO(), mv)
			if err != nil {
				klog.Errorf("ack message failed: %v", err)
			}
		}
	}
}

type PaymentProducerMsg struct {
	OrderUuid       string
	TransactionUuid string
}

func dbDecreaseStockBiz(mv *golang.MessageView) error {
	// Unmarshal message
	var paymentProducerMsg PaymentProducerMsg
	err := json.Unmarshal(mv.GetBody(), &paymentProducerMsg)
	if err != nil {
		klog.Errorf("dbDecreaseStockBiz unmarshal error: %v", err)
		return err
	}
	//Get order items
	orderRes, err := orderAgent.GetOrder(context.Background(), &order.GetOrderReq{
		Uuid: paymentProducerMsg.OrderUuid,
	})
	if err != nil {
		klog.Errorf("dbDecreaseStockBiz get order items failed: %v", err)
		return err
	}

	orderItems := orderRes.Order.Items

	//Decrease stock
	for _, orderItem := range orderItems {
		// 使用 gorm.Expr 进行安全的库存扣减
		res := mysql.DB.Model(&model.Product{}).
			Where("uuid = ? AND stock >= ?", orderItem.ProductUuid, orderItem.Quantity).
			Update("stock", gorm.Expr("stock - ?", orderItem.Quantity))
		if res.Error != nil {
			klog.Errorf("dbDecreaseStockBiz decrease stock failed: %v", res.Error)
			return res.Error
		}
	}
	return nil
}
