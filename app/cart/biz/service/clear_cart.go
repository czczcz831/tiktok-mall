package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/cart/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/cart/biz/model"
	cart "github.com/czczcz831/tiktok-mall/app/cart/kitex_gen/cart"
)

type ClearCartService struct {
	ctx context.Context
} // NewClearCartService new ClearCartService
func NewClearCartService(ctx context.Context) *ClearCartService {
	return &ClearCartService{ctx: ctx}
}

// Run create note info
func (s *ClearCartService) Run(req *cart.ClearCartReq) (resp *cart.ClearCartResp, err error) {
	// Finish your business logic.

	deleteResp := mysql.DB.Where("user_id = ?", req.UserUuid).Delete(&model.Cart{})

	if deleteResp.Error != nil {
		return nil, deleteResp.Error
	}

	return &cart.ClearCartResp{
		UserUuid: req.UserUuid,
	}, nil
}
