package main

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/auth/biz/service"
	auth "github.com/czczcz831/tiktok-mall/app/auth/kitex_gen/auth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	resp, err = service.NewDeliverTokenByRPCService(ctx).Run(req)

	return resp, err
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	resp, err = service.NewVerifyTokenByRPCService(ctx).Run(req)

	return resp, err
}
