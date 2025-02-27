package service

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/czczcz831/tiktok-mall/app/user/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/user/biz/model"
	"github.com/czczcz831/tiktok-mall/app/user/conf"
	user "github.com/czczcz831/tiktok-mall/app/user/kitex_gen/user"
	"github.com/czczcz831/tiktok-mall/common/utils"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	cryptoPassword := utils.MD5Crypto(req.Password, conf.GetConf().MD5Secret)
	loginUser := &model.User{Email: req.Email, Password: cryptoPassword}

	if err := mysql.DB.Where(loginUser).First(&loginUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user.LoginResp{
				UserUuid: "",
			}, nil
		}
		return nil, err
	}

	return &user.LoginResp{
		UserUuid: loginUser.UUID,
	}, nil
}
