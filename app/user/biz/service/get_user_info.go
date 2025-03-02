package service

import (
	"context"

	"errors"
	"github.com/czczcz831/tiktok-mall/app/user/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/user/biz/model"
	user "github.com/czczcz831/tiktok-mall/app/user/kitex_gen/user"
	"github.com/czczcz831/tiktok-mall/common/errno"
	"gorm.io/gorm"
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
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New(errno.ErrUserNotFound)
		}
		return nil, result.Error
	}

	return &user.GetUserInfoResp{
		Email: userIns.Email,
	}, nil
}
