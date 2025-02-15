package rocketmq

import (
	"os"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/rocketmq/consumer"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/rocketmq/producer"
)

func Init() {
	os.Setenv("mq.consoleAppender.enabled", "true")
	golang.ResetLogger()

	consumer.Init()
	producer.Init()

	klog.Info("start rocketmq producer success")
}
