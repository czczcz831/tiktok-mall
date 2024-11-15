package service

import (
	"context"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
	"testing"
)

func TestGetProductList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetProductListService(ctx)
	// init req and assert value

	req := &product.GetProductListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
