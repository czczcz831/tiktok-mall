package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	checkout "github.com/czczcz831/tiktok-mall/client/checkout/kitex_gen/checkout"
	checkoutAgent "github.com/czczcz831/tiktok-mall/client/checkout/rpc/checkout"
)

type UpdateAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateAddressService(Context context.Context, RequestContext *app.RequestContext) *UpdateAddressService {
	return &UpdateAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateAddressService) Run(req *api.UpdateAddressReq) (resp *api.UpdateAddressResp, err error) {
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

	updateAddressResp, err := checkoutAgent.UpdateAddress(h.Context, &checkout.UpdateAddressReq{
		Address: &checkout.Address{
			Uuid:          req.Address.UUID,
			UserUuid:      userUUIDStr,
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
	})

	if err != nil {
		return nil, err
	}

	return &api.UpdateAddressResp{
		Address: &api.Address{
			UUID:          updateAddressResp.Address.Uuid,
			StreetAddress: updateAddressResp.Address.StreetAddress,
			City:          updateAddressResp.Address.City,
			State:         updateAddressResp.Address.State,
			Country:       updateAddressResp.Address.Country,
			ZipCode:       updateAddressResp.Address.ZipCode,
		},
	}, nil
}
