package service

import (
	"context"
	"errors"

	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/mysql"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

type UpdateAddressService struct {
	ctx context.Context
} // NewUpdateAddressService new UpdateAddressService
func NewUpdateAddressService(ctx context.Context) *UpdateAddressService {
	return &UpdateAddressService{ctx: ctx}
}

// Run create note info
func (s *UpdateAddressService) Run(req *checkout.UpdateAddressReq) (resp *checkout.UpdateAddressResp, err error) {
	// Finish your business logic.

	updatedAddress := &model.Address{
		UserUUID:      req.Address.UserUuid,
		StreetAddress: req.Address.StreetAddress,
		City:          req.Address.City,
		State:         req.Address.State,
		Country:       req.Address.Country,
		ZipCode:       req.Address.ZipCode,
	}

	res := mysql.DB.Where("uuid = ? AND user_uuid = ?", req.Address.Uuid, req.Address.UserUuid).Updates(updatedAddress)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("address not found")
	}

	return &checkout.UpdateAddressResp{
		Address: &checkout.Address{
			Uuid:          req.Address.Uuid,
			UserUuid:      req.Address.UserUuid,
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
	}, nil
}
