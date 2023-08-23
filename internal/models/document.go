package models

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	UserId int64 `from:"user_id" json:"user_id"`
}
