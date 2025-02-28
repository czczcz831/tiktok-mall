package product

import (
	"context"
	product "github.com/czczcz831/tiktok-mall/client/product/kitex_gen/product"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/czczcz831/tiktok-mall/client/product/kitex_gen/product/productservice"
)

type RPCClient interface {
	KitexClient() productservice.Client
	Service() string
	CreateProduct(ctx context.Context, req *product.CreateProductReq, callOptions ...callopt.Option) (r *product.CreateProductResp, err error)
	UpdateProduct(ctx context.Context, req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.UpdateProductResp, err error)
	DeleteProduct(ctx context.Context, req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error)
	GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	GetProductList(ctx context.Context, req *product.GetProductListReq, callOptions ...callopt.Option) (r *product.GetProductListResp, err error)
	PreDecrStock(ctx context.Context, req *product.PreDecrStockReq, callOptions ...callopt.Option) (r *product.PreDecrStockResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := productservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient productservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() productservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateProduct(ctx context.Context, req *product.CreateProductReq, callOptions ...callopt.Option) (r *product.CreateProductResp, err error) {
	return c.kitexClient.CreateProduct(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateProduct(ctx context.Context, req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.UpdateProductResp, err error) {
	return c.kitexClient.UpdateProduct(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error) {
	return c.kitexClient.DeleteProduct(ctx, req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error) {
	return c.kitexClient.GetProduct(ctx, req, callOptions...)
}

func (c *clientImpl) GetProductList(ctx context.Context, req *product.GetProductListReq, callOptions ...callopt.Option) (r *product.GetProductListResp, err error) {
	return c.kitexClient.GetProductList(ctx, req, callOptions...)
}

func (c *clientImpl) PreDecrStock(ctx context.Context, req *product.PreDecrStockReq, callOptions ...callopt.Option) (r *product.PreDecrStockResp, err error) {
	return c.kitexClient.PreDecrStock(ctx, req, callOptions...)
}
