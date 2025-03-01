package producer

import (
	"encoding/json"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/payment/conf"
)

var PaymentProducer golang.Producer

type PaymentProducerMsg struct {
	OrderUuid       string
	TransactionUuid string
}

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

func PaymentSuccessTxChecker(msg *golang.MessageView) golang.TransactionResolution {
	var paymentProducerMsg PaymentProducerMsg
	err := json.Unmarshal(msg.GetBody(), &paymentProducerMsg)
	if err != nil {
		klog.Errorf("PaymentSuccessTxChecker unmarshal error: %v", err)
		return golang.ROLLBACK
	}

	//check if transaction is paid
	transactionIns := &model.Transaction{}
	findRes := mysql.DB.Where("uuid = ?", paymentProducerMsg.TransactionUuid).First(transactionIns)
	if findRes.Error != nil {
		klog.Errorf("PaymentSuccessTxChecker find transaction failed: %v", findRes.Error)
		return golang.ROLLBACK
	}

	if transactionIns.Status == model.TransactionStatusPaid {
		return golang.COMMIT
	}

	return golang.ROLLBACK
}
