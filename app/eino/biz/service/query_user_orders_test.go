package service

import (
	"context"

	"testing"

	eino "github.com/czczcz831/tiktok-mall/app/eino/kitex_gen/eino"
)

func TestQueryUserOrders_Run(t *testing.T) {
	ctx := context.Background()
	s := NewQueryUserOrdersService(ctx)
	// init req and assert value

	req := &eino.QueryUserOrdersReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
