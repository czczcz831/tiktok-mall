package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	order "github.com/czczcz831/tiktok-mall/client/order/kitex_gen/order"
	orderAgent "github.com/czczcz831/tiktok-mall/client/order/rpc/order"
)

type UpdateOrderAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateOrderAddressService(Context context.Context, RequestContext *app.RequestContext) *UpdateOrderAddressService {
	return &UpdateOrderAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateOrderAddressService) Run(req *api.UpdateOrderAddressReq) (resp *api.UpdateOrderAddressResp, err error) {
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

	_, err = orderAgent.UpdateOrderAddress(h.Context, &order.UpdateOrderAddressReq{
		Uuid:        req.OrderUUID,
		UserUuid:    userUUID.(string),
		AddressUuid: req.AddressUUID,
	})

	if err != nil {
		return nil, err
	}

	return &api.UpdateOrderAddressResp{
		Ok: ok,
	}, nil
}
