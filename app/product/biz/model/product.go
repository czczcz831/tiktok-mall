package model

type Product struct {
	Base
	Name        string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
	Price       int64  `gorm:"type:bigint"`
	Stock       int64  `gorm:"type:bigint"`
}
