package repositories

import (
	"store_backend/models"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r CartRepository) GetAll() ([]models.Cart, error) {
	var carts []models.Cart
	if err := r.db.Preload("Products").Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r CartRepository) GetByID(id uint) (*models.Cart, error) {
	var cart models.Cart

	if err := r.db.Scopes(
		WithProducts(),
	).First(&cart, id).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r CartRepository) Create(cart *models.Cart) (*models.Cart, error) {
	if err := r.db.Create(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (r CartRepository) Update(cart *models.Cart) (*models.Cart, error) {
	if err := r.db.Save(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (r CartRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Cart{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r CartRepository) AddProduct(cartID uint, productID uint) error {
	var cart models.Cart
	var product models.Product

	if err := r.db.First(&cart, cartID).Error; err != nil {
		return err
	}

	if err := r.db.First(&product, productID).Error; err != nil {
		return err
	}

	cart.Products = append(cart.Products, product)

	if err := r.db.Save(&cart).Error; err != nil {
		return err
	}

	return nil
}

func (r CartRepository) RemoveProduct(cartID uint, productID uint) error {
	var cart models.Cart

	if err := r.db.Scopes(
		WithProducts(),
	).First(&cart, cartID).Error; err != nil {
		return err
	}

	for i, product := range cart.Products {
		if product.ID == productID {
			cart.Products = append(cart.Products[:i], cart.Products[i+1:]...)
			break
		}
	}

	if err := r.db.Save(&cart).Error; err != nil {
		return err
	}

	return nil
}

func (r CartRepository) GetProducts(cartID uint) ([]models.Product, error) {
	var cart models.Cart

	if err := r.db.Scopes(
		WithProducts(),
	).First(&cart, cartID).Error; err != nil {
		return nil, err
	}

	return cart.Products, nil
}

func (r CartRepository) ClearCart(cartID uint) error {
	var cart models.Cart

	if err := r.db.First(&cart, cartID).Error; err != nil {
		return err
	}

	cart.Products = []models.Product{}

	if err := r.db.Save(&cart).Error; err != nil {
		return err
	}

	return nil
}

// Scopes

func WithProducts() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Products")
	}
}
