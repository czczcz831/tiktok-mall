package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	product "github.com/czczcz831/tiktok-mall/client/product/kitex_gen/product"
	productAgent "github.com/czczcz831/tiktok-mall/client/product/rpc/product"
)

type CreateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateProductService(Context context.Context, RequestContext *app.RequestContext) *CreateProductService {
	return &CreateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateProductService) Run(req *api.CreateProductReq) (resp *api.CreateProductResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	createResp, err := productAgent.CreateProduct(h.Context, &product.CreateProductReq{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	})

	if err != nil {
		return nil, err
	}

	return &api.CreateProductResp{
		Product: &api.Product{
			UUID:        createResp.Product.Uuid,
			Name:        createResp.Product.Name,
			Description: createResp.Product.Description,
			Price:       createResp.Product.Price,
			Stock:       createResp.Product.Stock,
		},
	}, nil
}
