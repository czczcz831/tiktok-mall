package service

import (
	"context"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

type GetAddressService struct {
	ctx context.Context
} // NewGetAddressService new GetAddressService
func NewGetAddressService(ctx context.Context) *GetAddressService {
	return &GetAddressService{ctx: ctx}
}

// Run create note info
func (s *GetAddressService) Run(req *checkout.GetAddressReq) (resp *checkout.GetAddressResp, err error) {
	// Finish your business logic.

	return
}
