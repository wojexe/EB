package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Product struct {
	Model
	Name       string
	Price      decimal.Decimal `gorm:"type:decimal(10,2);"`
	CategoryID *uint
	Category   *Category `json:"-"`
}

type Category struct {
	Model
	Name     string
	Products []Product `json:"-"`
}

type Cart struct {
	Model
	Products []Product `gorm:"many2many:cart_products;"`
}
