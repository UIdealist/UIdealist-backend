package models

import "gorm.io/gorm"

// SignIn struct to describe login user.
type Idea struct {
	gorm.Model
	// One-to-one relationship with Member table
	ID    string `gorm:"primaryKey" json:"id"`
	Title string `gorm:"column:idea_title" json:"title" validate:"required;lte=45"`
	Body  string `gorm:"column:idea_body" json:"body" validate:"required;lte=255"`

	// One-to-many relationship with Member table
	MemberID string `gorm:"column:mem_id" json:"memId"`
	Member   Member `json:"member" gorm:"foreignKey:MemberID;references:ID"`

	// One-to-many relationship with Project table (Not required)
	ProjectID string  `gorm:"column:pro_id" json:"proId"`
	Project   Project `json:"project" gorm:"foreignKey:ProjectID;references:ID"`
}

func (idea *Idea) TableName() string {
	return "idea"
}
