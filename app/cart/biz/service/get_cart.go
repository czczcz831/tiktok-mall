package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/cart/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/cart/biz/model"

	cart "github.com/czczcz831/tiktok-mall/app/cart/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.

	var getCarts []*model.Cart

	getResp := mysql.DB.Where("user_id = ?", req.UserUuid).Find(&getCarts)

	if getResp.Error != nil {
		return nil, getResp.Error
	}

	var cartItems []*cart.CartItem

	for _, v := range getCarts {
		cartItems = append(cartItems, &cart.CartItem{
			UserUuid:    v.UserID,
			ProductUuid: v.ProductID,
			Quantity:    int32(v.Quantity),
		})
	}

	return &cart.GetCartResp{
		Items: cartItems,
	}, nil
}
