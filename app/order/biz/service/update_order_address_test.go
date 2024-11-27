package service

import (
	"context"
	order "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
	"testing"
)

func TestUpdateOrderAddress_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateOrderAddressService(ctx)
	// init req and assert value

	req := &order.UpdateOrderAddressReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
