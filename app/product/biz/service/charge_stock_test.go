package service

import (
	"context"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
	"testing"
)

func TestChargeStock_Run(t *testing.T) {
	ctx := context.Background()
	s := NewChargeStockService(ctx)
	// init req and assert value

	req := &product.ChargeStockReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
