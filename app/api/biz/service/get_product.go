package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	product "github.com/czczcz831/tiktok-mall/client/product/kitex_gen/product"
	productAgent "github.com/czczcz831/tiktok-mall/client/product/rpc/product"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *api.GetProductReq) (resp *api.GetProductResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	getResp, err := productAgent.GetProduct(h.Context, &product.GetProductReq{
		Uuid: req.UUID,
	})

	if err != nil {
		return nil, err
	}

	return &api.GetProductResp{
		Product: &api.Product{
			UUID:        getResp.Product.Uuid,
			Name:        getResp.Product.Name,
			Description: getResp.Product.Description,
			Price:       getResp.Product.Price,
			Stock:       getResp.Product.Stock,
		},
	}, nil
}
