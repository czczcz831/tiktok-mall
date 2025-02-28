package mysql

import (
	"github.com/czczcz831/tiktok-mall/app/order/biz/model"
	"github.com/czczcz831/tiktok-mall/app/order/conf"

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
		panic(err)
	}

	err = DB.AutoMigrate(&model.Order{}, &model.OrderItem{})

	if err != nil {
		panic(err)
	}
}
