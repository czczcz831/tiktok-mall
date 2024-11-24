package model

type Order struct {
	Base
	UserUuid string `gorm:"type:char(36);index"`
	Total    int64  `gorm:"type:bigint;not null"`
	IsPaid   bool   `gorm:"type:tinyint(1);not null"`
}

type OrderItem struct {
	Base
	OrderUUID   string `gorm:"type:char(36);index"`
	ProductUuid string `gorm:"type:char(36);index"`
	Price       int64  `gorm:"type:bigint;not null"`
	Quantity    int64  `gorm:"type:bigint;not null"`
}
