package model

type User struct {
	Base
	Email    string `gorm:"type:varchar(255);uniqueIndex"`
	Password string `gorm:"type:varchar(255)"`
}
