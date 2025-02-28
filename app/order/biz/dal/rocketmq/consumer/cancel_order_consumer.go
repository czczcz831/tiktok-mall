package consumer

import (
	"time"

	"context"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/order/conf"
	"github.com/czczcz831/tiktok-mall/common/consts"

	"github.com/czczcz831/tiktok-mall/app/order/biz/model"
)

var (
	delayedCancelOrderConsumer golang.SimpleConsumer
)

const (
	// maximum waiting time for receive func
	awaitDuration = time.Second * 5
	// maximum number of messages received at one time
	maxMessageNum int32 = 16
	// invisibleDuration should > 20s
	invisibleDuration = time.Second * 20
	// receive messages in a loop
)

func delayedCancelOrderConsumerInit() error {
	var err error

	delayedCancelOrderConsumer, err = golang.NewSimpleConsumer(&golang.Config{
		Endpoint:      conf.GetConf().RocketMQ.Endpoint,
		ConsumerGroup: consts.RocketDelayCancelOrderConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    conf.GetConf().RocketMQ.AccessKey,
			AccessSecret: conf.GetConf().RocketMQ.AccessKey,
		},
	},
		golang.WithAwaitDuration(awaitDuration),
		golang.WithSubscriptionExpressions(map[string]*golang.FilterExpression{
			consts.RocketOrderNormalTopic: golang.NewFilterExpressionWithType(consts.RocketCreateOrderDelayedTag, golang.TAG),
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
			err := delayedCancelOrderConsumer.Ack(context.TODO(), mv)
			if err != nil {
				klog.Errorf("ack message failed: %v", err)
			}
			err = cancelOrderBiz(mv)
			if err != nil {
				klog.Errorf("clear cart failed: %v", err)
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
	return nil
}
