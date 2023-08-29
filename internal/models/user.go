package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname string `from:"nickname" json:"nickname"`
}
