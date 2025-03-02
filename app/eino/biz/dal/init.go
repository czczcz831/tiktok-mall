package dal

import (
	"github.com/czczcz831/tiktok-mall/app/eino/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/eino/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
