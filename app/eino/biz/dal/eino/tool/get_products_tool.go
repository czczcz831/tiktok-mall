package tool

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
	product "github.com/czczcz831/tiktok-mall/client/product/kitex_gen/product"
	productAgent "github.com/czczcz831/tiktok-mall/client/product/rpc/product"
	"github.com/czczcz831/tiktok-mall/common/errno"
)

var (
	GetProductsTool tool.InvokableTool
)

func InitGetProductsTool() {
	var err error
	GetProductsTool, err = utils.InferTool("get_products_tool", "Get Products by params", GetProductsFunc)
	if err != nil {
		klog.Fatalf("failed to infer tool: %v", err)
	}
}

type GetProductsParams struct {
	Page     int32   `json:"page" jsonschema:"description=page"`
	Limit    int32   `json:"limit" jsonschema:"description=limit"`
	Name     *string `json:"name,omitempty" jsonschema:"description=name"`
	MinPrice *int64  `json:"min_price,omitempty" jsonschema:"description=min_price"`
	MaxPrice *int64  `json:"max_price,omitempty" jsonschema:"description=max_price"`
}

func GetProductsFunc(_ context.Context, params *GetProductsParams) (string, error) {

	req := &product.GetProductListReq{
		Page:     params.Page,
		Limit:    params.Limit,
		Name:     params.Name,
		MinPrice: params.MinPrice,
		MaxPrice: params.MaxPrice,
	}

	productList, err := productAgent.GetProductList(context.Background(), req)
	if err != nil {
		switch err.Error() {
		case errno.ErrProductNotFound:
			return "", &packer.MyError{
				Code: packer.PRODUCT_NOT_FOUND_ERROR,
			}
		default:
			return "", err
		}
	}

	res, err := json.Marshal(productList)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
