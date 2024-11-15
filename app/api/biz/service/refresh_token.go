package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
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

	return &api.LoginResp{
		Token:        rpcResp.Token,
		RefreshToken: rpcResp.RefreshToken,
	}, nil
}
