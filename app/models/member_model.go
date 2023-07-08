package models

import (
	"gorm.io/gorm"
)

// SignIn struct to describe login user.
type Member struct {
	gorm.Model
	ID   string `gorm:"primaryKey" json:"id"`
	Type string `gorm:"column:mem_type" json:"type"`
}

func (member *Member) TableName() string {
	return "member"
}
