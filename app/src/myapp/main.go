package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"myapp/db"
	"myapp/controller"
)



func main() {
	db.Init()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339_nano} \n",
	}))

	e.Logger.SetLevel(log.DEBUG)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!")
	})
	e.POST("/results", controller.ResultsCreate)
	e.PATCH("/results/:id", controller.ResultsUpdate)
	e.GET("/stages", controller.StagesIndex)
	e.GET("/stages/:id", controller.StagesShow)
	e.POST("/stages", controller.StagesCreate)
	e.DELETE("/stages/:id", controller.StagesDelete)
	e.Logger.Fatal(e.Start(":1323"))
}

