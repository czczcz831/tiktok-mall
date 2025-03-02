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

type GetUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserInfoService(Context context.Context, RequestContext *app.RequestContext) *GetUserInfoService {
	return &GetUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserInfoService) Run(req *api.GetUserInfoReq) (resp *api.GetUserInfoResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	userUUID, ok := h.RequestContext.Get("uuid")
	if !ok {
		return nil, &packer.MyError{
			Code: packer.UNKNOWN_SERVER_ERROR,
		}
	}

	userUUIDStr, _ := userUUID.(string)

	getUserInfoResp, err := userAgent.GetUserInfo(h.Context, &user.GetUserInfoReq{
		UserUuid: userUUIDStr,
	})

	if err != nil {
		return nil, err
	}

	roles, err := casbin.CasbinEnforcer.GetRolesForUser(userUUIDStr)
	if err != nil {
		return nil, err
	}

	return &api.GetUserInfoResp{
		Email: getUserInfoResp.Email,
		Roles: roles,
	}, nil
}
