package repositories

import "gorm.io/gorm"

type Repositories struct {
	Products   *ProductRepository
	Categories *CategoryRepository
	Carts      *CartRepository
}

func Initialize(db *gorm.DB) Repositories {
	return Repositories{
		Products:   NewProductRepository(db),
		Categories: NewCategoryRepository(db),
		Carts:      NewCartRepository(db),
	}
}
