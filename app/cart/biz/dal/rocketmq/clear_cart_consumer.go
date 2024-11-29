package rocketmq

import (
	"encoding/json"
	"time"

	"context"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/cart/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/cart/biz/model"
	"github.com/czczcz831/tiktok-mall/app/cart/conf"
	order "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
)

var (
	clearCartConsumer golang.SimpleConsumer
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

func clearCartConsumerInit() error {
	var err error
	clearCartConsumer, err = golang.NewSimpleConsumer(&golang.Config{
		Endpoint:      conf.GetConf().RocketMQ.Endpoint,
		ConsumerGroup: conf.GetConf().RocketMQ.ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    conf.GetConf().RocketMQ.AccessKey,
			AccessSecret: conf.GetConf().RocketMQ.AccessKey,
		},
	},
		golang.WithAwaitDuration(awaitDuration),
		golang.WithSubscriptionExpressions(map[string]*golang.FilterExpression{
			conf.GetConf().RocketMQ.TxTopic: golang.SUB_ALL,
		}),
	)

	klog.Infof("topic: %s", conf.GetConf().RocketMQ.TxTopic)
	klog.Infof("consumer group: %s", conf.GetConf().RocketMQ.ConsumerGroup)
	klog.Infof("endpoint: %s", conf.GetConf().RocketMQ.Endpoint)

	if err != nil {
		klog.Fatalf("new simple consumer failed: %v", err)
	}

	err = clearCartConsumer.Start()
	if err != nil {
		klog.Fatalf("start simple consumer failed: %v", err)
	}

	// Start handlers
	{
		go clearCartOrderConsumerHandler()
	}

	return nil

}

func clearCartOrderConsumerHandler() {
	defer clearCartConsumer.GracefulStop()
	for {
		klog.Info("start recevie message")
		mvs, err := clearCartConsumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
		if err != nil {
			klog.Errorf("receive message failed: %v", err)
		}
		// ack message
		for _, mv := range mvs {
			klog.Infof("message 6666666666666666666666666666666666: %v", mv)
			err := clearCartConsumer.Ack(context.TODO(), mv)
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

func clearCartBiz(mv *golang.MessageView) error {
	//Unmarshal message

	req := &order.CreateOrderReq{}
	err := json.Unmarshal(mv.GetBody(), req)
	if err != nil {
		klog.Errorf("unmarshal message failed: %v", err)
		return err
	}

	//clear cart

	userUuid := req.UserUuid

	productIds := make([]string, 0)

	for _, item := range req.Items {
		productId := item.ProductUuid
		productIds = append(productIds, productId)
	}
	klog.Infof("productIds: %v", productIds)
	klog.Infof("userUuid: %s", userUuid)

	res := mysql.DB.Model(&model.Cart{}).Where("user_id = ?", userUuid).Where("product_id IN (?)", productIds).Delete(&model.Cart{})
	if res.Error != nil {
		klog.Errorf("delete cart failed: %v", res.Error)
		return res.Error
	}
	return nil
}
