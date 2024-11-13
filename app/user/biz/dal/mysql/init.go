package mysql

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/user/biz/model"
	"github.com/czczcz831/tiktok-mall/app/user/conf"

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

	DB.AutoMigrate(&model.User{})

	klog.Info("mysql init success")

	if err != nil {
		panic(err)
	}
}
