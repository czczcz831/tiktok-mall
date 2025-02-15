package dal

import (
	"github.com/czczcz831/tiktok-mall/app/cart/biz/dal/rocketmq"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
	rocketmq.Init()
}
