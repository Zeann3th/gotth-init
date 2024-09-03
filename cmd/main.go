package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zeann3th/htmx+tailwind/internal/templates"
)

var userCount = 430000

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return templates.Index().Render(context.Background(), c.Response().Writer)
	})
	e.Static("/assets", "/internal/assets")
	e.Static("/css", "/internal/css")
	e.Logger.Fatal(e.Start(":4000"))
}
