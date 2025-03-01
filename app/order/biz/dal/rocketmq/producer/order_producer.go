package producer

import (
	"encoding/json"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/czczcz831/tiktok-mall/app/order/conf"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/order/biz/model"
	order "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
)

var (
	OrderProducer golang.Producer
	err           error
)

type OrderProducerMsg struct {
	OrderUuid string
	UserUuid  string
	Items     []*order.OrderItem
}

func orderProducetInit() {
	OrderProducer, err = golang.NewProducer(
		&golang.Config{
			Endpoint: conf.GetConf().RocketMQ.Endpoint,
			Credentials: &credentials.SessionCredentials{
				AccessKey:    conf.GetConf().RocketMQ.AccessKey,
				AccessSecret: conf.GetConf().RocketMQ.SecretKey,
			},
		},
		golang.WithTransactionChecker(&golang.TransactionChecker{
			Check: createOrderTxChecker,
		}),
	)

	if err != nil {
		klog.Fatal("init order producer failed: %v", err)
	}

	err = OrderProducer.Start()

	if err != nil {
		klog.Fatal("order producer start failed", err)
	}

}

func createOrderTxChecker(msg *golang.MessageView) golang.TransactionResolution {

	var createOrder model.Order

	var orderMsg OrderProducerMsg

	err := json.Unmarshal(msg.GetBody(), &orderMsg)

	if err != nil {
		klog.Errorf("CreateOrderTxChecker unmarshal error: %v", err)
	}

	orderResp := mysql.DB.Where("uuid = ?", orderMsg.OrderUuid).First(&createOrder)
	if orderResp.Error != nil {
		return golang.ROLLBACK
	}

	return golang.COMMIT
}
