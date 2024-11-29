package rocketmq

import (
	"encoding/json"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/model"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	order "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
)

func CreateOrderTxChecker(msg *golang.MessageView) golang.TransactionResolution {

	var createOrder model.Order

	var createOrderReq order.CreateOrderReq

	err := json.Unmarshal(msg.GetBody(), &createOrderReq)

	if err != nil {
		klog.Errorf("CreateOrderTxChecker unmarshal error: %v", err)
	}

	orderResp := mysql.DB.Where("uuid = ?", createOrderReq.UserUuid).First(&createOrder)
	if orderResp.Error != nil {
		return golang.ROLLBACK
	}

	return golang.COMMIT
}
