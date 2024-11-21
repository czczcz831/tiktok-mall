package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

func TestCreateCreditCard_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewCreateCreditCardService(ctx)
	// init req and assert value

	req := &checkout.CreateCreditCardReq{
		UserUuid:           "1855968708639035392",
		CreditCardNumber:   "123",
		CreditCardCvv:      123,
		CreditCardExpMonth: 1,
		CreditCardExpYear:  2025,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
