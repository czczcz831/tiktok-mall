package service

import (
	"context"
	"testing"

	auth "github.com/czczcz831/tiktok-mall/app/auth/kitex_gen/auth"
)

func TestVerifyTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewVerifyTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.VerifyTokenReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
