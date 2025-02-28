package main

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/product/biz/service"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// CreateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateProduct(ctx context.Context, req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	resp, err = service.NewCreateProductService(ctx).Run(req)

	return resp, err
}

// UpdateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	resp, err = service.NewUpdateProductService(ctx).Run(req)

	return resp, err
}

// DeleteProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	resp, err = service.NewDeleteProductService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// GetProductList implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProductList(ctx context.Context, req *product.GetProductListReq) (resp *product.GetProductListResp, err error) {
	resp, err = service.NewGetProductListService(ctx).Run(req)

	return resp, err
}

// PreDecrStock implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) PreDecrStock(ctx context.Context, req *product.PreDecrStockReq) (resp *product.PreDecrStockResp, err error) {
	resp, err = service.NewPreDecrStockService(ctx).Run(req)

	return resp, err
}
