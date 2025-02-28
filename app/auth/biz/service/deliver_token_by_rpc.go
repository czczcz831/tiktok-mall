package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	auth "github.com/czczcz831/tiktok-mall/app/auth/kitex_gen/auth"
	"github.com/czczcz831/tiktok-mall/app/user/conf"
	"github.com/czczcz831/tiktok-mall/common/utils"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	tokenExpire := conf.GetConf().JWT.TokenExpire
	refreshTokenExpire := conf.GetConf().JWT.RefreshTokenExpire
	privateKeyString := conf.GetConf().JWT.PrivateSecret

	token, refreshToken, err := utils.SignToken(req.UserUuid, privateKeyString, tokenExpire, refreshTokenExpire)

	if err != nil {
		klog.Fatalf("SignToken failed: %v", err)
		return nil, err
	}

	return &auth.DeliveryResp{
		Token:                   token,
		RefreshToken:            refreshToken,
		TokenExpireAfter:        int64(tokenExpire),
		RefreshTokenExpireAfter: int64(refreshTokenExpire),
	}, nil
}
