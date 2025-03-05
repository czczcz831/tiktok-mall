package service

import (
	"context"
	payment "github.com/czczcz831/tiktok-mall/app/payment/kitex_gen/payment"
	"testing"
)

func TestCancelCharge_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCancelChargeService(ctx)
	// init req and assert value

	req := &payment.CancelChargeReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
