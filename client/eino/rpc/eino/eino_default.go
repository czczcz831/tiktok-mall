package eino

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	eino "github.com/czczcz831/tiktok-mall/client/eino/kitex_gen/eino"
)

func CallAssistantAgent(ctx context.Context, req *eino.CallAssistantAgentReq, callOptions ...callopt.Option) (resp *eino.CallAssistantAgentResp, err error) {
	resp, err = defaultClient.CallAssistantAgent(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CallAssistantAgent call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
