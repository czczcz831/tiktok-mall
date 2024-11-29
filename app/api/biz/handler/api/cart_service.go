package api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/czczcz831/tiktok-mall/app/api/biz/service"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
)

// AddProductToCart .
// @router /cart/add_product [POST]
func AddProductToCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.AddProductToCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &api.AddProductToCartResp{}
	resp, err = service.NewAddProductToCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ClearCart .
// @router /cart/:user_uuid [DELETE]
func ClearCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ClearCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &api.ClearCartResp{}
	resp, err = service.NewClearCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetCart .
// @router /cart/:user_uuid [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &api.GetCartResp{}
	resp, err = service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
