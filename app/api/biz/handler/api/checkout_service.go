package api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/czczcz831/tiktok-mall/app/api/biz/service"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

// CreateAddress .
// @router /checkout/address [POST]
func CreateAddress(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateAddressReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &api.CreateAddressResp{}
	resp, err = service.NewCreateAddressService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateAddress .
// @router /checkout/address [PUT]
func UpdateAddress(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UpdateAddressReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &api.UpdateAddressResp{}
	resp, err = service.NewUpdateAddressService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteAddress .
// @router /checkout/address/:uuid [DELETE]
func DeleteAddress(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DeleteAddressReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &api.DeleteAddressResp{}
	resp, err = service.NewDeleteAddressService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetAddress .
// @router /checkout/address/:user_uuid [GET]
func GetAddress(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetAddressReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &api.GetAddressResp{}
	resp, err = service.NewGetAddressService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// Checkout .
// @router /checkout [POST]
func Checkout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CheckoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &api.CheckoutResp{}
	resp, err = service.NewCheckoutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
