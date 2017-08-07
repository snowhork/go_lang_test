package db

import "github.com/jinzhu/gorm"

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
