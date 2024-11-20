package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/product/biz/dal"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
)

func TestGetProduct_Run(t *testing.T) {
	ctx := context.Background()
	dal.Init()
	s := NewGetProductService(ctx)
	// init req and assert value

	req := &product.GetProductReq{
		Uuid: "1857379168810668032",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
