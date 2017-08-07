package myapp

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"strconv"

	"myapp/model"
	"myapp/db"
	"myapp/controller"
)



func main() {
	db.Init()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339_nano} \n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!")
	})
	e.POST("/results", controller.ResultsCreate)
	e.GET("/users", func(c echo.Context) error {
		user := model.User{}
		user.Id = 3
		db.Db.First(&user)
		return c.String(http.StatusOK, "Users, Index name = " + user.Name + strconv.FormatInt(user.Id, 10))
	})
	e.GET("/stages", controller.StagesIndex)
	e.GET("/hoges", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hoge, Inde")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

