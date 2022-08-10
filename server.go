package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/", func(c echo.Context) error {
		user := []map[string]interface{}{
			{
				"name": "Harry Potter",
				"city": "London",
			},
			{
				"name": "Don Quixote",
				"city": "Madrid",
			},
			{
				"name": "Joan of Arc",
				"city": "Paris",
			},
			{
				"name": "Rosa Park",
				"city": "Alabama",
			},
		}

		return c.JSON(http.StatusOK, user)

	})

	e.Logger.Fatal(e.Start(":3000"))
}
