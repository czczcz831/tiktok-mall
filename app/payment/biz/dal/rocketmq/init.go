package rocketmq

import (
	"os"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/rocketmq/consumer"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/rocketmq/producer"
)

func Init() {
	os.Setenv("mq.consoleAppender.enabled", "true")
	golang.ResetLogger()
	consumer.Init()
	producer.Init()
}
