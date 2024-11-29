package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	product "github.com/czczcz831/tiktok-mall/client/product/kitex_gen/product"
	productAgent "github.com/czczcz831/tiktok-mall/client/product/rpc/product"
)

type UpdateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProductService(Context context.Context, RequestContext *app.RequestContext) *UpdateProductService {
	return &UpdateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProductService) Run(req *api.UpdateProductReq) (resp *api.UpdateProductResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	updateResp, err := productAgent.UpdateProduct(h.Context, &product.UpdateProductReq{
		Product: &product.Product{
			Uuid:        req.Product.UUID,
			Name:        req.Product.Name,
			Description: req.Product.Description,
			Price:       req.Product.Price,
			Stock:       req.Product.Stock,
		},
	})

	if err != nil {
		return nil, err
	}

	return &api.UpdateProductResp{
		Product: &api.Product{
			UUID:        updateResp.Product.Uuid,
			Name:        updateResp.Product.Name,
			Description: updateResp.Product.Description,
			Price:       updateResp.Product.Price,
			Stock:       updateResp.Product.Stock,
		},
	}, nil
}
