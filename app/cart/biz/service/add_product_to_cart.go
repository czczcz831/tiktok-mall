package service

import (
	"context"
	cart "github.com/czczcz831/tiktok-mall/app/cart/kitex_gen/cart"
)

type AddProductToCartService struct {
	ctx context.Context
} // NewAddProductToCartService new AddProductToCartService
func NewAddProductToCartService(ctx context.Context) *AddProductToCartService {
	return &AddProductToCartService{ctx: ctx}
}

// Run create note info
func (s *AddProductToCartService) Run(req *cart.AddProductToCartReq) (resp *cart.AddProductToCartResp, err error) {
	// Finish your business logic.

	return
}
