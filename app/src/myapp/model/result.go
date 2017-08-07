package model

import "time"

type Result struct {
	Id int64
	UserId int64
	Status bool
	CreatedAt time.Time
}


