package eino

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
	"github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order/orderservice"
	"github.com/kitex-contrib/registry-consul/resolver"
	"github.com/czczcz831/tiktok-mall/client/order"
)

var (
	// OrderClient 是用于与Order微服务通信的客户端
	OrderClient orderservice.Client
	// Agent 是Eino的Agent实例
	Agent *compose.Agent
	// OrderModel 是处理订单查询的大语言模型
	OrderModel model.ChatModel
)

// Init 初始化Eino组件和Order客户端
func Init() error {
	// 初始化Order客户端
	r, err := resolver.NewConsulResolver("127.0.0.1:8500") // 根据您的Consul地址进行修改
	if err != nil {
		return fmt.Errorf("failed to create consul resolver: %v", err)
	}

	OrderClient = orderservice.MustNewClient(
		"order-service",
		client.WithResolver(r.(discovery.Resolver)),
		client.WithRPCTimeout(3*time.Second),
	)

	// 设置OpenAI API密钥
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}

	// 初始化OpenAI客户端
	openAIClient := openai.NewChatModel(&openai.Config{
		APIKey:         apiKey,
		Model:          "gpt-3.5-turbo", // 您可以根据需要选择模型
		Temperature:    0.7,
		TimeoutSeconds: 60,
	})
	OrderModel = openAIClient

	// 定义工具函数
	orderTools := []tool.Tool{
		getUserOrdersTool(),
		getOrderDetailTool(),
	}

	// 创建Agent
	Agent = compose.NewAgent(compose.NewAgentOptions().
		WithLLM(openAIClient).
		WithTools(orderTools).
		WithDescription("此助手可以帮助用户查询和筛选订单信息。").
		WithName("订单助手"))

	return nil
}

// getUserOrdersTool 定义获取用户订单列表的工具
func getUserOrdersTool() tool.Tool {
	return utils.NewStructTool(
		"get_user_orders",
		"获取用户的所有订单",
		getUserOrders,
		&schema.ParamsSchema{
			Type: schema.StringType,
			Properties: map[string]*schema.PropertySchema{
				"user_uuid": {Type: schema.StringType, Description: "用户ID"},
			},
			Required: []string{"user_uuid"},
		},
	)
}

// getOrderDetailTool 定义获取订单详情的工具
func getOrderDetailTool() tool.Tool {
	return utils.NewStructTool(
		"get_order_detail",
		"获取特定订单的详细信息",
		getOrderDetail,
		&schema.ParamsSchema{
			Type: schema.StringType,
			Properties: map[string]*schema.PropertySchema{
				"order_uuid": {Type: schema.StringType, Description: "订单ID"},
			},
			Required: []string{"order_uuid"},
		},
	)
}

// getUserOrders 实现获取用户订单的逻辑
func getUserOrders(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	userUUID, ok := params["user_uuid"].(string)
	if !ok || userUUID == "" {
		return nil, fmt.Errorf("invalid user_uuid")
	}

	// 调用Order微服务获取用户订单
	req := &order.GetUserOrdersReq{
		UserUuid: userUUID,
	}

	resp, err := OrderClient.GetUserOrders(ctx, req)
	if err != nil {
		klog.Errorf("failed to get user orders: %v", err)
		return nil, fmt.Errorf("获取订单失败: %v", err)
	}

	// 格式化返回结果，便于AI理解和处理
	type OrderInfo struct {
		UUID       string `json:"uuid"`
		Total      int64  `json:"total"`
		Status     int32  `json:"status"`
		StatusDesc string `json:"status_desc"`
		CreatedAt  string `json:"created_at"`
		ItemCount  int    `json:"item_count"`
	}

	orders := make([]OrderInfo, 0, len(resp.Orders))
	for _, o := range resp.Orders {
		statusDesc := getOrderStatusDesc(o.Status)
		createdTime := time.Unix(o.CreatedAt, 0).Format("2006-01-02 15:04:05")

		orders = append(orders, OrderInfo{
			UUID:       o.Uuid,
			Total:      o.Total,
			Status:     o.Status,
			StatusDesc: statusDesc,
			CreatedAt:  createdTime,
			ItemCount:  len(o.Items),
		})
	}

	return map[string]interface{}{
		"total":  resp.Total,
		"orders": orders,
	}, nil
}

