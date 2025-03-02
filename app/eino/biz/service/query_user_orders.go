package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/eino/biz/dal/eino"
	einothrift "github.com/czczcz831/tiktok-mall/app/eino/kitex_gen/eino"
)

// QueryUserOrdersService 提供订单查询服务
type QueryUserOrdersService struct {
	ctx context.Context
}

// NewQueryUserOrdersService 创建查询服务实例
func NewQueryUserOrdersService(ctx context.Context) *QueryUserOrdersService {
	return &QueryUserOrdersService{ctx: ctx}
}

// Run 运行查询服务，处理用户请求
func (s *QueryUserOrdersService) Run(req *einothrift.QueryUserOrdersReq) (resp *einothrift.QueryUserOrdersResp, err error) {
	klog.Infof("处理用户订单查询: 用户=%s, 查询=%s", req.UserUuid, req.QueryContent)

	// 初始化响应
	resp = einothrift.NewQueryUserOrdersResp()

	// 创建订单查询处理器
	processor := eino.NewOrderQueryProcessor(req.UserUuid, req.QueryContent)

	// 执行订单查询
	result, err := processor.QueryUserOrders(s.ctx)
	if err != nil {
		klog.Errorf("订单查询处理失败: %v", err)
		return nil, err
	}

	// 将处理结果转换为用户友好的文本
	responseText := result.ToUserFriendlyText()

	// 创建响应（注意：QueryUserOrdersResp可能没有Content字段，应根据实际定义调整）
	// resp.Content = responseText // 如果没有Content字段则注释掉此行

	// 添加日志，记录AI生成的回复内容
	klog.Infof("AI回复内容: %s", responseText)

	resp.Orders = make([]*einothrift.Order, 0, len(result.Orders))
	resp.Total = result.Total

	// 可选：将订单数据转换为thrift格式以供前端使用
	for _, orderInfo := range result.Orders {
		thriftOrder := einothrift.NewOrder()
		thriftOrder.Uuid = orderInfo.UUID
		thriftOrder.Total = orderInfo.Total
		thriftOrder.Status = orderInfo.Status

		// 转换商品项
		thriftOrder.Items = make([]*einothrift.OrderItem, 0, len(orderInfo.Items))
		for _, item := range orderInfo.Items {
			thriftItem := einothrift.NewOrderItem()
			thriftItem.ProductUuid = item.ProductUUID
			thriftItem.Price = item.Price
			thriftItem.Quantity = item.Quantity

			thriftOrder.Items = append(thriftOrder.Items, thriftItem)
		}

		resp.Orders = append(resp.Orders, thriftOrder)
	}

	klog.Infof("订单查询处理完成，返回 %d 个结果", resp.Total)
	return resp, nil
}
