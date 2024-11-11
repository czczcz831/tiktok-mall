package dal

import (
	"github.com/czczcz831/tiktok-mall/app/user/biz/dal/casbin"
	"github.com/czczcz831/tiktok-mall/app/user/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
	casbin.Init()
}
