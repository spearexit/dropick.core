package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Workflow struct {
	gorm.Model
	UserId int64          `from:"user_id" json:"user_id"`
	Tags   datatypes.JSON `gorm:"type:json" from:"tags" json:"tags"`
}