// getOrderDetail 实现获取订单详情的逻辑
func getOrderDetail(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	orderUUID, ok := params["order_uuid"].(string)
	if !ok || orderUUID == "" {
		return nil, fmt.Errorf("invalid order_uuid")
	}

	// 调用Order微服务获取订单详情
	req := &order.GetOrderReq{
		Uuid: orderUUID,
	}

	resp, err := OrderClient.GetOrder(ctx, req)
	if err != nil {
		klog.Errorf("failed to get order detail: %v", err)
		return nil, fmt.Errorf("获取订单详情失败: %v", err)
	}

	if resp.Order == nil {
		return nil, fmt.Errorf("订单不存在")
	}

	// 格式化商品项
	type ItemInfo struct {
		ProductUUID string `json:"product_uuid"`
		Price       int64  `json:"price"`
		Quantity    int64  `json:"quantity"`
		Subtotal    int64  `json:"subtotal"`
	}

	items := make([]ItemInfo, 0, len(resp.Order.Items))
	for _, item := range resp.Order.Items {
		items = append(items, ItemInfo{
			ProductUUID: item.ProductUuid,
			Price:       item.Price,
			Quantity:    item.Quantity,
			Subtotal:    item.Price * item.Quantity,
		})
	}

	// 格式化返回结果
	return map[string]interface{}{
		"uuid":         resp.Order.Uuid,
		"user_uuid":    resp.Order.UserUuid,
		"address_uuid": resp.Order.AddressUuid,
		"total":        resp.Order.Total,
		"status":       resp.Order.Status,
		"status_desc":  getOrderStatusDesc(resp.Order.Status),
		"created_at":   time.Unix(resp.Order.CreatedAt, 0).Format("2006-01-02 15:04:05"),
		"items":        items,
	}, nil
}

// getOrderStatusDesc 根据状态码返回状态描述
func getOrderStatusDesc(status int32) string {
	switch status {
	case 0:
		return "待付款"
	case 1:
		return "已付款"
	case 2:
		return "已发货"
	case 3:
		return "已完成"
	case 4:
		return "已取消"
	default:
		return "未知状态"
	}
}

// ProcessOrderQuery 处理用户的订单查询请求
func ProcessOrderQuery(ctx context.Context, userID string, query string) (string, error) {
	// 构建系统提示，指导模型如何处理订单查询
	systemPrompt := `你是一个订单查询助手，可以帮助用户查询和筛选订单信息。请根据用户的请求，调用合适的工具来获取订单信息，然后根据用户的需求对订单进行筛选和分析。
在分析订单时，你可以考虑以下因素：
1. 订单的状态（待付款、已付款、已发货、已完成、已取消）
2. 订单的创建时间
3. 订单的金额
4. 订单中的商品数量

请提供清晰、简洁的回复，并确保根据用户需求进行筛选。如果用户没有明确的筛选条件，应返回所有订单的概要信息。`

	// 设置用户上下文信息
	userContext := fmt.Sprintf("用户ID: %s", userID)

	// 使用Agent处理查询
	resp, err := Agent.Chat(ctx, &schema.ChatRequest{
		Messages: []*schema.ChatMessage{
			{Role: schema.ChatMessageRoleSystem, Content: systemPrompt},
			{Role: schema.ChatMessageRoleSystem, Content: userContext},
			{Role: schema.ChatMessageRoleUser, Content: query},
		},
	})

	if err != nil {
		klog.Errorf("Agent chat failed: %v", err)
		return "", fmt.Errorf("处理查询失败: %v", err)
	}

	// 返回处理结果
	content := resp.Message.Content
	if content == "" {
		content = "抱歉，我无法处理您的请求，请稍后再试。"
	}

	return content, nil
}

// OrderClient 获取订单服务客户端
func GetOrderClient() orderservice.Client {
	return orderclient.GetOrderClient()
}

// 处理订单查询的结构体
type OrderQueryProcessor struct {
	// 用于存储用户查询的上下文和参数
	UserID string
	Query  string
}

// NewOrderQueryProcessor 创建订单查询处理器
func NewOrderQueryProcessor(userID, query string) *OrderQueryProcessor {
	return &OrderQueryProcessor{
		UserID: userID,
		Query:  query,
	}
}

// QueryUserOrders 查询用户订单
func (p *OrderQueryProcessor) QueryUserOrders(ctx context.Context) (*OrderQueryResult, error) {
	// 使用Order微服务客户端获取用户订单
	client := GetOrderClient()
	req := &order.GetUserOrdersReq{
		UserUuid: p.UserID,
	}
	
	resp, err := client.GetUserOrders(ctx, req)
	if err != nil {
		klog.Errorf("获取用户订单失败: %v", err)
		return nil, fmt.Errorf("获取订单数据出错: %v", err)
	}
	
	// 转换订单为我们的内部格式
	orders := make([]OrderInfo, 0, len(resp.Orders))
	for _, o := range resp.Orders {
		statusDesc := getOrderStatusDesc(o.Status)
		createdTime := time.Unix(o.CreatedAt, 0).Format("2006-01-02 15:04:05")
		
		// 转换商品项
		items := make([]OrderItemInfo, 0, len(o.Items))
		for _, item := range o.Items {
			items = append(items, OrderItemInfo{
				ProductUUID: item.ProductUuid,
				Price:       item.Price,
				Quantity:    item.Quantity,
				Subtotal:    item.Price * item.Quantity,
			})
		}
		
		orders = append(orders, OrderInfo{
			UUID:       o.Uuid,
			Total:      o.Total,
			Status:     o.Status,
			StatusDesc: statusDesc,
			CreatedAt:  createdTime,
			Items:      items,
		})
	}
	
	result := &OrderQueryResult{
		Total:  resp.Total,
		Orders: orders,
		Query:  p.Query,
	}
	
	// 分析用户查询并筛选订单
	filteredResult, err := p.analyzeAndFilterOrders(result)
	if err != nil {
		klog.Warnf("筛选订单过程中出现错误: %v", err)
		// 返回原始结果，而不是错误
		return result, nil
	}
	
	return filteredResult, nil
}

