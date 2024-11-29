package main

import (
	"context"
	"github.com/czczcz831/tiktok-mall/app/checkout/biz/service"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type CheckoutServiceImpl struct{}

// CreateAddress implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) CreateAddress(ctx context.Context, req *checkout.CreateAddressReq) (resp *checkout.CreateAddressResp, err error) {
	resp, err = service.NewCreateAddressService(ctx).Run(req)

	return resp, err
}

// UpdateAddress implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) UpdateAddress(ctx context.Context, req *checkout.UpdateAddressReq) (resp *checkout.UpdateAddressResp, err error) {
	resp, err = service.NewUpdateAddressService(ctx).Run(req)

	return resp, err
}

// DeleteAddress implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) DeleteAddress(ctx context.Context, req *checkout.DeleteAddressReq) (resp *checkout.DeleteAddressResp, err error) {
	resp, err = service.NewDeleteAddressService(ctx).Run(req)

	return resp, err
}

// GetAddress implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) GetAddress(ctx context.Context, req *checkout.GetAddressReq) (resp *checkout.GetAddressResp, err error) {
	resp, err = service.NewGetAddressService(ctx).Run(req)

	return resp, err
}

// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	resp, err = service.NewCheckoutService(ctx).Run(req)

	return resp, err
}
