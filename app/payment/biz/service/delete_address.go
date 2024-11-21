package service

import (
	"context"
	checkout "github.com/czczcz831/tiktok-mall/app/payment/kitex_gen/checkout"
)

type DeleteAddressService struct {
	ctx context.Context
} // NewDeleteAddressService new DeleteAddressService
func NewDeleteAddressService(ctx context.Context) *DeleteAddressService {
	return &DeleteAddressService{ctx: ctx}
}

// Run create note info
func (s *DeleteAddressService) Run(req *checkout.DeleteAddressReq) (resp *checkout.DeleteAddressResp, err error) {
	// Finish your business logic.

	return
}
