package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/mysql"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

type DeleteCreditCardService struct {
	ctx context.Context
} // NewDeleteCreditCardService new DeleteCreditCardService
func NewDeleteCreditCardService(ctx context.Context) *DeleteCreditCardService {
	return &DeleteCreditCardService{ctx: ctx}
}

// Run create note info
func (s *DeleteCreditCardService) Run(req *checkout.DeleteCreditCardReq) (resp *checkout.DeleteCreditCardResp, err error) {
	// Finish your business logic.

	res := mysql.DB.Delete(&model.CreditCard{}, "uuid = ?", req.Uuid)

	if res.Error != nil {
		return nil, res.Error
	}

	return &checkout.DeleteCreditCardResp{
		Uuid: req.Uuid,
	}, nil
}
