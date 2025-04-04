package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string
	Price      decimal.Decimal `gorm:"type:decimal(10,2);"`
	CategoryID *uint
	Category   *Category
}

type Category struct {
	gorm.Model
	Name     string
	Products []Product
}

type Cart struct {
	gorm.Model
	Products []Product `gorm:"many2many:cart_products;"`
}
