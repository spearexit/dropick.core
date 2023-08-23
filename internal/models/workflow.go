package models

import (
	"gorm.io/gorm"
)

type Workflow struct {
	gorm.Model
	UserId int64 `from:"user_id" json:"user_id"`
}
