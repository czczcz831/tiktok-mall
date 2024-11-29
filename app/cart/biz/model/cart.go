package model

type Cart struct {
	Base
	UserID    string `gorm:"type:string"`
	ProductID string `gorm:"type:string"`
	Quantity  uint   `gorm:"type:int"`
}
