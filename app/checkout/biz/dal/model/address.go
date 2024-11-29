package model

type Address struct {
	Base
	UserUUID      string `gorm:"type:char(36);not null" json:"user_uuid"`
	StreetAddress string `gorm:"type:varchar(255);not null" json:"street_address"`
	City          string `gorm:"type:varchar(100);not null" json:"city"`
	State         string `gorm:"type:varchar(100);not null" json:"state"`
	Country       string `gorm:"type:varchar(100);not null" json:"country"`
	ZipCode       int64  `gorm:"type:bigint;not null" json:"zip_code"`
}
