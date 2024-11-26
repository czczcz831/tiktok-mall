package rocketmq

import (
	"encoding/json"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/cart/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/cart/biz/model"
	order "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
)

func clearCartBiz(mv *golang.MessageView) error {
	//Unmarshal message

	req := &order.CreateOrderReq{}
	err := json.Unmarshal(mv.GetBody(), req)
	if err != nil {
		klog.Errorf("unmarshal message failed: %v", err)
		return err
	}

	//clear cart

	userUuid := req.UserUuid

	productIds := make([]string, 0)

	for _, item := range req.Items {
		productId := item.ProductUuid
		productIds = append(productIds, productId)
	}
	klog.Infof("productIds: %v", productIds)
	klog.Infof("userUuid: %s", userUuid)

	res := mysql.DB.Model(&model.Cart{}).Where("user_id = ?", userUuid).Where("product_id IN (?)", productIds).Delete(&model.Cart{})
	if res.Error != nil {
		klog.Errorf("delete cart failed: %v", res.Error)
		return res.Error
	}
	return nil
}
