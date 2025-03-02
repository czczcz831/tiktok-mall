package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/product/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/product/biz/dal/redis"
	"github.com/czczcz831/tiktok-mall/app/product/biz/model"

	product "github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product"
	"github.com/czczcz831/tiktok-mall/common/errno"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	// 查之前的库存数
	oldProduct := &model.Product{}
	result := mysql.DB.Where("uuid = ?", req.Product.Uuid).First(oldProduct)
	if result.Error != nil {
		return nil, errors.New(errno.ErrProductNotFound)
	}

	oldStock := oldProduct.Stock

	// 更新产品信息到MySQL
	dbProduct := &model.Product{}
	updateResp := mysql.DB.Model(dbProduct).Where("uuid = ?", req.Product.Uuid).Updates(
		&model.Product{
			Name:        req.Product.Name,
			Description: req.Product.Description,
			Price:       req.Product.Price,
			Stock:       req.Product.Stock,
		})

	if updateResp.Error != nil {
		return nil, updateResp.Error
	}

	if updateResp.RowsAffected == 0 {
		return nil, errors.New(errno.ErrDatabaseConsistent)
	}

	//计算库存diff
	stockDiff := req.Product.Stock - oldStock

	//同步更新Redis中的库存
	if stockDiff != 0 {
		redisKey := fmt.Sprintf("stock:%s", req.Product.Uuid)

		// 检查Redis中是否已存在该库存键
		exists, err := redis.RedisClient.Exists(s.ctx, redisKey).Result()
		if err != nil {
			klog.Errorf("Failed to check Redis key existence: %v", err)
		} else if exists == 1 {
			// Redis中已存在库存记录，需要更新
			var redisOp error
			_, redisOp = redis.RedisClient.IncrBy(s.ctx, redisKey, stockDiff).Result()
			if redisOp != nil {
				klog.Errorf("Failed to update Redis stock: %v", redisOp)
			}
		}
	}

	return &product.UpdateProductResp{
		Product: &product.Product{
			Uuid:        dbProduct.UUID,
			Name:        dbProduct.Name,
			Description: dbProduct.Description,
			Price:       dbProduct.Price,
			Stock:       dbProduct.Stock,
		},
	}, nil
}
