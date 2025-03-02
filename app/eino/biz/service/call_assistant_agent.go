package service

import (
	"context"

	"github.com/cloudwego/eino/schema"
	agent "github.com/czczcz831/tiktok-mall/app/eino/biz/dal/eino/agent"
	eino "github.com/czczcz831/tiktok-mall/app/eino/kitex_gen/eino"
)

type CallAssistantAgentService struct {
	ctx context.Context
} // NewCallAssistantAgentService new CallAssistantAgentService
func NewCallAssistantAgentService(ctx context.Context) *CallAssistantAgentService {
	return &CallAssistantAgentService{ctx: ctx}
}

// Run create note info
func (s *CallAssistantAgentService) Run(req *eino.CallAssistantAgentReq) (resp *eino.CallAssistantAgentResp, err error) {
	// Finish your business logic.

	reply, err := agent.AssistantAgent.Generate(context.Background(), []*schema.Message{
		{
			Role:    schema.User,
			Content: "用户UUID:" + req.UserUuid + " 用户的问题:" + req.Content,
		},
	})
	if err != nil {
		return nil, err
	}

	return &eino.CallAssistantAgentResp{
		Reply: reply.Content,
	}, nil
}
