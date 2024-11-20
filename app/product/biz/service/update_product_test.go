package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/product/biz/dal"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
	_ "github.com/joho/godotenv/autoload"
)

func TestUpdateProduct_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewUpdateProductService(ctx)
	// init req and assert value

	req := &product.UpdateProductReq{
		Product: &product.Product{
			Uuid:        "1857372730428198912",
			Name:        "VIVO手机",
			Description: "666",
			Price:       8848,
			Stock:       1,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
