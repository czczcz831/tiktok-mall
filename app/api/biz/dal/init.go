package dal

import (
	"github.com/czczcz831/tiktok-mall/app/api/biz/dal/casbin"
	"github.com/czczcz831/tiktok-mall/app/api/biz/dal/redis"
)

func Init() {
	redis.Init()
	casbin.Init()
	// mysql.Init()
}
