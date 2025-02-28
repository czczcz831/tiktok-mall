package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/czczcz831/tiktok-mall/app/product/biz/dal/redis"
	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
	goredis "github.com/redis/go-redis/v9"
)

type PreDecrStockService struct {
	ctx context.Context
}

// NewPreDecrStockService 创建新的 PreDecrStockService
func NewPreDecrStockService(ctx context.Context) *PreDecrStockService {
	return &PreDecrStockService{ctx: ctx}
}

// Run 执行库存预扣减逻辑
func (s *PreDecrStockService) Run(req *product.PreDecrStockReq) (resp *product.PreDecrStockResp, err error) {
	// 记录成功扣减的 OrderItem
	var deductedItems []*product.OrderItem

	// 逐个处理 OrderItem
	for _, item := range req.Items {
		err := PreStockDeduct(item, 0)
		if err != nil {
			// 如果某个 OrderItem 扣减失败，回滚之前已扣减的库存
			for _, deductedItem := range deductedItems {
				rollbackStock(deductedItem)
			}
			return nil, err
		}
		// 记录成功扣减的 OrderItem
		deductedItems = append(deductedItems, item)
	}

	return &product.PreDecrStockResp{
		Ok: true,
	}, nil
}

// rollbackStock 回滚库存
func rollbackStock(item *product.OrderItem) {
	redisKey := fmt.Sprintf("stock:%s", item.Uuid)
	redis.RedisClient.IncrBy(context.Background(), redisKey, item.Quantity)
}

// PreStockDeduct 预扣减库存
func PreStockDeduct(item *product.OrderItem, retry int) error {
	// 尝试扣减库存，如果键不存在则加锁并从数据库初始化库存
	// 如果库存不足，返回错误
	if retry >= 3 {
		return errors.New("max retry exceeded")
	}

	var luaScript = goredis.NewScript(`
        if redis.call("EXISTS", KEYS[1]) == 0 then
            return -1
        end
        local value = redis.call("Get", KEYS[1])
        if tonumber(value) - tonumber(KEYS[2]) >= 0 then
            local leftStock = redis.call("DecrBy", KEYS[1], KEYS[2])
            return leftStock
        else
            return -2
        end
    `)

	redisKey := fmt.Sprintf("stock:%s", item.Uuid)

	res, err := luaScript.Run(context.Background(), redis.RedisClient, []string{redisKey, strconv.Itoa(int(item.Quantity))}).Result()
	if err != nil {
		return err
	}

	switch res.(int64) {
	case -1:
		// 键不存在，加锁并初始化库存
		lockKey := fmt.Sprintf("lock:%s", redisKey)
		rsMutex := redis.Rs.NewMutex(lockKey)
		err = rsMutex.TryLock()
		if err != nil {
			// 其他协程已获取锁，等待后重试
			time.Sleep(time.Millisecond * (100 << retry))
			return PreStockDeduct(item, retry+1)
		}
		// 获取锁成功，从数据库初始化库存
		resp, err := NewGetProductService(context.Background()).Run(&product.GetProductReq{Uuid: item.Uuid})
		if err != nil {
			rsMutex.Unlock()
			return err
		}
		redis.RedisClient.Set(context.Background(), redisKey, resp.Product.Stock, 0)
		rsMutex.Unlock()
		// 重试扣减
		return PreStockDeduct(item, retry)
	case -2:
		return errors.New("stock not enough")
	default:
		return nil
	}
}
