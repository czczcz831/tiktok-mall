package service

import (
	"context"
	"encoding/json"
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
	delayedTime            = consts.RocketPaymentDelayedTime
)

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	//Check if transaction exist,
	//if exist,which means it tried to pay for order and failed,user is trying to pay again
	//otherwise,create a new transaction and tried to call 3rd party service to pay for order

	// Check if transaction exist

	transactionIns := &model.Transaction{}

	findRes := mysql.DB.Where("order_uuid = ?", req.OrderUuid).First(transactionIns)

	if findRes.Error != nil && errors.Is(findRes.Error, gorm.ErrRecordNotFound) {
		nodeId := conf.GetConf().NodeID

		transactionId, err := utils.UUIDGenerate(nodeId)
		if err != nil {
			return nil, err
		}
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
			Topic: consts.RocketPaymentDelayedTopic,
			Body:  []byte(transactionIns.UUID),
			Tag:   &rocketCreatePaymentTag,
		}

		delayedPaymentMsg.SetDelayTimestamp(time.Now().Add(delayedTime))

		_, err = producer.PaymentProducer.Send(context.TODO(), delayedPaymentMsg)
		if err != nil {
			return nil, err
		}

	} else if findRes.Error != nil {
		//sth goes wrong with db
		return nil, findRes.Error
	}

	if transactionIns.Status == model.TransactionStatusPaid {
		return nil, errors.New("transaction already paid")
	}
	//RocketMQ Transaction begin
	//Inform order service that transaction is paid,send half message
	tag := consts.RocketPaidSuccessTag

	paidSuccessData := &producer.PaymentProducerMsg{
		OrderUuid:       req.OrderUuid,
		TransactionUuid: transactionIns.UUID,
	}

	paidSuccessDataBytes, err := json.Marshal(paidSuccessData)
	if err != nil {
		return nil, err
	}

	paidSuccessMsg := &rocketGolang.Message{
		Topic: consts.RocketPaymentTransactionTopic,
		Body:  paidSuccessDataBytes,
		Tag:   &tag,
	}
	rocketTx := producer.PaymentProducer.BeginTransaction()
	//Send half message to rocketmq
	_, err = producer.PaymentProducer.SendWithTransaction(context.TODO(), paidSuccessMsg, rocketTx)
	if err != nil {
		return nil, err
	}

	paidRes := call3rdPartyService()
	if paidRes {
		transactionIns.Status = model.TransactionStatusPaid
		saveRes := mysql.DB.Save(transactionIns)
		if saveRes.Error != nil {
			return nil, saveRes.Error
		}
		//paid successfully,commit transaction
		rocketTx.Commit()
	} else {
		rocketTx.RollBack()
	}

	return &payment.ChargeResp{
		TransactionUuid: transactionIns.UUID,
		Success:         paidRes,
	}, nil
}

func call3rdPartyService() bool {
	// call 3rd party service
	// simulate the time cost
	time.Sleep(100 * time.Millisecond)
	// randomly return success or failure
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 0
}
