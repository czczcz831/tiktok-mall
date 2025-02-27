package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"

	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	checkout "github.com/czczcz831/tiktok-mall/client/checkout/kitex_gen/checkout"
	checkoutAgent "github.com/czczcz831/tiktok-mall/client/checkout/rpc/checkout"
)

type DeleteAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteAddressService(Context context.Context, RequestContext *app.RequestContext) *DeleteAddressService {
	return &DeleteAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteAddressService) Run(req *api.DeleteAddressReq) (resp *api.DeleteAddressResp, err error) {
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

	deleteAddressResp, err := checkoutAgent.DeleteAddress(h.Context, &checkout.DeleteAddressReq{
		Uuid: userUUIDStr,
	})

	if err != nil {
		return nil, err
	}

	return &api.DeleteAddressResp{
		UUID: deleteAddressResp.Uuid,
	}, nil
}
