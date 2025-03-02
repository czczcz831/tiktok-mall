package agent

import (
	"context"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/kitex/pkg/klog"
	myTool "github.com/czczcz831/tiktok-mall/app/eino/biz/dal/eino/tool"
	"github.com/czczcz831/tiktok-mall/app/eino/conf"
)

var (
	AssistantAgent *react.Agent
)

func InitAssistantAgent() {
	// 初始化 tools
	assistantTools := []tool.BaseTool{
		myTool.GetProductsTool,
		myTool.GetUserOrdersTool,
		myTool.CheckoutTool,
		myTool.GetUserAddressesTool,
	}

	// 创建并配置 ChatModel
	temperature := float32(0.7)
	chatModel, err := ark.NewChatModel(context.Background(), &ark.ChatModelConfig{
		Model:       conf.GetConf().ArkModel,
		APIKey:      conf.GetConf().ArkApiKey,
		Temperature: &temperature,
	})
	if err != nil {
		klog.Fatalf("failed to create chat model: %v", err)
	}
	// 获取工具信息, 用于绑定到 ChatModel
	toolInfos := make([]*schema.ToolInfo, 0, len(assistantTools))
	var info *schema.ToolInfo
	for _, assistantTool := range assistantTools {
		info, err = assistantTool.Info(context.Background())
		if err != nil {
			klog.Fatalf("get ToolInfo failed, err=%v", err)
			return
		}
		toolInfos = append(toolInfos, info)
	}

	// 将 tools 绑定到 ChatModel
	err = chatModel.BindTools(toolInfos)
	if err != nil {
		klog.Fatalf("BindTools failed, err=%v", err)
		return
	}

	// 创建 agent
	AssistantAgent, err = react.NewAgent(context.Background(), &react.AgentConfig{
		Model: chatModel,
		ToolsConfig: compose.ToolsNodeConfig{
			Tools: assistantTools,
		},
		MessageModifier: react.NewPersonaModifier("你是一个购物助理，请根据用户的问题给出回答并调用合适的工具。当没有足够的参数或者参数觉得不合适，不要调用工具函数"),
	})

	if err != nil {
		klog.Fatalf("NewAgent failed, err=%v", err)
		return
	}

}
