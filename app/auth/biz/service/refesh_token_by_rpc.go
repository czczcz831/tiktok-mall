package service

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	auth "github.com/czczcz831/tiktok-mall/app/auth/kitex_gen/auth"
	"github.com/czczcz831/tiktok-mall/app/user/conf"
	"github.com/czczcz831/tiktok-mall/common/utils"
)

type RefeshTokenByRPCService struct {
	ctx context.Context
} // NewRefeshTokenByRPCService new RefeshTokenByRPCService
func NewRefeshTokenByRPCService(ctx context.Context) *RefeshTokenByRPCService {
	return &RefeshTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *RefeshTokenByRPCService) Run(req *auth.RefeshTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	publicKeyString := conf.GetConf().JWT.PublicSecret
	tokenExpire := conf.GetConf().JWT.TokenExpire
	refreshTokenExpire := conf.GetConf().JWT.RefreshTokenExpire
	privateKeyString := conf.GetConf().JWT.PrivateSecret

	uuid, isRt, err := utils.VerifyToken(req.RefreshToken, publicKeyString)
	if err != nil {
		return nil, err
	}
	klog.Info("isRt: ", isRt)
	if !isRt {
		return nil, errors.New("not a refresh token")
	}

	token, refreshToken, err := utils.SignToken(uuid, privateKeyString, tokenExpire, refreshTokenExpire)
	if err != nil {
		return nil, err
	}

	return &auth.DeliveryResp{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
