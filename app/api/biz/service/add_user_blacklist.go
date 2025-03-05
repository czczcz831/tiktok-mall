package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/dal/casbin"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

type AddUserBlacklistService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddUserBlacklistService(Context context.Context, RequestContext *app.RequestContext) *AddUserBlacklistService {
	return &AddUserBlacklistService{RequestContext: RequestContext, Context: Context}
}

func (h *AddUserBlacklistService) Run(req *api.AddUserBlacklistReq) (resp *api.AddUserBlacklistResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	ok, err := casbin.CasbinEnforcer.AddRoleForUser(req.UserUUID, casbin.BANNED_ROLE)

	if err != nil {
		return nil, &packer.MyError{
			Code: packer.UNKNOWN_SERVER_ERROR,
			Err:  err,
		}
	}

	if !ok {
		return nil, &packer.MyError{
			Code: packer.UNKNOWN_SERVER_ERROR,
			Err:  err,
		}
	}

	return &api.AddUserBlacklistResp{
		UserUUID: req.UserUUID,
	}, nil
}