// 分析用户查询并筛选订单
func (p *OrderQueryProcessor) analyzeAndFilterOrders(result *OrderQueryResult) (*OrderQueryResult, error) {
	// 简单查询分析 - 在实际情况下，这里可以使用更复杂的LLM进行分析
	query := p.Query
	
	// 创建一个筛选后的结果
	filteredResult := &OrderQueryResult{
		Total:  0,
		Orders: []OrderInfo{},
		Query:  query,
	}
	
	// 基于状态筛选
	var statusFilter *int32 = nil
	if contains(query, "待付款") {
		var status int32 = 0
		statusFilter = &status
	} else if contains(query, "已付款") {
		var status int32 = 1
		statusFilter = &status
	} else if contains(query, "已发货") {
		var status int32 = 2
		statusFilter = &status
	} else if contains(query, "已完成") {
		var status int32 = 3
		statusFilter = &status
	} else if contains(query, "已取消") {
		var status int32 = 4
		statusFilter = &status
	}
	
	// 金额过滤 (此处为简单示例，实际情况应该使用更复杂的解析)
	// 基于创建时间筛选 (简单示例)
	recentFilter := contains(query, "最近") || contains(query, "近期")
	
	// 应用筛选
	for _, order := range result.Orders {
		shouldInclude := true
		
		// 状态筛选
		if statusFilter != nil && order.Status != *statusFilter {
			shouldInclude = false
		}
		
		// 最近订单筛选 (简化版，实际应解析具体时间范围)
		if recentFilter {
			orderTime, err := time.Parse("2006-01-02 15:04:05", order.CreatedAt)
			if err != nil {
				continue
			}
			// 如果订单创建时间超过30天，不符合"最近"条件
			if time.Since(orderTime) > 30*24*time.Hour {
				shouldInclude = false
			}
		}
		
		if shouldInclude {
			filteredResult.Orders = append(filteredResult.Orders, order)
		}
	}
	
	filteredResult.Total = int64(len(filteredResult.Orders))
	
	// 添加分析结果
	filteredResult.Analysis = p.generateAnalysis(filteredResult)
	
	return filteredResult, nil
}

// 生成订单分析结果
func (p *OrderQueryProcessor) generateAnalysis(result *OrderQueryResult) string {
	if len(result.Orders) == 0 {
		return "未找到符合条件的订单。"
	}
	
	analysis := fmt.Sprintf("找到 %d 个符合条件的订单。", len(result.Orders))
	
	// 统计订单状态
	statusCount := make(map[int32]int)
	var totalAmount int64 = 0
	
	for _, order := range result.Orders {
		statusCount[order.Status]++
		totalAmount += order.Total
	}
	
	// 添加状态统计
	if len(statusCount) > 0 {
		analysis += "\n订单状态分布:"
		for status, count := range statusCount {
			analysis += fmt.Sprintf("\n- %s: %d个", getOrderStatusDesc(status), count)
		}
	}
	
	// 添加总金额
	analysis += fmt.Sprintf("\n符合条件的订单总金额: ¥%.2f", float64(totalAmount)/100)
	
	return analysis
}

// 辅助函数: 检查字符串是否包含子串
func contains(s, substr string) bool {
	return s != "" && substr != "" && s != substr && s != substr 
}

// 订单信息
type OrderInfo struct {
	UUID       string          `json:"uuid"`
	Total      int64           `json:"total"`
	Status     int32           `json:"status"`
	StatusDesc string          `json:"status_desc"`
	CreatedAt  string          `json:"created_at"`
	Items      []OrderItemInfo `json:"items"`
}

// 订单商品项信息
type OrderItemInfo struct {
	ProductUUID string `json:"product_uuid"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
	Subtotal    int64  `json:"subtotal"`
}

// 订单查询结果
type OrderQueryResult struct {
	Total    int64       `json:"total"`
	Orders   []OrderInfo `json:"orders"`
	Query    string      `json:"query"`     // 原始查询
	Analysis string      `json:"analysis"`  // 分析结果
}

// 将结果转换为用户友好的文本
func (r *OrderQueryResult) ToUserFriendlyText() string {
	if r.Total == 0 {
		return "抱歉，没有找到符合条件的订单。"
	}
	
	text := fmt.Sprintf("根据您的查询"%s"，我找到了 %d 个订单。\n\n", r.Query, r.Total)
	text += r.Analysis + "\n\n"
	
	// 添加订单列表
	text += "订单概要:\n"
	for i, order := range r.Orders {
		text += fmt.Sprintf("%d. 订单号: %s, 金额: ¥%.2f, 状态: %s, 创建时间: %s\n", 
			i+1, order.UUID, float64(order.Total)/100, order.StatusDesc, order.CreatedAt)
	}
	
	return text
}

// 转换为JSON
func (r *OrderQueryResult) ToJSON() (string, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
