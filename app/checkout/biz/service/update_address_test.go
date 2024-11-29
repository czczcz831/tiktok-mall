package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal"

	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

func TestUpdateAddress_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewUpdateAddressService(ctx)
	// init req and assert value

	req := &checkout.UpdateAddressReq{
		Address: &checkout.Address{
			Uuid:          "1859566119110283264",
			UserUuid:      "1855968708639035392",
			StreetAddress: "123 Main St",
			City:          "Anytown",
			State:         "CA",
			Country:       "USA",
			ZipCode:       12345,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
