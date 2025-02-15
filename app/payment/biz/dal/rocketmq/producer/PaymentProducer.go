package producer

import (
	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/payment/conf"
)

var PaymentProducer golang.Producer

func paymentProducerInit() {
	var err error

	PaymentProducer, err = golang.NewProducer(
		&golang.Config{
			Endpoint: conf.GetConf().RocketMQ.Endpoint,
			Credentials: &credentials.SessionCredentials{
				AccessKey:    conf.GetConf().RocketMQ.AccessKey,
				AccessSecret: conf.GetConf().RocketMQ.SecretKey,
			},
		},
	)
	if err != nil {
		klog.Fatalf("new producer failed: %v", err)
	}

	err = PaymentProducer.Start()
	if err != nil {
		klog.Fatalf("start producer failed: %v", err)
	}
}
