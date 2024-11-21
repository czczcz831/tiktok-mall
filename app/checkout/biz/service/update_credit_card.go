package service

import (
	"context"
	"errors"

	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/mysql"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

type UpdateCreditCardService struct {
	ctx context.Context
} // NewUpdateCreditCardService new UpdateCreditCardService
func NewUpdateCreditCardService(ctx context.Context) *UpdateCreditCardService {
	return &UpdateCreditCardService{ctx: ctx}
}

// Run create note info
func (s *UpdateCreditCardService) Run(req *checkout.UpdateCreditCardReq) (resp *checkout.UpdateCreditCardResp, err error) {
	// Finish your business logic.

	updatedCreditCard := &model.CreditCard{
		UserUUID:           req.CreditCard.UserUuid,
		CreditCardNumber:   req.CreditCard.CreditCardNumber,
		CreditCardCVV:      req.CreditCard.CreditCardCvv,
		CreditCardExpMonth: req.CreditCard.CreditCardExpMonth,
		CreditCardExpYear:  req.CreditCard.CreditCardExpYear,
	}

	res := mysql.DB.Where("uuid = ?", req.CreditCard.Uuid).Updates(updatedCreditCard)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("credit card not found")
	}

	return &checkout.UpdateCreditCardResp{
		CreditCard: &checkout.CreditCard{
			Uuid:               req.CreditCard.Uuid,
			UserUuid:           req.CreditCard.UserUuid,
			CreditCardNumber:   req.CreditCard.CreditCardNumber,
			CreditCardCvv:      req.CreditCard.CreditCardCvv,
			CreditCardExpMonth: req.CreditCard.CreditCardExpMonth,
			CreditCardExpYear:  req.CreditCard.CreditCardExpYear,
		},
	}, nil
}
