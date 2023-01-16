package models

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	UserId   int     `json:"userId" validate:"nonzero"`
	Distance float32 `json:"distance" validate:"nonzero"`
	Duration int64   `json:"duration" validate:"nonzero"`
}
