package handlers

import (
	"errors"
	"log"
	"net/http"
	"store_backend/models"
	"store_backend/repositories"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProductsHandler struct {
	repos repositories.Repositories
}

func (h *ProductsHandler) RegisterRoutes(e *echo.Echo) error {
	products := e.Group("/products")

	products.GET("", h.GetProducts)
	products.GET("/:id", h.GetProduct)
	products.POST("", h.CreateProduct)
	products.PUT("/:id", h.UpdateProduct)
	products.DELETE("/:id", h.DeleteProduct)

	return nil
}

func (h *ProductsHandler) GetProducts(c echo.Context) error {
	products, err := h.repos.Products.GetAll(nil)

	if err != nil {
		log.Printf("error getting products: %v", err)
		return c.NoContent(501)
	}

	return c.JSON(http.StatusOK, products)
}

type GetProductRequest struct {
	ID uint `param:"id" validate:"required"`
}

func (h *ProductsHandler) GetProduct(c echo.Context) error {
	req := GetProductRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	product, err := h.repos.Products.GetByID(req.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.NoContent(http.StatusNotFound)
		} else {
			log.Printf("error getting product: %v", err)
			return c.NoContent(500)
		}
	}

	return c.JSON(http.StatusOK, product)
}

type CreateProductRequest struct {
	Name       string          `json:"name" validate:"required"`
	Price      decimal.Decimal `json:"price" validate:"required"`
	CategoryID *uint           `json:"categoryId"`
}

func (h *ProductsHandler) CreateProduct(c echo.Context) error {
	data := CreateProductRequest{}

	if err := bindAndValidate(c, &data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	p := models.Product{Name: data.Name, Price: data.Price, CategoryID: data.CategoryID}
	product, err := h.repos.Products.Create(&p)
	if err != nil {
		log.Printf("error creating product: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, product)
}

type UpdateProductRequest struct {
	ID         uint            `param:"id" validate:"required"`
	Name       string          `json:"name" validate:"required"`
	Price      decimal.Decimal `json:"price" validate:"required"`
	CategoryID *uint           `json:"categoryId"`
}

func (h *ProductsHandler) UpdateProduct(c echo.Context) error {
	data := UpdateProductRequest{}

	if err := bindAndValidate(c, &data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	product, err := h.repos.Products.GetByID(data.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.NoContent(http.StatusNotFound)
		} else {
			log.Printf("error getting product: %v", err)
			return c.NoContent(500)
		}
	}

	product.Name = data.Name
	product.Price = data.Price
	product.CategoryID = data.CategoryID

	product, err = h.repos.Products.Update(product)
	if err != nil {
		log.Printf("error updating product: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update product",
		})
	}

	return c.JSON(http.StatusOK, product)
}

type DeleteProductRequest struct {
	ID uint `param:"id"`
}

func (h *ProductsHandler) DeleteProduct(c echo.Context) error {
	data := DeleteProductRequest{}

	if err := bindAndValidate(c, &data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Check if product exists
	_, err := h.repos.Products.GetByID(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Product not found",
			})
		}

		log.Printf("error getting product for deletion: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete product",
		})
	}

	err = h.repos.Products.Delete(data.ID)

	if err != nil {
		log.Printf("error deleting product: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete product",
		})
	}

	return c.NoContent(http.StatusNoContent)
}
