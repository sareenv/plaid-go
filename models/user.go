package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"uniqueIndex"`
	Items []PlaidItem
}

type PlaidItem struct {
	ItemID      string `gorm:"uniqueIndex"`
	AccessToken string
	gorm.Model
	UserID uint
	User   User
}
