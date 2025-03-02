package tool

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	order "github.com/czczcz831/tiktok-mall/client/order/kitex_gen/order"
	orderAgent "github.com/czczcz831/tiktok-mall/client/order/rpc/order"
)

var (
	GetUserOrdersTool tool.InvokableTool
)

func InitUserOrdersTool() {
	var err error
	GetUserOrdersTool, err = utils.InferTool("get_user_orders_tool", "Get User Orders by user_uuid", GetUserOrdersFunc)
	if err != nil {
		klog.Fatalf("failed to infer tool: %v", err)
	}
}

type GetUserOrdersParams struct {
	UserUUID string `json:"user_uuid" jsonschema:"description=user_uuid"`
}

func GetUserOrdersFunc(_ context.Context, params *GetUserOrdersParams) (string, error) {

	req := &order.GetUserOrdersReq{
		UserUuid: params.UserUUID,
	}

	userOrders, err := orderAgent.GetUserOrders(context.Background(), req)
	if err != nil {
		return "", err
	}

	res, err := json.Marshal(userOrders)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
