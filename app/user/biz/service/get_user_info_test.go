package service

import (
	"context"
	user "github.com/czczcz831/tiktok-mall/app/user/kitex_gen/user"
	"testing"
)

func TestGetUserInfo_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetUserInfoService(ctx)
	// init req and assert value

	req := &user.GetUserInfoReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
