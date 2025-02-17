package consumer

import (
	"context"
	"time"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/order/conf"
	"github.com/czczcz831/tiktok-mall/common/consts"

	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/model"
)

var delayedCancelPaymentConsumer golang.SimpleConsumer

const (
	// maximum waiting time for receive func
	awaitDuration = time.Second * 5
	// maximum number of messages received at one time
	maxMessageNum int32 = 16
	// invisibleDuration should > 20s
	invisibleDuration = time.Second * 20
	// receive messages in a loop
)

func delayedCancelPaymentConsumerInit() error {
	var err error

	delayedCancelPaymentConsumer, err = golang.NewSimpleConsumer(&golang.Config{
		Endpoint:      conf.GetConf().RocketMQ.Endpoint,
		ConsumerGroup: consts.RocketDelayCancelPaymentConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    conf.GetConf().RocketMQ.AccessKey,
			AccessSecret: conf.GetConf().RocketMQ.AccessKey,
		},
	},
		golang.WithAwaitDuration(awaitDuration),
		golang.WithSubscriptionExpressions(map[string]*golang.FilterExpression{
			consts.RocketPaymentNormalTopic: golang.NewFilterExpressionWithType(consts.RocketCreatePaymentTag, golang.TAG),
		}),
	)

	klog.Infof("endpoint: %s", conf.GetConf().RocketMQ.Endpoint)

	if err != nil {
		klog.Fatalf("new simple consumer failed: %v", err)
	}

	err = delayedCancelPaymentConsumer.Start()
	if err != nil {
		klog.Fatalf("start simple consumer failed: %v", err)
	}

	// Start handlers
	{
		go delayedCancelPaymentConsumerHandler()
	}

	return nil
}

func delayedCancelPaymentConsumerHandler() {
	defer delayedCancelPaymentConsumer.GracefulStop()
	for {
		klog.Info("start recevie message")
		mvs, err := delayedCancelPaymentConsumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
		if err != nil {
			klog.Errorf("receive message failed: %v", err)
		}
		// ack message
		for _, mv := range mvs {
			err := delayedCancelPaymentConsumer.Ack(context.TODO(), mv)
			if err != nil {
				klog.Errorf("ack message failed: %v", err)
			}
			err = cancelPaymentBiz(mv)
			if err != nil {
				klog.Errorf("cancel payment failed: %v", err)
			}
		}
	}
}

func cancelPaymentBiz(mv *golang.MessageView) error {
	// Unmarshal message

	transactionUuid := string(mv.GetBody())
	// delete order

	klog.Infof("transactionUuid %s", transactionUuid)

	res := mysql.DB.Model(&model.Transaction{}).Where("uuid = ?", transactionUuid).Where("status != ?", model.TransactionStatusPaid).Update("status", model.TransactionStatusCancel)
	if res.Error != nil {
		klog.Errorf("cancel payment failed: %v", res.Error)
		return res.Error
	}
	return nil
}
