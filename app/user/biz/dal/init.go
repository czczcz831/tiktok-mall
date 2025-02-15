package dal

import (
	"github.com/czczcz831/tiktok-mall/app/user/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
