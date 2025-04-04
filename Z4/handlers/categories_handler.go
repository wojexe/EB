package handlers

import (
	"store_backend/repositories"

	"github.com/labstack/echo/v4"
)

type CategoriesHandler struct {
	repos repositories.Repositories
}

func (h *CategoriesHandler) RegisterRoutes(e *echo.Echo) error {
	return nil
}
