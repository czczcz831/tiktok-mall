package service

import (
	"context"
	"errors"

	"github.com/czczcz831/tiktok-mall/app/product/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/product/biz/model"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
	"github.com/czczcz831/tiktok-mall/common/errno"
	_ "github.com/joho/godotenv/autoload"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	// Finish your business logic.

	res := mysql.DB.Where("uuid = ?", req.Uuid).Delete(&model.Product{})

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New(errno.ErrProductNotFound)
	}

	return &product.DeleteProductResp{
		Uuid: req.Uuid,
	}, nil
}
