package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/checkout/conf"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
	"github.com/czczcz831/tiktok-mall/common/utils"
)

type CreateCreditCardService struct {
	ctx context.Context
} // NewCreateCreditCardService new CreateCreditCardService
func NewCreateCreditCardService(ctx context.Context) *CreateCreditCardService {
	return &CreateCreditCardService{ctx: ctx}
}

// Run create note info
func (s *CreateCreditCardService) Run(req *checkout.CreateCreditCardReq) (resp *checkout.CreateCreditCardResp, err error) {
	// Finish your business logic.

	nodeId := conf.GetConf().NodeID

	uuid, err := utils.UUIDGenerate(nodeId)

	if err != nil {
		return nil, err
	}

	createCreditCard := &model.CreditCard{
		Base: model.Base{
			UUID: uuid,
		},
		UserUUID:           req.UserUuid,
		CreditCardNumber:   req.CreditCardNumber,
		CreditCardCVV:      req.CreditCardCvv,
		CreditCardExpMonth: req.CreditCardExpMonth,
		CreditCardExpYear:  req.CreditCardExpYear,
	}

	res := mysql.DB.Create(createCreditCard)

	if res.Error != nil {
		return nil, res.Error
	}

	return &checkout.CreateCreditCardResp{
		CreditCard: &checkout.CreditCard{
			Uuid: createCreditCard.UUID,
		},
	}, nil
}
