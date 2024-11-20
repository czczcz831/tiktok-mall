package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/cart/biz/dal"
	cart "github.com/czczcz831/tiktok-mall/app/cart/kitex_gen/cart"
)

func TestClearCart_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewClearCartService(ctx)
	// init req and assert value

	req := &cart.ClearCartReq{
		UserUuid: "1855968708639035392",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
