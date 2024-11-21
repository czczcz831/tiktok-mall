package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	cart "github.com/czczcz831/tiktok-mall/client/cart/kitex_gen/cart"
	cartAgent "github.com/czczcz831/tiktok-mall/client/cart/rpc/cart"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *api.GetCartReq) (resp *api.GetCartResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	getCartResp, err := cartAgent.GetCart(h.Context, &cart.GetCartReq{
		UserUuid: req.UserUUID,
	})

	if err != nil {
		return nil, err
	}

	items := make([]*api.CartItem, 0, len(getCartResp.Items))

	for _, item := range getCartResp.Items {
		items = append(items, &api.CartItem{
			UserUUID:    item.UserUuid,
			ProductUUID: item.ProductUuid,
			Quantity:    int64(item.Quantity),
		})
	}

	return &api.GetCartResp{
		Total: int64(len(items)),
		Items: items,
	}, nil
}
