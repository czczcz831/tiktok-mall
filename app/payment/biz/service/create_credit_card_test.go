package service

import (
	"context"
	checkout "github.com/czczcz831/tiktok-mall/app/payment/kitex_gen/checkout"
	"testing"
)

func TestCreateCreditCard_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateCreditCardService(ctx)
	// init req and assert value

	req := &checkout.CreateCreditCardReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
