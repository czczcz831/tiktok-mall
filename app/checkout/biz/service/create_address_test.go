package service

import (
	"context"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
	"testing"
)

func TestCreateAddress_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateAddressService(ctx)
	// init req and assert value

	req := &checkout.CreateAddressReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
