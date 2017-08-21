package model

type User struct {
	Id int64
	Name string `sql:"size:255"`
}

