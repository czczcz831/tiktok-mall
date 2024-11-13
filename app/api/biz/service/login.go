package service

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	"github.com/czczcz831/tiktok-mall/app/user/rpc/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *api.LoginReq) (resp *api.LoginResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password and confirm password are not the same")
	}

	user.DefaultClient()

	return &api.LoginResp{
		Token:        "",
		RefreshToken: "",
	}, nil
}
