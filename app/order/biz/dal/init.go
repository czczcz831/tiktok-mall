package dal

import (
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
