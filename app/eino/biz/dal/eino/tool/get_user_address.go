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
	GetUserAddressesTool tool.InvokableTool
)

func InitGetUserAddressesTool() {
	var err error
	GetUserAddressesTool, err = utils.InferTool("get_user_addresses_tool", "Get user addresses by user UUID", GetUserAddressesFunc)
	if err != nil {
		klog.Fatalf("failed to infer tool: %v", err)
	}
}

type GetUserAddressesParams struct {
	UserUuid string `json:"user_uuid" jsonschema:"description=user_uuid of the customer"`
}

func GetUserAddressesFunc(_ context.Context, params *GetUserAddressesParams) (string, error) {
	req := &checkout.GetAddressReq{
		UserUuid: params.UserUuid,
	}

	addressResp, err := checkoutAgent.GetAddress(context.Background(), req)
	if err != nil {
		return "", err
	}

	res, err := json.Marshal(addressResp)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
