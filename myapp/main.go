package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!")
	})
	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "Users, Index")
	})
	e.GET("/hoges", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hoge, Index")
	})
	e.Logger.Fatal(e.Start(":1323"))
}