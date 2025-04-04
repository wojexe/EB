package database

import (
	"math/rand"
	"store_backend/models"

	"github.com/go-faker/faker/v4"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Seed populates the database with example data
func Seed(db *gorm.DB) error {
	// Clear existing data
	db.Exec("DELETE FROM cart_products")
	db.Exec("DELETE FROM products")
	db.Exec("DELETE FROM categories")
	db.Exec("DELETE FROM carts")

	// Create categories
	categories := []models.Category{
		{Name: "Electronics"},
		{Name: "Clothing"},
		{Name: "Books"},
		{Name: "Home & Kitchen"},
		{Name: "Sports & Outdoors"},
	}

	for i := range categories {
		if err := db.Create(&categories[i]).Error; err != nil {
			return err
		}
	}

	// Create products
	products := make([]models.Product, 0)
	for _, category := range categories {
		// Create 5-10 products per category
		numProducts := rand.Intn(6) + 5
		for i := 0; i < numProducts; i++ {
			// Create price between $5.99 and $999.99
			price := decimal.NewFromFloat(float64(rand.Intn(99400)+599) / 100)

			product := models.Product{
				Name:       faker.Word() + " " + faker.Word(),
				Price:      price,
				CategoryID: &category.ID,
			}
			if err := db.Create(&product).Error; err != nil {
				return err
			}
			products = append(products, product)
		}
	}

	// Create carts
	for i := 0; i < 5; i++ {
		cart := models.Cart{}
		if err := db.Create(&cart).Error; err != nil {
			return err
		}

		// Add random products to cart (1-5 products)
		numProductsInCart := rand.Intn(5) + 1
		cartProducts := make([]models.Product, 0)

		for j := 0; j < numProductsInCart; j++ {
			randomProduct := products[rand.Intn(len(products))]
			cartProducts = append(cartProducts, randomProduct)
		}

		if err := db.Model(&cart).Association("Products").Append(cartProducts); err != nil {
			return err
		}
	}

	return nil
}

// GetRandomProducts returns a specified number of random products
func GetRandomProducts(db *gorm.DB, count int) ([]models.Product, error) {
	var products []models.Product
	err := db.Order("RANDOM()").Limit(count).Find(&products).Error
	return products, err
}
