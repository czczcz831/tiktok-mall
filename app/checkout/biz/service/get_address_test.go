package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

func TestGetAddress_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewGetAddressService(ctx)
	// init req and assert value

	req := &checkout.GetAddressReq{
		UserUuid: "1855968708639035392",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
