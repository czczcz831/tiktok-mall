package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/product/biz/dal"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
)

func TestDeleteProduct_Run(t *testing.T) {
	ctx := context.Background()
	sc := NewCreateProductService(ctx)
	sd := NewDeleteProductService(ctx)
	dal.Init()
	// init req and assert value
	req1 := &product.CreateProductReq{
		Name:        "小米手机",
		Description: "555j",
		Price:       999,
		Stock:       9999,
	}

	resp1, err := sc.Run(req1)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp1)

	req2 := &product.DeleteProductReq{
		Uuid: resp1.Product.Uuid,
	}

	resp2, err := sd.Run(req2)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp2)

	// todo: edit your unit test

}
