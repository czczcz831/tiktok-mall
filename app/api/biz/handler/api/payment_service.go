package api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/czczcz831/tiktok-mall/app/api/biz/service"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

// Charge .
// @router /payment/charge [POST]
func Charge(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ChargeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &api.ChargeResp{}
	resp, err = service.NewChargeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CancelCharge .
// @router /payment/cancel [POST]
func CancelCharge(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CancelChargeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCancelChargeService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
