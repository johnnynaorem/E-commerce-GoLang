package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Email    string `gorm:"uniqueIndex"`
	Phone    string `gorm:"uniqueIndex"`
	Address  string
	City     string
	Role     string
}
