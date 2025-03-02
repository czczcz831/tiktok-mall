package tool

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	checkout "github.com/czczcz831/tiktok-mall/client/checkout/kitex_gen/checkout"
	checkoutAgent "github.com/czczcz831/tiktok-mall/client/checkout/rpc/checkout"
)

var (
	CheckoutTool tool.InvokableTool
)

func InitCheckoutTool() {
	var err error
	CheckoutTool, err = utils.InferTool("checkout_tool", "Checkout to create an order", CheckoutFunc)
	if err != nil {
		klog.Fatalf("failed to infer tool: %v", err)
	}
}

type OrderItem struct {
	ProductUuid string `json:"product_uuid" jsonschema:"description=product_uuid"`
	Quantity    int64  `json:"quantity" jsonschema:"description=quantity"`
}

type CheckoutParams struct {
	UserUuid    string      `json:"user_uuid" jsonschema:"description=user_uuid"`
	FirstName   string      `json:"first_name" jsonschema:"description=first_name"`
	LastName    string      `json:"last_name" jsonschema:"description=last_name"`
	Email       string      `json:"email" jsonschema:"description=email"`
	AddressUuid string      `json:"address_uuid" jsonschema:"description=address_uuid"`
	Items       []OrderItem `json:"items" jsonschema:"description=items"`
}

func CheckoutFunc(_ context.Context, params *CheckoutParams) (string, error) {
	// 转换参数
	items := make([]*checkout.OrderItem, 0, len(params.Items))
	for _, item := range params.Items {
		items = append(items, &checkout.OrderItem{
			ProductUuid: item.ProductUuid,
			Quantity:    item.Quantity,
		})
	}

	req := &checkout.CheckoutReq{
		UserUuid:    params.UserUuid,
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		Email:       params.Email,
		AddressUuid: params.AddressUuid,
		Items:       items,
	}

	checkoutResp, err := checkoutAgent.Checkout(context.Background(), req)
	if err != nil {
		return "", err
	}

	res, err := json.Marshal(checkoutResp)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
