package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"strconv"

	"myapp/model"
	"myapp/controller"
)

var Db *gorm.DB

func Init() {
	DBMS     := "mysql"
	USER     := "user"
	PASS     := "pass"
	PROTOCOL := "tcp(db:3306)"
	DBNAME   := "myapp"


	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	Db, _ = gorm.Open(DBMS, CONNECT)
	Db.LogMode(true)
}

func main() {
	Init()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339_nano} \n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!")
	})
	e.GET("/user", func(c echo.Context) error {
		user := model.User{}
		user.Id = 3
		Db.First(&user)
		return c.String(http.StatusOK, "Users, Index name = " + user.Name + strconv.FormatInt(user.Id, 10))
	})
	e.GET("/stages", controller.StagesIndex)
	e.GET("/hoges", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hoge, Inde")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

