package dal

import (
	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/app/checkout/biz/dal/rocketmq"
)

func Init() {
	mysql.Init()
	rocketmq.Init()
}
