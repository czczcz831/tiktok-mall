package dal

import (
	"github.com/czczcz831/tiktok-mall/app/cart/biz/dal/mysql"
	// "github.com/czczcz831/tiktok-mall/app/cart/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
