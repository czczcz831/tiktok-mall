package model

const (
	TransactionStatusUnpaid = 0
	TransactionStatusPaid   = 1
	TransactionStatusCancel = -1
)

type Transaction struct {
	Base
	UserUuid  string `gorm:"type:char(36);"`
	OrderUuid string `gorm:"type:char(36);"`
	Amount    int64  `gorm:"type:bigint;"`
	Status    int32  `gorm:"type:int;"`
}
