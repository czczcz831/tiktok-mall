package service

import (
	"context"
	cart "github.com/czczcz831/tiktok-mall/app/cart/kitex_gen/cart"
	"testing"
)

func TestAddProductToCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddProductToCartService(ctx)
	// init req and assert value

	req := &cart.AddProductToCartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
