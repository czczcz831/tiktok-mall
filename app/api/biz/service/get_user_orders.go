package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	order "github.com/czczcz831/tiktok-mall/client/order/kitex_gen/order"
	orderAgent "github.com/czczcz831/tiktok-mall/client/order/rpc/order"
)

type GetUserOrdersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserOrdersService(Context context.Context, RequestContext *app.RequestContext) *GetUserOrdersService {
	return &GetUserOrdersService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserOrdersService) Run(req *api.GetUserOrdersReq) (resp *api.GetUserOrdersResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userUUID, ok := h.RequestContext.Get("uuid")
	if !ok {
		return nil, &packer.MyError{
			Code: packer.UNKNOWN_SERVER_ERROR,
		}
	}

	userUUIDStr, _ := userUUID.(string)

	getUserOrdersResp, err := orderAgent.GetUserOrders(h.Context, &order.GetUserOrdersReq{
		UserUuid: userUUIDStr,
	})

	if err != nil {
		return nil, err
	}

	orders := make([]*api.Order, 0, len(getUserOrdersResp.Orders))
	for _, order := range getUserOrdersResp.Orders {
		items := make([]*api.OrderItemWithPrice, 0, len(order.Items))
		for _, item := range order.Items {
			items = append(items, &api.OrderItemWithPrice{
				ProductUUID: item.ProductUuid,
				Quantity:    item.Quantity,
				Price:       item.Price,
			})
		}
		orders = append(orders, &api.Order{
			UUID:        order.Uuid,
			UserUUID:    order.UserUuid,
			AddressUUID: order.AddressUuid,
			Total:       order.Total,
			Status:      order.Status,
			CreatedAt:   order.CreatedAt,
			Items:       items,
		})
	}

	return &api.GetUserOrdersResp{
		Total:  getUserOrdersResp.Total,
		Orders: orders,
	}, nil
}
