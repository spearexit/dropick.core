package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname    string         `from:"nickname" json:"nickname"`
	Email       string         `from:"email" json:"email"`
	Password    string         `from:"password" json:"password"`
	GroupId     int64          `from:"group_id" json:"group_id"`
	Permissions datatypes.JSON `gorm:"type:json" from:"permissions" json:"permissions"`
}
