package handlers

import (
	"log"
	"store_backend/repositories"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	RegisterRoutes(e *echo.Echo) error
}

func Initialize(repos repositories.Repositories) []Handler {
	return []Handler{
		&ProductsHandler{repos: repos},
		&CategoriesHandler{repos: repos},
		&CartHandler{repos: repos},
	}
}

func bindAndValidate(c echo.Context, model interface{}) map[string]string {
	if err := c.Bind(model); err != nil {
		log.Printf("Binding error: %v", err)

		return map[string]string{"error": "Invalid request format"}
	}

	if err := c.Validate(model); err != nil {
		log.Printf("Validation error: %v", err)

		return map[string]string{"error": "Invalid request format"}
	}

	return nil
}
