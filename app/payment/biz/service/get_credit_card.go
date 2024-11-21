package service

import (
	"context"
	checkout "github.com/czczcz831/tiktok-mall/app/payment/kitex_gen/checkout"
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

	return
}
