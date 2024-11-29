package dal

import (
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/redis"
	"github.com/czczcz831/tiktok-mall/app/order/biz/dal/rocketmq"
)

func Init() {
	redis.Init()
	mysql.Init()
	rocketmq.Init()
}
