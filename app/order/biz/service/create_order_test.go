package service

import (
	"context"
	"testing"

	"github.com/czczcz831/tiktok-mall/app/order/biz/dal"
	order "github.com/czczcz831/tiktok-mall/app/order/kitex_gen/order"
	_ "github.com/joho/godotenv/autoload"
)

func TestCreateOrder_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewCreateOrderService(ctx)
	// init req and assert value

	req := &order.CreateOrderReq{
		UserUuid: "czczcz-test-uuid",
		Total:    100,
		Items: []*order.OrderItem{
			{
				ProductUuid: "czczcz-test-product-uuid",
				Price:       100,
				Quantity:    1,
			},
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
