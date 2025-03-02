package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	eino "github.com/czczcz831/tiktok-mall/client/eino/kitex_gen/eino"
	einoAgent "github.com/czczcz831/tiktok-mall/client/eino/rpc/eino"
)

type CallAssistantAgentService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCallAssistantAgentService(Context context.Context, RequestContext *app.RequestContext) *CallAssistantAgentService {
	return &CallAssistantAgentService{RequestContext: RequestContext, Context: Context}
}

func (h *CallAssistantAgentService) Run(req *api.CallAssistantAgentReq) (resp *api.CallAssistantAgentResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userUUID, ok := h.RequestContext.Get("uuid")
	if !ok {
		return nil, &packer.MyError{
			Code: packer.UNKNOWN_SERVER_ERROR,
		}
	}

	userUUIDStr, _ := userUUID.(string)

	reply, err := einoAgent.CallAssistantAgent(h.Context, &eino.CallAssistantAgentReq{
		UserUuid: userUUIDStr,
		Content:  req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &api.CallAssistantAgentResp{
		Reply: reply.Reply,
	}, nil
}
