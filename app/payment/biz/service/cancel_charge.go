package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/mysql"
	payment "github.com/czczcz831/tiktok-mall/app/payment/kitex_gen/payment"
)

type CancelChargeService struct {
	ctx context.Context
} // NewCancelChargeService new CancelChargeService
func NewCancelChargeService(ctx context.Context) *CancelChargeService {
	return &CancelChargeService{ctx: ctx}
}

// Run create note info
func (s *CancelChargeService) Run(req *payment.CancelChargeReq) (resp *payment.CancelChargeResp, err error) {
	// Finish your business logic.

	res := mysql.DB.Model(&model.Transaction{}).Where("uuid = ?", req.TransactionUuid).Where("user_uuid = ? ", req.UserUuid).Where("status = ?", model.TransactionStatusUnpaid).Update("status", model.TransactionStatusCancel)
	if res.Error != nil {
		klog.Errorf("cancel payment failed: %v", res.Error)
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return &payment.CancelChargeResp{
			Success: false,
		}, nil
	}

	return &payment.CancelChargeResp{
		Success: true,
	}, nil

}
