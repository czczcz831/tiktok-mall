package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
	_ "github.com/joho/godotenv/autoload"
)

func TestCreateAddress_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewCreateAddressService(ctx)
	// init req and assert value

	req := &checkout.CreateAddressReq{
		UserUuid:      "1855968708639035392",
		StreetAddress: "123",
		City:          "123",
		State:         "123",
		Country:       "USA",
		ZipCode:       123456,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
