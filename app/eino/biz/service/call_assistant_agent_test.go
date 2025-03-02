package service

import (
	"context"
	eino "github.com/czczcz831/tiktok-mall/app/eino/kitex_gen/eino"
	"testing"
)

func TestCallAssistantAgent_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCallAssistantAgentService(ctx)
	// init req and assert value

	req := &eino.CallAssistantAgentReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
