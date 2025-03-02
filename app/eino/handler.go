package main

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/eino/biz/service"
	eino "github.com/czczcz831/tiktok-mall/app/eino/kitex_gen/eino"
)

// EinoServiceImpl implements the last service interface defined in the IDL.
type EinoServiceImpl struct{}

// CallAssistantAgent implements the EinoServiceImpl interface.
func (s *EinoServiceImpl) CallAssistantAgent(ctx context.Context, req *eino.CallAssistantAgentReq) (resp *eino.CallAssistantAgentResp, err error) {
	resp, err = service.NewCallAssistantAgentService(ctx).Run(req)

	return resp, err
}
