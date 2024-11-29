package service

import (
	"context"
	"errors"

	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	order "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
)

type UpdateOrderAddressService struct {
	ctx context.Context
} // NewUpdateOrderAddressService new UpdateOrderAddressService
func NewUpdateOrderAddressService(ctx context.Context) *UpdateOrderAddressService {
	return &UpdateOrderAddressService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderAddressService) Run(req *order.UpdateOrderAddressReq) (resp *order.UpdateOrderAddressResp, err error) {
	// Finish your business logic.

	updateRes := mysql.DB.Model(&model.Order{}).Where("uuid = ?", req.Uuid).Where("user_uuid = ?", req.UserUuid).Update("address_uuid", req.AddressUuid)
	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	if updateRes.RowsAffected == 0 {
		return nil, errors.New("order not found")
	}

	return &order.UpdateOrderAddressResp{
		AddressUuid: req.AddressUuid,
	}, nil
}
