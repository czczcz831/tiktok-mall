package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/product/biz/dal"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
	_ "github.com/joho/godotenv/autoload"
)

func TestCreateProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateProductService(ctx)
	// init req and assert value
	dal.Init()

	req := &product.CreateProductReq{
		Name:        "小米手机",
		Description: "555j",
		Price:       999,
		Stock:       9999,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
