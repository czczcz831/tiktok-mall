package dal

import (
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/redis"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal/rocketmq"
)

func Init() {
	redis.Init()
	mysql.Init()
	rocketmq.Init()
}
