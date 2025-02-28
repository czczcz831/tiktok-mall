package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/mysql"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

type GetAddressService struct {
	ctx context.Context
} // NewGetAddressService new GetAddressService
func NewGetAddressService(ctx context.Context) *GetAddressService {
	return &GetAddressService{ctx: ctx}
}

// Run create note info
func (s *GetAddressService) Run(req *checkout.GetAddressReq) (resp *checkout.GetAddressResp, err error) {
	// Finish your business logic.

	addresses := []*model.Address{}

	res := mysql.DB.Find(&addresses, "user_uuid = ?", req.UserUuid)

	if res.Error != nil {
		return nil, res.Error
	}

	addressResp := []*checkout.Address{}

	for _, address := range addresses {
		addressResp = append(addressResp, &checkout.Address{
			Uuid:          address.UUID,
			UserUuid:      address.UserUUID,
			StreetAddress: address.StreetAddress,
			City:          address.City,
			State:         address.State,
			Country:       address.Country,
			ZipCode:       address.ZipCode,
		})
	}

	return &checkout.GetAddressResp{
		Addresses: addressResp,
	}, nil
}
