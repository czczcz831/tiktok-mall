package model

import "gorm.io/gorm"

type Base struct {
	gorm.Model
	UUID string `gorm:"type:char(36);uniqueIndex"`
}
