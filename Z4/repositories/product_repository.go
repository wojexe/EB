package repositories

import (
	"store_backend/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

type GetAllProductsOptions struct {
	Page       int
	PageSize   int
	CategoryID uint
	MinPrice   float64
	MaxPrice   float64
}

func DefaultGetAllProductsOptions() *GetAllProductsOptions {
	return &GetAllProductsOptions{
		Page:       1,
		PageSize:   25,
		CategoryID: 0,
		MinPrice:   0,
		MaxPrice:   0,
	}
}

func (r ProductRepository) GetAll(opts *GetAllProductsOptions) ([]models.Product, error) {
	var products []models.Product

	if opts == nil {
		opts = DefaultGetAllProductsOptions()
	}

	err := r.db.Scopes(
		WithCategory(),
		Paginate(opts.Page, opts.PageSize),
		ByCategory(opts.CategoryID),
		PriceRange(opts.MinPrice, opts.MaxPrice),
		OrderBy("created_at", "desc"),
	).Find(&products).Error

	return products, err
}

func (r ProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r ProductRepository) Create(product *models.Product) (*models.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r ProductRepository) Update(product *models.Product) (*models.Product, error) {
	if err := r.db.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r ProductRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Scopes

func WithCategory() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Category")
	}
}

func ByCategory(categoryID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if categoryID > 0 {
			return db.Where("category_id = ?", categoryID)
		}
		return db
	}
}

func PriceRange(min, max float64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if min > 0 {
			db = db.Where("price >= ?", min)
		}
		if max > 0 {
			db = db.Where("price <= ?", max)
		}
		return db
	}
}

func InStock() func(db *gorm.DB) *gorm.DB {
	// not implemented in DB
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("stock_quantity > 0")
	}
}
