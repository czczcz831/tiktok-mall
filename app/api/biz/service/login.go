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
	user "github.com/czczcz831/tiktok-mall/client/user/kitex_gen/user"
	userAgent "github.com/czczcz831/tiktok-mall/client/user/rpc/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *api.LoginReq) (resp *api.LoginResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	loginResp, err := userAgent.DefaultClient().Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, &packer.MyError{
			Code: packer.UNKNOWN_SERVER_ERROR,
			Err:  err,
		}
	}

	if loginResp.UserUuid == "" {
		return nil, &packer.MyError{
			Code: packer.INVALID_ACCOUNT_PASSWORD_ERROR,
			Err:  err,
		}
	}

	uuid := loginResp.UserUuid

	deliveryTokenResp, err := authAgent.DefaultClient().DeliverTokenByRPC(h.Context, &auth.DeliverTokenReq{
		UserUuid: uuid,
	})

	if err != nil {
		return nil, &packer.MyError{
			Code: packer.AUTH_DELIBER_TOKEN_ERROR,
			Err:  err,
		}
	}

	redisKey := fmt.Sprintf("token:%s", deliveryTokenResp.Token)
	_, err = redis.RedisClient.Set(context.Background(), redisKey, "", time.Duration(deliveryTokenResp.TokenExpireAfter)*time.Hour).Result()
	if err != nil {
		return nil, &packer.MyError{
			Code: packer.UNKNOWN_SERVER_ERROR,
			Err:  err,
		}
	}

	return &api.LoginResp{
		Token:        deliveryTokenResp.Token,
		RefreshToken: deliveryTokenResp.RefreshToken,
	}, nil

}
