package rocketmq

import (
	"context"
	"os"
	"time"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/cart/conf"
)

var (
	cartConsumer golang.SimpleConsumer
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

func Init() {
	os.Setenv("mq.consoleAppender.enabled", "true")
	var err error

	cartConsumer, err = golang.NewSimpleConsumer(&golang.Config{
		Endpoint:      conf.GetConf().RocketMQ.Endpoint,
		ConsumerGroup: conf.GetConf().RocketMQ.ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    conf.GetConf().RocketMQ.AccessKey,
			AccessSecret: conf.GetConf().RocketMQ.AccessKey,
		},
	},
		golang.WithAwaitDuration(awaitDuration),
		golang.WithSubscriptionExpressions(map[string]*golang.FilterExpression{
			conf.GetConf().RocketMQ.Topic: golang.SUB_ALL,
		}),
	)

	klog.Infof("topic: %s", conf.GetConf().RocketMQ.Topic)
	klog.Infof("consumer group: %s", conf.GetConf().RocketMQ.ConsumerGroup)
	klog.Infof("endpoint: %s", conf.GetConf().RocketMQ.Endpoint)

	if err != nil {
		klog.Fatalf("new simple consumer failed: %v", err)
	}

	err = cartConsumer.Start()
	if err != nil {
		klog.Fatalf("start simple consumer failed: %v", err)
	}

	// Start handlers
	{
		go createOrderConsumerHandler()
	}
}

func createOrderConsumerHandler() {
	defer cartConsumer.GracefulStop()
	for {
		klog.Info("start recevie message")
		mvs, err := cartConsumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
		if err != nil {
			klog.Errorf("receive message failed: %v", err)
		}
		// ack message
		for _, mv := range mvs {
			klog.Infof("message 6666666666666666666666666666666666: %v", mv)
			err := cartConsumer.Ack(context.TODO(), mv)
			if err != nil {
				klog.Errorf("ack message failed: %v", err)
			}
			err = clearCartBiz(mv)
			if err != nil {
				klog.Errorf("clear cart failed: %v", err)
			}
		}
	}
}
