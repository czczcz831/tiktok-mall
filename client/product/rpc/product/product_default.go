package product

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	product "github.com/czczcz831/tiktok-mall/client/product/kitex_gen/product"
)

func CreateProduct(ctx context.Context, req *product.CreateProductReq, callOptions ...callopt.Option) (resp *product.CreateProductResp, err error) {
	resp, err = defaultClient.CreateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateProduct(ctx context.Context, req *product.UpdateProductReq, callOptions ...callopt.Option) (resp *product.UpdateProductResp, err error) {
	resp, err = defaultClient.UpdateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteProduct(ctx context.Context, req *product.DeleteProductReq, callOptions ...callopt.Option) (resp *product.DeleteProductResp, err error) {
	resp, err = defaultClient.DeleteProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (resp *product.GetProductResp, err error) {
	resp, err = defaultClient.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProductList(ctx context.Context, req *product.GetProductListReq, callOptions ...callopt.Option) (resp *product.GetProductListResp, err error) {
	resp, err = defaultClient.GetProductList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProductList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func PreDecrStock(ctx context.Context, req *product.PreDecrStockReq, callOptions ...callopt.Option) (resp *product.PreDecrStockResp, err error) {
	resp, err = defaultClient.PreDecrStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PreDecrStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ChargeStock(ctx context.Context, req *product.ChargeStockReq, callOptions ...callopt.Option) (resp *product.ChargeStockResp, err error) {
	resp, err = defaultClient.ChargeStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ChargeStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
