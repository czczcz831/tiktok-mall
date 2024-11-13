package dal

import (
	"github.com/czczcz831/tiktok-mall/app/api/biz/dal/redis"
)

func Init() {
	redis.Init()
	// mysql.Init()
}
