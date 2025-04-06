package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Product struct {
	Model
	Name       string          `json:"name"`
	Price      decimal.Decimal `json:"price" gorm:"type:decimal(10,2);"`
	CategoryID *uint           `json:"categoryId"`
	Category   *Category       `json:"-"`
}

type Category struct {
	Model
	Name     string    `json:"name"`
	Products []Product `json:"-"`
}

type Cart struct {
	Model
	Products []Product `gorm:"many2many:cart_products;"`
}
