package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/user/biz/dal/mysql"
	user "github.com/czczcz831/tiktok-mall/app/user/kitex_gen/user"
)

func TestRegister_Run(t *testing.T) {
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:    "czczcz831@gmail.com",
		Password: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
