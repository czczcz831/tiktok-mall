package service

import (
	"context"
	"errors"
	"github.com/czczcz831/tiktok-mall/app/product/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/product/biz/model"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
	"github.com/czczcz831/tiktok-mall/common/errno"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	dbProdcut := &model.Product{}
	res := mysql.DB.Model(dbProdcut).Where("uuid = ?", req.Uuid).First(dbProdcut)

	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, errors.New(errno.ErrProductNotFound)
		}
		return nil, errors.New(errno.ErrDatabaseSystem)
	}

	return &product.GetProductResp{
		Product: &product.Product{
			Uuid:        dbProdcut.UUID,
			Name:        dbProdcut.Name,
			Description: dbProdcut.Description,
			Price:       dbProdcut.Price,
			Stock:       dbProdcut.Stock,
		},
	}, nil
}
