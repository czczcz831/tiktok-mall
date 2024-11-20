package service

import (
	"context"

	"errors"

	"github.com/czczcz831/tiktok-mall/app/cart/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/cart/biz/model"
	"github.com/czczcz831/tiktok-mall/app/cart/conf"
	cart "github.com/czczcz831/tiktok-mall/app/cart/kitex_gen/cart"
	"github.com/czczcz831/tiktok-mall/common/utils"
)

type AddProductToCartService struct {
	ctx context.Context
} // NewAddProductToCartService new AddProductToCartService
func NewAddProductToCartService(ctx context.Context) *AddProductToCartService {
	return &AddProductToCartService{ctx: ctx}
}

// Run create note info
func (s *AddProductToCartService) Run(req *cart.AddProductToCartReq) (resp *cart.AddProductToCartResp, err error) {
	// Finish your business logic.

	cartItem := &model.Cart{}

	findRes := mysql.DB.Where("user_id = ?", req.Item.UserUuid).Where("product_id = ?", req.Item.ProductUuid).Find(cartItem)

	if findRes.Error != nil {
		return nil, findRes.Error
	}

	//if cartItem not exist
	if findRes.RowsAffected == 0 {
		//create cartItem with quantity

		if req.Item.Quantity <= 0 {
			return nil, errors.New("quantity should be positive when createing")
		}
		nodeId := conf.GetConf().NodeID
		uuid, err := utils.UUIDGenerate(nodeId)

		if err != nil {
			return nil, err
		}

		createCart := &model.Cart{
			Base: model.Base{
				UUID: uuid,
			},
			UserID:    req.Item.UserUuid,
			ProductID: req.Item.ProductUuid,
			Quantity:  uint(req.Item.Quantity),
		}

		createResp := mysql.DB.Create(createCart)
		if createResp.Error != nil {
			return nil, createResp.Error
		}
		return &cart.AddProductToCartResp{
			Item: &cart.CartItem{
				UserUuid:    req.Item.UserUuid,
				ProductUuid: req.Item.ProductUuid,
				Quantity:    req.Item.Quantity,
			},
		}, nil
	}

	//if cartItem exist, do increase or decrease till 0

	//decrease
	afterQuantity := int32(cartItem.Quantity) + req.Item.Quantity
	if afterQuantity <= 0 {
		deleteResp := mysql.DB.Delete(cartItem)
		if deleteResp.Error != nil {
			return nil, deleteResp.Error
		}
		return &cart.AddProductToCartResp{
			Item: &cart.CartItem{
				UserUuid:    req.Item.UserUuid,
				ProductUuid: req.Item.ProductUuid,
				Quantity:    int32(cartItem.Quantity) + req.Item.Quantity,
			},
		}, nil
	}

	//increase
	updateResp := mysql.DB.Model(cartItem).Where("uuid = ?", cartItem.UUID).Update("quantity", afterQuantity)

	if updateResp.Error != nil {
		return nil, updateResp.Error
	}

	if updateResp.RowsAffected == 0 {
		return nil, errors.New("!!! Produt may be deleted by other service")
	}

	return &cart.AddProductToCartResp{
		Item: &cart.CartItem{
			UserUuid:    req.Item.UserUuid,
			ProductUuid: req.Item.ProductUuid,
			Quantity:    afterQuantity,
		},
	}, nil
}
