package main

import (
	"net/http"

	"goth-project/templates"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return templates.Index().Render(c.Request().Context(), c.Response().Writer)
	})

	e.POST("/api/greet", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<p class='text-green-600 font-bold'>Hello from the GOTH stack!</p>")
	})

	e.Logger.Fatal(e.Start(":8080"))
}