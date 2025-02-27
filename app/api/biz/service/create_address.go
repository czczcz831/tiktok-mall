package service

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	checkout "github.com/czczcz831/tiktok-mall/client/checkout/kitex_gen/checkout"
	checkoutAgent "github.com/czczcz831/tiktok-mall/client/checkout/rpc/checkout"
)

type CreateAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateAddressService(Context context.Context, RequestContext *app.RequestContext) *CreateAddressService {
	return &CreateAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateAddressService) Run(req *api.CreateAddressReq) (resp *api.CreateAddressResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	userUUID, ok := h.RequestContext.Get("uuid")
	if !ok {
		return nil, &packer.MyError{
			Code: packer.UNKNOWN_SERVER_ERROR,
			Err:  errors.New("user uuid not found"),
		}
	}

	userUUIDStr, _ := userUUID.(string)

	createAddressResp, err := checkoutAgent.CreateAddress(h.Context, &checkout.CreateAddressReq{
		UserUuid:      userUUIDStr,
		StreetAddress: req.StreetAddress,
		City:          req.City,
		State:         req.State,
		Country:       req.Country,
		ZipCode:       req.ZipCode,
	})

	if err != nil {
		return nil, &packer.MyError{
			Code: packer.CHECKOUT_ERROR,
			Err:  err,
		}
	}

	return &api.CreateAddressResp{
		Address: &api.Address{
			UUID:          createAddressResp.Address.Uuid,
			StreetAddress: createAddressResp.Address.StreetAddress,
			City:          createAddressResp.Address.City,
			State:         createAddressResp.Address.State,
			Country:       createAddressResp.Address.Country,
			ZipCode:       createAddressResp.Address.ZipCode,
		},
	}, nil
}
