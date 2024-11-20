package service

import (
	"context"
	cart "github.com/czczcz831/tiktok-mall/app/cart/kitex_gen/cart"
	"testing"
)

func TestClearCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewClearCartService(ctx)
	// init req and assert value

	req := &cart.ClearCartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
