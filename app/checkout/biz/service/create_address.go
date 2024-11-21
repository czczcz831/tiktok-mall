package service

import (
	"context"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

type CreateAddressService struct {
	ctx context.Context
} // NewCreateAddressService new CreateAddressService
func NewCreateAddressService(ctx context.Context) *CreateAddressService {
	return &CreateAddressService{ctx: ctx}
}

// Run create note info
func (s *CreateAddressService) Run(req *checkout.CreateAddressReq) (resp *checkout.CreateAddressResp, err error) {
	// Finish your business logic.

	return
}
