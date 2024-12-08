package main

import (
	"net/http"
	"pictureframe/environment"
	"pictureframe/routes"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Set headers to disable caching for static files in dev
			if environment.IsDevelopment() {
				c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
				c.Response().Header().Set("Pragma", "no-cache")
				c.Response().Header().Set("Expires", "0")
				return next(c)
			}
			if c.Request().URL.String() == "/random.txt/image" {
				c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
				c.Response().Header().Set("Pragma", "no-cache")
				c.Response().Header().Set("Expires", "0")
				return next(c)
			}

			return next(c)
		}
	})
	e.Static("/", "static")

	if environment.IsDevelopment() {
		e.Use(
			func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					if c.Request().URL.Path == "/css/style.css" {
						c.Response().Header().Set("Cache-Control", "no-store")
						c.Response().Header().Set("Pragma", "no-cache")
						c.Response().Header().Set("Expires", "0")
					}
					return next(c)
				}
			},
		)
	}
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} ${method} ${uri} ${status}
`,
	}))
	e.Use(middleware.Recover())

	routes.Init(e)

	e.Logger.Fatal(e.Start(":" + environment.EnvVar("PORT")))
}
