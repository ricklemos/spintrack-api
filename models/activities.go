package models

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	UserId   int
	Distance float32
	Duration int64
}
