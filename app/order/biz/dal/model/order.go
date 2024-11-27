package model

// Status 0:unpaid 1:paid -1:cancelled
const (
	OrderStatusUnpaid    = 0
	OrderStatusPaid      = 1
	OrderStatusCancelled = -1
)

type Order struct {
	Base
	UserUuid    string `gorm:"type:char(36);index"`
	AddressUuid string `gorm:"type:char(36);"`
	Total       int64  `gorm:"type:bigint;not null"`
	Status      int    `gorm:"type:int;not null"`
}

type OrderItem struct {
	Base
	OrderUUID   string `gorm:"type:char(36);index"`
	ProductUuid string `gorm:"type:char(36);index"`
	Price       int64  `gorm:"type:bigint;not null"`
	Quantity    int64  `gorm:"type:bigint;not null"`
}
