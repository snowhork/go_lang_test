package model

import "time"

type Stage struct {
	Id int
	Name string `sql:"size:255"`
	UserId int
	Csv string
	CreatedAt time.Time
}
