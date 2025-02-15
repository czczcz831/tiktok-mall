package service

import (
	"context"
	"errors"
	"math/rand"
	"time"

	rocketGolang "github.com/apache/rocketmq-clients/golang"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/rocketmq/producer"
	"github.com/czczcz831/tiktok-mall/app/payment/conf"
	payment "github.com/czczcz831/tiktok-mall/app/payment/kitex_gen/payment"
	"github.com/czczcz831/tiktok-mall/common/consts"
	"github.com/czczcz831/tiktok-mall/common/utils"
	"gorm.io/gorm"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

var (
	rocketCreatePaymentTag = consts.RocketCreatePaymentTag
	delayedTime            = time.Minute * 5
)

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.

	// call 3rd party service
	nodeId := conf.GetConf().NodeID

	transactionId, err := utils.UUIDGenerate(nodeId)
	if err != nil {
		return nil, err
	}

	// Check if transaction exist

	transactionIns := &model.Transaction{}

	findRes := mysql.DB.Where("order_uuid = ?", transactionId).First(transactionIns)

	if findRes.Error != nil && errors.Is(findRes.Error, gorm.ErrRecordNotFound) {
		transactionIns = &model.Transaction{
			Base:      model.Base{UUID: transactionId},
			UserUuid:  req.UserUuid,
			OrderUuid: req.OrderUuid,
			Amount:    req.Amount,
			Status:    model.TransactionStatusUnpaid,
		}
		createRes := mysql.DB.Create(transactionIns)
		if createRes.Error != nil {
			return nil, createRes.Error
		}
		// Send delayed msg to rocketmq
		delayedPaymentMsg := &rocketGolang.Message{
			Topic: consts.RocketPaymentNormalTopic,
			Body:  []byte(transactionIns.UUID),
			Tag:   &rocketCreatePaymentTag,
		}

		delayedPaymentMsg.SetDelayTimestamp(time.Now().Add(delayedTime))

		_, err = producer.PaymentProducer.Send(context.TODO(), delayedPaymentMsg)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, findRes.Error
	}

	if transactionIns.Status == model.TransactionStatusPaid {
		return nil, errors.New("transaction already paid")
	}

	paidRes := call3rdPartyService()
	if paidRes {
		transactionIns.Status = model.TransactionStatusPaid
	}

	saveRes := mysql.DB.Save(transactionIns)
	if saveRes.Error != nil {
		return nil, saveRes.Error
	}

	return &payment.ChargeResp{
		TransactionUuid: transactionIns.UUID,
		Success:         paidRes,
	}, nil
}

func call3rdPartyService() bool {
	// call 3rd party service
	// simulate the time cost
	time.Sleep(500 * time.Millisecond)
	// randomly return success or failure
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 0
}
