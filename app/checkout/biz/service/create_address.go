package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/checkout/conf"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
	"github.com/czczcz831/tiktok-mall/common/utils"
)

type CreateAddressService struct {
	ctx context.Context
} // NewCreateAddressService new CreateAddressService
func NewCreateAddressService(ctx context.Context) *CreateAddressService {
	return &CreateAddressService{ctx: ctx}
}

// Run create note info
func (s *CreateAddressService) Run(req *checkout.CreateAddressReq) (resp *checkout.CreateAddressResp, err error) {
	// Finish your business logic.

	nodeId := conf.GetConf().NodeID

	uuid, err := utils.UUIDGenerate(nodeId)

	if err != nil {
		return nil, err
	}

	createAddress := &model.Address{
		Base: model.Base{
			UUID: uuid,
		},
		UserUUID:      req.UserUuid,
		StreetAddress: req.StreetAddress,
		City:          req.City,
		State:         req.State,
		Country:       req.Country,
		ZipCode:       req.ZipCode,
	}

	res := mysql.DB.Create(createAddress)

	if res.Error != nil {
		return nil, res.Error
	}

	return &checkout.CreateAddressResp{
		Address: &checkout.Address{
			Uuid:          createAddress.UUID,
			UserUuid:      uuid,
			StreetAddress: createAddress.StreetAddress,
			City:          createAddress.City,
			State:         createAddress.State,
			Country:       createAddress.Country,
			ZipCode:       createAddress.ZipCode,
		},
	}, nil

}
