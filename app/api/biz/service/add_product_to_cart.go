package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
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

	userUUID, ok := h.RequestContext.Get("uuid")
	if !ok {
		return nil, &packer.MyError{
			Code: packer.UNKNOWN_SERVER_ERROR,
		}
	}

	userUUIDStr, _ := userUUID.(string)

	addResp, err := cartAgent.AddProductToCart(h.Context, &cart.AddProductToCartReq{
		Item: &cart.CartItem{
			UserUuid:    userUUIDStr,
			ProductUuid: req.Item.ProductUUID,
			Quantity:    int32(req.Item.Quantity),
		},
	})

	if err != nil {
		return nil, err
	}

	return &api.AddProductToCartResp{
		Item: &api.CartItem{
			ProductUUID: addResp.Item.ProductUuid,
			Quantity:    int64(addResp.Item.Quantity),
		},
	}, nil
}
