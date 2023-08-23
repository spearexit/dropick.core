package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	TagName          string   `from:"tag_type" json:"tag_type"`
}
