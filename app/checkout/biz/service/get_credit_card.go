package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/mysql"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

type GetCreditCardService struct {
	ctx context.Context
} // NewGetCreditCardService new GetCreditCardService
func NewGetCreditCardService(ctx context.Context) *GetCreditCardService {
	return &GetCreditCardService{ctx: ctx}
}

// Run create note info
func (s *GetCreditCardService) Run(req *checkout.GetCreditCardReq) (resp *checkout.GetCreditCardResp, err error) {
	// Finish your business logic.

	creditCards := []*model.CreditCard{}

	res := mysql.DB.Where("user_uuid = ?", req.UserUuid).Find(&creditCards)

	if res.Error != nil {
		return nil, res.Error
	}

	creditCardResp := []*checkout.CreditCard{}

	for _, creditCard := range creditCards {
		creditCardResp = append(creditCardResp, &checkout.CreditCard{
			Uuid: creditCard.UUID,
		})
	}

	return &checkout.GetCreditCardResp{
		CreditCards: creditCardResp,
	}, nil
}
