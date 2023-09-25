package models

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	UserId     int64  `from:"user_id" json:"user_id"`
	WorkflowId int64  `from:"workflow_id" json:"workflow_id"`
	Title      string `from:"title" json:"title"`
	Content    string `from:"content" json:"content"`
}
