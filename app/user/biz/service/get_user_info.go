package service

import (
	"context"

	"github.com/czczcz831/tiktok-mall/app/user/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/user/biz/model"
	user "github.com/czczcz831/tiktok-mall/app/user/kitex_gen/user"
)

type GetUserInfoService struct {
	ctx context.Context
} // NewGetUserInfoService new GetUserInfoService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// Run create note info
func (s *GetUserInfoService) Run(req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	// Finish your business logic.

	userIns := &model.User{}
	result := mysql.DB.Where("uuid = ?", req.UserUuid).First(userIns)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user.GetUserInfoResp{
		Email: userIns.Email,
	}, nil
}
