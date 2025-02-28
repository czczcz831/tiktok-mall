package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/product/biz/dal"
	"github.com/czczcz831/tiktok-mall/app/product/biz/dal/redis"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

func TestPreDecrStock_Run(t *testing.T) {
	ctx := context.Background()
	dal.Init()

	// 准备测试数据
	testProduct := &product.Product{
		Uuid:  "test-product-1",
		Name:  "Test Product",
		Stock: 100,
	}

	// 初始化 Redis 中的库存
	redisKey := fmt.Sprintf("stock:%s", testProduct.Uuid)
	err := redis.RedisClient.Set(ctx, redisKey, testProduct.Stock, 0).Err()
	assert.NoError(t, err)

	tests := []struct {
		name    string
		req     *product.PreDecrStockReq
		wantErr bool
	}{
		{
			name: "正常扣减库存",
			req: &product.PreDecrStockReq{
				Items: []*product.OrderItem{
					{
						Uuid:     testProduct.Uuid,
						Quantity: 10,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "库存不足",
			req: &product.PreDecrStockReq{
				Items: []*product.OrderItem{
					{
						Uuid:     testProduct.Uuid,
						Quantity: 1000,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "多个商品扣减",
			req: &product.PreDecrStockReq{
				Items: []*product.OrderItem{
					{
						Uuid:     testProduct.Uuid,
						Quantity: 5,
					},
					{
						Uuid:     "non-exist-product",
						Quantity: 1,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewPreDecrStockService(ctx)
			resp, err := s.Run(tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)

			// 验证库存是否正确扣减
			if !tt.wantErr {
				for _, item := range tt.req.Items {
					currentStock, err := redis.RedisClient.Get(ctx, fmt.Sprintf("stock:%s", item.Uuid)).Int64()
					assert.NoError(t, err)
					assert.Equal(t, testProduct.Stock-item.Quantity, currentStock)
				}
			}
		})
	}

	// 清理测试数据
	redis.RedisClient.Del(ctx, redisKey)
}
