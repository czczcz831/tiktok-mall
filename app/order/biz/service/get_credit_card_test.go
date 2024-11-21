package service

import (
	"context"
	checkout "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/checkout"
	"testing"
)

func TestGetCreditCard_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetCreditCardService(ctx)
	// init req and assert value

	req := &checkout.GetCreditCardReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
