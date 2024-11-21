package service

import (
	"context"
	checkout "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/checkout"
	"testing"
)

func TestDeleteCreditCard_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteCreditCardService(ctx)
	// init req and assert value

	req := &checkout.DeleteCreditCardReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
