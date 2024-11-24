package rocketmq

import (
	"log"
	"os"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/czczcz831/tiktok-mall/app/checkout/conf"
)

var (
	CheckoutProducer golang.Producer
)

func Init() {
	os.Setenv("mq.consoleAppender.enabled", "true")
	var err error
	log.Println(conf.GetConf().RocketMQ.Endpoint)
	golang.ResetLogger()
	CheckoutProducer, err = golang.NewProducer(
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
		panic(err)
	}

	err = CheckoutProducer.Start()
	if err != nil {
		panic(err)
	}

	log.Println("start rocketmq producer success")
}
