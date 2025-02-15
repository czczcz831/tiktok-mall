package rocketmq

import (
	"os"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/czczcz831/tiktok-mall/app/cart/biz/dal/rocketmq/consumer"
)

func Init() {
	os.Setenv("mq.consoleAppender.enabled", "true")
	golang.ResetLogger()
	consumer.Init()
}
