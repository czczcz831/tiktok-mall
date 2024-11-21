package model

type CreditCard struct {
	Base
	UserUUID           string `gorm:"type:char(36);not null" json:"user_uuid"`
	CreditCardNumber   string `gorm:"type:varchar(255);not null" json:"credit_card_number"`
	CreditCardCVV      int64  `gorm:"type:bigint;not null" json:"credit_card_cvv"`
	CreditCardExpMonth int64  `gorm:"type:bigint;not null" json:"credit_card_exp_month"`
	CreditCardExpYear  int64  `gorm:"type:bigint;not null" json:"credit_card_exp_year"`
}
