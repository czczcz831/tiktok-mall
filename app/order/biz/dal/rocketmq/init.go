package rocketmq

import (
	"log"
	"os"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/checkout/conf"
)

var (
	CreateOrderTxProducer golang.Producer
)

func Init() {
	os.Setenv("mq.consoleAppender.enabled", "true")
	var err error
	log.Println(conf.GetConf().RocketMQ.Endpoint)
	golang.ResetLogger()

	err = delayedCancelOrderConsumerInit()
	if err != nil {
		klog.Fatal(err)
	}

	CreateOrderTxProducer, err = golang.NewProducer(
		&golang.Config{
			Endpoint: conf.GetConf().RocketMQ.Endpoint,
			Credentials: &credentials.SessionCredentials{
				AccessKey:    conf.GetConf().RocketMQ.AccessKey,
				AccessSecret: conf.GetConf().RocketMQ.SecretKey,
			},
		},
		golang.WithTransactionChecker(&golang.TransactionChecker{
			Check: CreateOrderTxChecker,
		}),
		golang.WithTopics(conf.GetConf().RocketMQ.Topic),
	)

	if err != nil {
		klog.Fatalf("new producer failed: %v", err)
	}

	err = CreateOrderTxProducer.Start()
	if err != nil {
		klog.Fatalf("start producer failed: %v", err)
	}

	log.Println("start rocketmq producer success")
}
