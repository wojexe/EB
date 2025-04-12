package server

import (
	"fmt"
	"log"
	"store_backend/environment"
	"store_backend/handlers"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

type Server struct {
	echo *echo.Echo
}

func Initialize(handlers []handlers.Handler, env environment.Environment) Server {
	e := echo.New()

	e.HideBanner = true
	e.Validator = &customValidator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}

	configureMiddleware(e, env)

	for _, handler := range handlers {
		handler.RegisterRoutes(e)
	}

	return Server{echo: e}
}

func (s Server) Start() {
	log.Printf("Available routes:\n%s", printRoutes(s.echo.Routes()))

	s.echo.Logger.Fatal(s.echo.Start(":1323"))
}

func configureMiddleware(e *echo.Echo, env environment.Environment) {
	slogEcho := slogecho.New(env.Logger)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(slogEcho)
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "*"},
		AllowCredentials: true,
	}))
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
