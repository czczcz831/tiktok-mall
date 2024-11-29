package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/api/conf"
	"github.com/czczcz831/tiktok-mall/app/product/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/product/biz/model"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
	"github.com/czczcz831/tiktok-mall/common/utils"
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// Finish your business logic.

	nodeId := conf.GetConf().NodeID

	uuid, err := utils.UUIDGenerate(nodeId)

	if err != nil {
		return nil, err
	}

	productIns := &model.Product{
		Base:        model.Base{UUID: uuid},
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	res := mysql.DB.Create(productIns)
	if res.Error != nil {
		return nil, res.Error
	}

	return &product.CreateProductResp{
		Product: &product.Product{
			Uuid:        productIns.UUID,
			Name:        productIns.Name,
			Description: productIns.Description,
			Price:       productIns.Price,
			Stock:       productIns.Stock,
		},
	}, nil
}
