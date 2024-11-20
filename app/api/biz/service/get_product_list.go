package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	product "github.com/czczcz831/tiktok-mall/client/product/kitex_gen/product"
	productAgent "github.com/czczcz831/tiktok-mall/client/product/rpc/product"
)

type GetProductListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductListService(Context context.Context, RequestContext *app.RequestContext) *GetProductListService {
	return &GetProductListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductListService) Run(req *api.GetProductListReq) (resp *api.GetProductListResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	getListResp, err := productAgent.GetProductList(h.Context, &product.GetProductListReq{
		Page:     req.Page,
		Limit:    req.Limit,
		Name:     req.Name,
		MinPrice: req.MinPrice,
		MaxPrice: req.MaxPrice,
	})

	if err != nil {
		return nil, err
	}

	products := make([]*api.Product, 0, len(getListResp.Products))

	for _, product := range getListResp.Products {
		products = append(products, &api.Product{
			UUID:        product.Uuid,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
		})
	}

	return &api.GetProductListResp{
		Total:    getListResp.Total,
		Products: products,
	}, nil
}
