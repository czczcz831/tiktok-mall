package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/mysql"
	checkout "github.com/czczcz831/tiktok-mall/app/checkout/kitex_gen/checkout"
)

type DeleteAddressService struct {
	ctx context.Context
} // NewDeleteAddressService new DeleteAddressService
func NewDeleteAddressService(ctx context.Context) *DeleteAddressService {
	return &DeleteAddressService{ctx: ctx}
}

// Run create note info
func (s *DeleteAddressService) Run(req *checkout.DeleteAddressReq) (resp *checkout.DeleteAddressResp, err error) {
	// Finish your business logic.

	res := mysql.DB.Delete(&model.Address{}, "uuid = ?", req.Uuid)

	if res.Error != nil {
		return nil, res.Error
	}

	return &checkout.DeleteAddressResp{
		Uuid: req.Uuid,
	}, nil
}
