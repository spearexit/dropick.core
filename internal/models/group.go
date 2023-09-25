package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name string `from:"name" json:"name"`
	Type string `from:"type" json:"type"`
}
