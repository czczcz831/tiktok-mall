package service

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/dal/redis"
	"github.com/czczcz831/tiktok-mall/app/api/biz/utils/packer"
	api "github.com/czczcz831/tiktok-mall/app/api/hertz_gen/api"
	auth "github.com/czczcz831/tiktok-mall/client/auth/kitex_gen/auth"
	authAgent "github.com/czczcz831/tiktok-mall/client/auth/rpc/auth"
)

type RefreshTokenService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRefreshTokenService(Context context.Context, RequestContext *app.RequestContext) *RefreshTokenService {
	return &RefreshTokenService{RequestContext: RequestContext, Context: Context}
}

func (h *RefreshTokenService) Run(req *api.RefreshTokenReq) (resp *api.LoginResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	rpcResp, err := authAgent.DefaultClient().RefeshTokenByRPC(
		h.Context,
		&auth.RefeshTokenReq{
			RefreshToken: req.RefreshToken,
		},
	)

	if err != nil {
		return nil, err
	}

	redisKey := fmt.Sprintf("token:%s", rpcResp.Token)
	_, err = redis.RedisClient.Set(context.Background(), redisKey, "", time.Duration(rpcResp.TokenExpireAfter)*time.Hour).Result()
	if err != nil {
		return nil, &packer.MyError{
			Code: packer.UNKNOWN_SERVER_ERROR,
			Err:  err,
		}
	}

	return &api.LoginResp{
		Token:        rpcResp.Token,
		RefreshToken: rpcResp.RefreshToken,
	}, nil
}
