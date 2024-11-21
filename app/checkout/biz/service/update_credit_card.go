package service

import (
	"context"
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

	return
}
