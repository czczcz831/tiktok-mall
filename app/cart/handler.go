package main

import (
	"context"
	"github.com/czczcz831/tiktok-mall/app/cart/biz/service"
	cart "github.com/czczcz831/tiktok-mall/app/cart/kitex_gen/cart"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddProductToCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddProductToCart(ctx context.Context, req *cart.AddProductToCartReq) (resp *cart.AddProductToCartResp, err error) {
	resp, err = service.NewAddProductToCartService(ctx).Run(req)

	return resp, err
}

// ClearCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) ClearCart(ctx context.Context, req *cart.ClearCartReq) (resp *cart.ClearCartResp, err error) {
	resp, err = service.NewClearCartService(ctx).Run(req)

	return resp, err
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	resp, err = service.NewGetCartService(ctx).Run(req)

	return resp, err
}
