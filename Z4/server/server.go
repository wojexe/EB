package server

import (
	"fmt"
	"log"
	"store_backend/handlers"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func Initialize(handlers []handlers.Handler) Server {
	e := echo.New()

	e.HideBanner = true
	e.Validator = &customValidator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}

	configureMiddleware(e)

	for _, handler := range handlers {
		handler.RegisterRoutes(e)
	}

	return Server{echo: e}
}

func (s Server) Start() {
	log.Printf("Available routes:\n%s", printRoutes(s.echo.Routes()))

	s.echo.Logger.Fatal(s.echo.Start(":1323"))
}

func configureMiddleware(e *echo.Echo) {
	e.Pre(middleware.AddTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
}

func printRoutes(routes []*echo.Route) string {
	formatted := make([]string, len(routes))
	for i, route := range routes {
		formatted[i] = formatRoute(route)
	}
	return strings.Join(formatted, "\n")
}

func formatRoute(r *echo.Route) string {
	return fmt.Sprintf("%s\t%s", r.Method, r.Path)
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
