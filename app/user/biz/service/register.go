package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/czczcz831/tiktok-mall/app/user/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/user/biz/model"
	"github.com/czczcz831/tiktok-mall/app/user/conf"
	"github.com/czczcz831/tiktok-mall/common/utils"

	user "github.com/czczcz831/tiktok-mall/app/user/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.

	nodeId := conf.GetConf().NodeID
	node, err := snowflake.NewNode(nodeId % 1024)
	if err != nil {
		return nil, err
	}
	uuid := node.Generate().String()

	newUser := &model.User{
		Base:     model.Base{UUID: uuid},
		Email:    req.Email,
		Password: (utils.MD5Crypto(req.Password, conf.GetConf().MD5Secret)),
	}

	// 保存用户对象到数据库
	if err := mysql.DB.Create(newUser).Error; err != nil {
		return nil, err
	}

	// 返回响应
	return &user.RegisterResp{
		UserUuid: uuid,
	}, nil
}
