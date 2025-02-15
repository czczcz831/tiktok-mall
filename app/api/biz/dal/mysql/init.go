package mysql

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/czczcz831/tiktok-mall/app/api/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		hlog.Fatal(err)
	}
}
