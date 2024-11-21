package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

func TestUpdateCreditCard_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewUpdateCreditCardService(ctx)
	// init req and assert value

	req := &checkout.UpdateCreditCardReq{
		CreditCard: &checkout.CreditCard{
			Uuid:               "1859568215935127552",
			UserUuid:           "1855968708639035392",
			CreditCardNumber:   "1234567890123456",
			CreditCardCvv:      123,
			CreditCardExpMonth: 1,
			CreditCardExpYear:  2025,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
