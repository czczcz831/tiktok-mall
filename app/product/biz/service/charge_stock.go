package service

import (
	"context"

	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
)

type ChargeStockService struct {
	ctx context.Context
} // NewChargeStockService new ChargeStockService
func NewChargeStockService(ctx context.Context) *ChargeStockService {
	return &ChargeStockService{ctx: ctx}
}

// Run create note info
func (s *ChargeStockService) Run(req *product.ChargeStockReq) (resp *product.ChargeStockResp, err error) {
	// Finish your business logic.

	for _, item := range req.Items {
		rollbackStock(item)
	}

	return &product.ChargeStockResp{
		Ok: true,
	}, nil
}
