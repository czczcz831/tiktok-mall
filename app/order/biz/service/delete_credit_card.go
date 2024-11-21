package service

import (
	"context"
	checkout "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/checkout"
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

	return
}
