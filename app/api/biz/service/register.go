package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/dal/casbin"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	user "github.com/czczcz831/tiktok-mall/client/user/kitex_gen/user"
	userAgent "github.com/czczcz831/tiktok-mall/client/user/rpc/user"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *api.RegisterReq) (resp *api.RegisterResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	if req.Password != req.ConfirmPassword {
		return nil, &packer.MyError{
			Code: packer.PASSWORD_NOT_MATCH_ERROR,
			Err:  err,
		}
	}

	rpcResp, err := userAgent.DefaultClient().Register(h.Context, &user.RegisterReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, &packer.MyError{
			Code: packer.USER_REGISTER_ERROR,
			Err:  err,
		}
	}

	//初始添加CUSTOMER角色
	casbin.CasbinEnforcer.AddRoleForUser(rpcResp.UserUuid, casbin.CUSTOMER_ROLE)

	return &api.RegisterResp{
		UserUUID: rpcResp.UserUuid,
	}, nil

}
