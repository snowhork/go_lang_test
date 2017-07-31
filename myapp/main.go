package main

import (
	"net/http"
	"github.com/labstack/echo"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"myapp/models"
	"strconv"
)

var db  *gorm.DB

func Init() {
	DBMS     := "mysql"
	USER     := "user"
	PASS     := "pass"
	PROTOCOL := "tcp(db:3306)"
	DBNAME   := "myapp"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db, _ = gorm.Open(DBMS, CONNECT)
}

func main() {
	Init()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!")
	})
	e.GET("/user", func(c echo.Context) error {
		user := models.User{}
		user.Id = 3
		db.First(&user)
		return c.String(http.StatusOK, "Users, Index name = " + user.Name + strconv.FormatInt(user.Id, 10))
	})
	e.GET("/hoges", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hoge, Inde")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

