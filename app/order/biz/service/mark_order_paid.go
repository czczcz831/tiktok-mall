package service

import (
	"context"
	"errors"

	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	order "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// Finish your business logic.

	updateRes := mysql.DB.Model(&model.Order{}).Where("uuid = ?", req.Uuid).Update("status", model.OrderStatusPaid)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	if updateRes.RowsAffected == 0 {
		return nil, errors.New("order not found")
	}

	return
}
