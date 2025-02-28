package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	checkout "github.com/czczcz831/tiktok-mall/client/checkout/kitex_gen/checkout"
	checkoutAgent "github.com/czczcz831/tiktok-mall/client/checkout/rpc/checkout"
)

type GetAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetAddressService(Context context.Context, RequestContext *app.RequestContext) *GetAddressService {
	return &GetAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *GetAddressService) Run(req *api.GetAddressReq) (resp *api.GetAddressResp, err error) {
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

	getAddressResp, err := checkoutAgent.GetAddress(h.Context, &checkout.GetAddressReq{
		UserUuid: userUUIDStr,
	})

	if err != nil {
		return nil, err
	}

	addresses := make([]*api.Address, 0)
	for _, address := range getAddressResp.Addresses {
		addresses = append(addresses, &api.Address{
			UUID:          address.Uuid,
			StreetAddress: address.StreetAddress,
			City:          address.City,
			State:         address.State,
			Country:       address.Country,
			ZipCode:       address.ZipCode,
		})
	}

	return &api.GetAddressResp{
		Addresses: addresses,
	}, nil
}
