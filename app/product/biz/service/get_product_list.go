package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/product/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/product/biz/model"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
)

type GetProductListService struct {
	ctx context.Context
} // NewGetProductListService new GetProductListService
func NewGetProductListService(ctx context.Context) *GetProductListService {
	return &GetProductListService{ctx: ctx}
}

// Run create note info
func (s *GetProductListService) Run(req *product.GetProductListReq) (resp *product.GetProductListResp, err error) {
	// Finish your business logic.

	var dbProducts []*model.Product

	query := mysql.DB.Model(&model.Product{})

	if req.Name != nil {
		query.Where("name like ?", "%"+*req.Name+"%")
	}

	if req.MinPrice != nil {
		query.Where("price >= ?", req.MinPrice)
	}

	if req.MaxPrice != nil {
		query.Where("price <= ?", req.MaxPrice)
	}

	findRes := query.Offset(int(req.Page) - 1).Limit(int(req.Limit)).Find(&dbProducts)

	if findRes.Error != nil {
		return nil, findRes.Error
	}

	var total int64

	cntRes := query.Count(&total)

	if cntRes.Error != nil {
		return nil, cntRes.Error
	}

	var productResp []*product.Product
	for _, dbProduct := range dbProducts {
		productResp = append(productResp, &product.Product{
			Uuid:        dbProduct.UUID,
			Name:        dbProduct.Name,
			Description: dbProduct.Description,
			Price:       dbProduct.Price,
			Stock:       dbProduct.Stock,
		})
	}

	return &product.GetProductListResp{
		Total:    total,
		Products: productResp,
	}, nil
}
