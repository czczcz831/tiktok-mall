package model

type Checkout struct {
	Base
	UserUUID  string `gorm:"type:char(36);not null" json:"user_uuid"`
	FirstName string `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName  string `gorm:"type:varchar(255);not null" json:"last_name"`
	Email     string `gorm:"type:varchar(255);not null" json:"email"`

	AddressUUID    string `gorm:"type:char(36);not null" json:"address_uuid"`
}
