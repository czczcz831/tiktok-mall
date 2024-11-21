package service

import (
	"context"
	checkout "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/checkout"
)

type UpdateAddressService struct {
	ctx context.Context
} // NewUpdateAddressService new UpdateAddressService
func NewUpdateAddressService(ctx context.Context) *UpdateAddressService {
	return &UpdateAddressService{ctx: ctx}
}

// Run create note info
func (s *UpdateAddressService) Run(req *checkout.UpdateAddressReq) (resp *checkout.UpdateAddressResp, err error) {
	// Finish your business logic.

	return
}
