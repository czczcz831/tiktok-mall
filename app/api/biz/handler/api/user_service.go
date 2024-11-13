package api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/czczcz831/tiktok-mall/app/api/biz/service"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

// Login .
// @router /user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewLoginService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
