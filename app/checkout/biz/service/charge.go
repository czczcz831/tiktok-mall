package service

import (
	"context"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *checkout.ChargeReq) (resp *checkout.ChargeResp, err error) {
	// Finish your business logic.

	return
}
