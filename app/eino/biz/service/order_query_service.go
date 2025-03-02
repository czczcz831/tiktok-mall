package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/eino/biz/dal/eino"
)

// OrderQueryService 提供直接调用的订单查询服务
type OrderQueryService struct{}

// NewOrderQueryService 创建订单查询服务实例
func NewOrderQueryService() *OrderQueryService {
	return &OrderQueryService{}
}

// ProcessQuery 处理用户的订单查询请求
func (s *OrderQueryService) ProcessQuery(ctx context.Context, userID, query string) (string, error) {
	klog.Infof("开始处理订单查询: 用户=%s 查询=%s", userID, query)

	// 创建订单查询处理器
	processor := eino.NewOrderQueryProcessor(userID, query)

	// 执行订单查询
	result, err := processor.QueryUserOrders(ctx)
	if err != nil {
		klog.Errorf("订单查询处理失败: %v", err)
		return "抱歉，处理您的订单查询请求时出现错误，请稍后再试。", err
	}

	// 将结果转换为用户友好的文本
	response := result.ToUserFriendlyText()

	// 记录生成的回复内容
	klog.Infof("生成的订单查询回复: %s", response)

	return response, nil
}
