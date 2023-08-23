package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	GroupType          string   `from:"group_type" json:"group_type"`
}
