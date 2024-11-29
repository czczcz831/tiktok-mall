package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	cart "github.com/czczcz831/tiktok-mall/client/cart/kitex_gen/cart"
	cartAgent "github.com/czczcz831/tiktok-mall/client/cart/rpc/cart"
)

type AddProductToCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddProductToCartService(Context context.Context, RequestContext *app.RequestContext) *AddProductToCartService {
	return &AddProductToCartService{RequestContext: RequestContext, Context: Context}
}

func (h *AddProductToCartService) Run(req *api.AddProductToCartReq) (resp *api.AddProductToCartResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	addResp, err := cartAgent.AddProductToCart(h.Context, &cart.AddProductToCartReq{
		Item: &cart.CartItem{
			UserUuid:    req.Item.UserUUID,
			ProductUuid: req.Item.UserUUID,
			Quantity:    int32(req.Item.Quantity),
		},
	})

	if err != nil {
		return nil, err
	}

	return &api.AddProductToCartResp{
		Item: &api.CartItem{
			UserUUID:    addResp.Item.UserUuid,
			ProductUUID: addResp.Item.ProductUuid,
			Quantity:    int64(addResp.Item.Quantity),
		},
	}, nil
}
