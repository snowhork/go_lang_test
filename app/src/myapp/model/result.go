package model

import (
	"time"
)

type Result struct {
	Id int
	UserId int `validate:"nonzero"`
	StageId int `validate:"nonzero"`
	Status bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

