package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/product/biz/dal"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
	_ "github.com/joho/godotenv/autoload"
)

func TestGetProductList_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewGetProductListService(ctx)
	// init req and assert value

	name := "小米"
	var minPrice int64 = 800
	var maxPrice int64 = 1000

	req := &product.GetProductListReq{
		Page:     1,
		Limit:    10,
		Name:     &name,
		MinPrice: &minPrice,
		MaxPrice: &maxPrice,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
