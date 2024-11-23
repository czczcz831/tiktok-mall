package rocketmq

import (
	"log"
	"os"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/czczcz831/tiktok-mall/app/checkout/conf"
)

var (
	CreateOrderProducer golang.Producer
)

func Init() {
	os.Setenv("mq.consoleAppender.enabled", "true")
	var err error
	log.Println(conf.GetConf().RocketMQ.Endpoint)
	golang.ResetLogger()
	CreateOrderProducer, err = golang.NewProducer(
		&golang.Config{
			Endpoint: conf.GetConf().RocketMQ.Endpoint,
			Credentials: &credentials.SessionCredentials{
				AccessKey:    conf.GetConf().RocketMQ.AccessKey,
				AccessSecret: conf.GetConf().RocketMQ.SecretKey,
			},
		},
		golang.WithTransactionChecker(&golang.TransactionChecker{
			Check: func(msg *golang.MessageView) golang.TransactionResolution {
				log.Printf("check transaction message: %v", msg)
				return golang.COMMIT
			},
		}),
		golang.WithTopics(conf.GetConf().RocketMQ.Topic),
	)

	if err != nil {
		panic(err)
	}

	err = CreateOrderProducer.Start()
	if err != nil {
		panic(err)
	}

	log.Println("start rocketmq producer success")
}
