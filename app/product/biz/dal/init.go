package dal

import (
	"github.com/czczcz831/tiktok-mall/app/product/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/product/biz/dal/redis"
	"github.com/czczcz831/tiktok-mall/app/product/biz/dal/rocketmq"
)

func Init() {
	redis.Init()
	mysql.Init()
	rocketmq.Init()
}
