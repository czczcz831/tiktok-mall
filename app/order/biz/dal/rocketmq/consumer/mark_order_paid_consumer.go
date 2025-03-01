package consumer

import (
	"encoding/json"

	"context"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/order/conf"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/rocketmq/producer"
	"github.com/czczcz831/tiktok-mall/common/consts"

	"github.com/czczcz831/tiktok-mall/app/order/biz/model"
)

var (
	markOrderPaidConsumer golang.SimpleConsumer
)

func markOrderPaidConsumerInit() error {
	var err error

	markOrderPaidConsumer, err = golang.NewSimpleConsumer(&golang.Config{
		Endpoint:      conf.GetConf().RocketMQ.Endpoint,
		ConsumerGroup: consts.RocketMarkOrderPaidConsumerGroup,
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

	err = markOrderPaidConsumer.Start()
	if err != nil {
		klog.Fatalf("start simple consumer failed: %v", err)
	}

	// Start handlers
	{
		go markOrderPaidConsumerHandler()
	}

	return nil

}

func markOrderPaidConsumerHandler() {
	defer markOrderPaidConsumer.GracefulStop()
	for {
		klog.Info("start recevie message")
		mvs, err := markOrderPaidConsumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
		if err != nil {
			klog.Errorf("receive message failed: %v", err)
		}
		// ack message
		for _, mv := range mvs {
			err = markOrderPaidBiz(mv)
			if err != nil {
				klog.Errorf("clear cart failed: %v", err)
			}
			err = markOrderPaidConsumer.Ack(context.TODO(), mv)
			if err != nil {
				klog.Errorf("ack message failed: %v", err)
			}
		}
	}
}

func markOrderPaidBiz(mv *golang.MessageView) error {
	//Unmarshal message
	var paymentProducerMsg producer.PaymentProducerMsg
	err := json.Unmarshal(mv.GetBody(), &paymentProducerMsg)
	if err != nil {
		klog.Errorf("markOrderPaidConsumerHandler unmarshal error: %v", err)
		return err
	}

	orderUuid := paymentProducerMsg.OrderUuid

	updateRes := mysql.DB.Model(&model.Order{}).Where("uuid = ?", orderUuid).Update("status", model.OrderStatusPaid)
	if updateRes.Error != nil {
		klog.Errorf("markOrderPaidConsumerHandler update order status failed: %v", updateRes.Error)
		return updateRes.Error
	}

	return nil

}
