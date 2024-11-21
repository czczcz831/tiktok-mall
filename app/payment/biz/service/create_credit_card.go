package service

import (
	"context"
	checkout "github.com/czczcz831/tiktok-mall/app/payment/kitex_gen/checkout"
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

	return
}
