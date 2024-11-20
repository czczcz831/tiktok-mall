package mysql

import (
	"github.com/czczcz831/tiktok-mall/app/product/biz/model"
	"github.com/czczcz831/tiktok-mall/app/product/conf"

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

	DB.AutoMigrate(&model.Product{})

	if err != nil {
		panic(err)
	}
}
