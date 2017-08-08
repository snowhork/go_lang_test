package model

import "time"

type Stage struct {
	Id int64
	Name string `sql:"size:255"`
	UserId int64
	Csv string
	CreatedAt time.Time
}
