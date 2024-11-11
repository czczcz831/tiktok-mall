package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/user/biz/dal/mysql"
	user "github.com/czczcz831/tiktok-mall/app/user/kitex_gen/user"
)

func TestLogin_Run(t *testing.T) {
	mysql.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{
		Email:    "czczcz831@gmail.com",
		Password: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
