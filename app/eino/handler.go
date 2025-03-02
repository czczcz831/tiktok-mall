package main

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/eino/biz/service"
	eino "github.com/czczcz831/tiktok-mall/app/eino/kitex_gen/eino"
)

// EinoServiceImpl implements the last service interface defined in the IDL.
type EinoServiceImpl struct{}

// QueryUserOrders implements the EinoServiceImpl interface.
func (s *EinoServiceImpl) QueryUserOrders(ctx context.Context, req *eino.QueryUserOrdersReq) (resp *eino.QueryUserOrdersResp, err error) {
	resp, err = service.NewQueryUserOrdersService(ctx).Run(req)

	return resp, err
}
