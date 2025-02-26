package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string
	Price       uint
	Quantity    int
	Category    string
}
