package service

import (
	"context"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
)

type GetProductListService struct {
	ctx context.Context
} // NewGetProductListService new GetProductListService
func NewGetProductListService(ctx context.Context) *GetProductListService {
	return &GetProductListService{ctx: ctx}
}

// Run create note info
func (s *GetProductListService) Run(req *product.GetProductListReq) (resp *product.GetProductListResp, err error) {
	// Finish your business logic.

	return
}
