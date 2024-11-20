package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/cart/biz/dal"
	cart "github.com/czczcz831/tiktok-mall/app/cart/kitex_gen/cart"
	_ "github.com/joho/godotenv/autoload"
)

func TestAddProductToCart_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewAddProductToCartService(ctx)
	// init req and assert value

	req := &cart.AddProductToCartReq{
		Item: &cart.CartItem{
			UserUuid:    "1855968708639035392",
			ProductUuid: "1857354284692901888",
			Quantity:    999,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
