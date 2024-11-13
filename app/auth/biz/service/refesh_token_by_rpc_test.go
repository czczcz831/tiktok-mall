package service

import (
	"context"
	"testing"

	auth "github.com/czczcz831/tiktok-mall/app/auth/kitex_gen/auth"
	_ "github.com/joho/godotenv/autoload"
)

func TestRefeshTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRefeshTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.RefeshTokenReq{
		RefreshToken: "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTczMzk3MzksInV1aWQiOiIxODU1OTY4NzA4NjM5MDM1MzkyIn0.jsPueC43tN7i_nYjcIx5fy4ufRzAgYHP4B7B9SLBMZ-ogvPuqp1g-wqAEp6e_BPdK5sZUeH4j-LixAd0RMswCQ",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
