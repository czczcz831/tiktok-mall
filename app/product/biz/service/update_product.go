package service

import (
	"context"

	"errors"

	"github.com/czczcz831/tiktok-mall/app/product/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/product/biz/model"

	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	// Finish your business logic.

	dbProduct := &model.Product{}

	updateResp := mysql.DB.Model(dbProduct).Where("uuid = ?", req.Product.Uuid).Updates(
		&model.Product{
			Name:        req.Product.Name,
			Description: req.Product.Description,
			Price:       req.Product.Price,
			Stock:       req.Product.Stock,
		})

	if updateResp.Error != nil {
		return nil, updateResp.Error
	}

	if updateResp.RowsAffected == 0 {
		return nil, errors.New("product not found")
	}

	return &product.UpdateProductResp{
		Product: &product.Product{
			Uuid:        dbProduct.UUID,
			Name:        dbProduct.Name,
			Description: dbProduct.Description,
			Price:       dbProduct.Price,
			Stock:       dbProduct.Stock,
		},
	}, nil

}
