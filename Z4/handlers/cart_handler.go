package handlers

import (
	"errors"
	"log"
	"net/http"
	"store_backend/models"
	"store_backend/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CartHandler struct {
	repos repositories.Repositories
}

func (h *CartHandler) RegisterRoutes(e *echo.Echo) error {
	carts := e.Group("/carts")

	carts.GET("", h.GetCarts)
	carts.GET("/:id", h.GetCart)
	carts.POST("", h.CreateCart)
	carts.DELETE("/:id", h.DeleteCart)

	cart_products := carts.Group("/:id/products")
	cart_products.GET("", h.GetCartProducts)
	cart_products.POST("/:productId", h.AddProductToCart)
	cart_products.DELETE("/:productId", h.RemoveProductFromCart)
	cart_products.DELETE("", h.ClearCart)

	return nil
}

func (h *CartHandler) GetCarts(c echo.Context) error {
	carts, err := h.repos.Carts.GetAll()

	if err != nil {
		log.Printf("error getting carts: %v", err)
		return c.NoContent(501)
	}

	return c.JSON(http.StatusOK, carts)
}

type GetCartRequest struct {
	ID uint `param:"id" validate:"required"`
}

func (h *CartHandler) GetCart(c echo.Context) error {
	req := GetCartRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	cart, err := h.repos.Carts.GetByID(req.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.NoContent(http.StatusNotFound)
		} else {
			log.Printf("error getting cart: %v", err)
			return c.NoContent(500)
		}
	}

	return c.JSON(http.StatusOK, cart)
}

func (h *CartHandler) CreateCart(c echo.Context) error {
	cart := &models.Cart{}

	newCart, err := h.repos.Carts.Create(cart)
	if err != nil {
		log.Printf("error creating cart: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create cart",
		})
	}

	return c.JSON(http.StatusCreated, newCart)
}

type DeleteCartRequest struct {
	ID uint `param:"id" validate:"required"`
}

func (h *CartHandler) DeleteCart(c echo.Context) error {
	data := DeleteCartRequest{}

	if err := bindAndValidate(c, &data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Check if cart exists
	_, err := h.repos.Carts.GetByID(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Cart not found",
			})
		}

		log.Printf("error getting cart for deletion: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete cart",
		})
	}

	err = h.repos.Carts.Delete(data.ID)

	if err != nil {
		log.Printf("error deleting cart: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete cart",
		})
	}

	return c.NoContent(http.StatusNoContent)
}

type AddProductToCartRequest struct {
	ID        uint `param:"id" validate:"required"`
	ProductID uint `param:"productId" validate:"required"`
}

func (h *CartHandler) AddProductToCart(c echo.Context) error {
	data := AddProductToCartRequest{}

	if err := bindAndValidate(c, &data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Check if cart exists
	_, err := h.repos.Carts.GetByID(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Cart not found",
			})
		}

		log.Printf("error getting cart: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to add product to cart",
		})
	}

	// Check if product exists
	_, err = h.repos.Products.GetByID(data.ProductID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Product not found",
			})
		}

		log.Printf("error getting product: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to add product to cart",
		})
	}

	err = h.repos.Carts.AddProduct(data.ID, data.ProductID)
	if err != nil {
		log.Printf("error adding product to cart: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to add product to cart",
		})
	}

	cart, err := h.repos.Carts.GetByID(data.ID)
	if err != nil {
		log.Printf("error getting updated cart: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Product added, but failed to retrieve updated cart",
		})
	}

	return c.JSON(http.StatusOK, cart)
}

type RemoveProductFromCartRequest struct {
	ID        uint `param:"id" validate:"required"`
	ProductID uint `param:"productId" validate:"required"`
}

func (h *CartHandler) RemoveProductFromCart(c echo.Context) error {
	data := RemoveProductFromCartRequest{}

	if err := bindAndValidate(c, &data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Check if cart exists
	_, err := h.repos.Carts.GetByID(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Cart not found",
			})
		}

		log.Printf("error getting cart: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to remove product from cart",
		})
	}

	err = h.repos.Carts.RemoveProduct(data.ID, data.ProductID)
	if err != nil {
		log.Printf("error removing product from cart: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to remove product from cart",
		})
	}

	cart, err := h.repos.Carts.GetByID(data.ID)
	if err != nil {
		log.Printf("error getting updated cart: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Product removed, but failed to retrieve updated cart",
		})
	}

	return c.JSON(http.StatusOK, cart)
}

type GetCartProductsRequest struct {
	ID uint `param:"id" validate:"required"`
}

func (h *CartHandler) GetCartProducts(c echo.Context) error {
	data := GetCartProductsRequest{}

	if err := bindAndValidate(c, &data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	products, err := h.repos.Carts.GetProducts(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Cart not found",
			})
		}

		log.Printf("error getting cart products: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get cart products",
		})
	}

	return c.JSON(http.StatusOK, products)
}

type ClearCartRequest struct {
	ID uint `param:"id" validate:"required"`
}

func (h *CartHandler) ClearCart(c echo.Context) error {
	data := ClearCartRequest{}

	if err := bindAndValidate(c, &data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := h.repos.Carts.ClearCart(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Cart not found",
			})
		}

		log.Printf("error clearing cart: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to clear cart",
		})
	}

	cart, err := h.repos.Carts.GetByID(data.ID)
	if err != nil {
		log.Printf("error getting updated cart: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Cart cleared, but failed to retrieve updated cart",
		})
	}

	return c.JSON(http.StatusOK, cart)
}
