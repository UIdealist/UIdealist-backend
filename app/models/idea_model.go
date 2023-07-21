package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Idea struct and related models definition
type Idea struct {
	// One-to-one relationship with Member table
	ID    string `gorm:"primaryKey" json:"id"`
	Title string `gorm:"column:idea_title" json:"title" validate:"required;lte=45"`
	Body  string `gorm:"column:idea_body" json:"body" validate:"required;lte=255"`

	// One-to-many relationship with Member table
	AuthorID string `gorm:"column:mem_id" json:"memId"`
	Author   Member `json:"member" gorm:"foreignKey:AuthorID;references:ID"`

	// One-to-many relationship with Project table (Not required)
	ProjectID *string  `gorm:"column:pro_id" json:"proId"`
	Project   *Project `json:"project" gorm:"foreignKey:ProjectID;references:ID"`

	// One-to-many relationship with Brainstorm table (Not required)
	BrainStormID *string     `gorm:"column:brs_id" json:"brainstormId"`
	BrainStorm   *BrainStorm `json:"brainstorm" gorm:"foreignKey:BrainStormID;references:ID"`
}

// Before creating the idea, create its UUID
func (idea *Idea) BeforeCreate(tx *gorm.DB) (err error) {
	// Create UUID
	idea.ID = uuid.New().String()
	return
}

func (idea *Idea) TableName() string {
	return "idea"
}

// BrainStorm model definition
type BrainStorm struct {
	ID   string `gorm:"primaryKey;column:brs_id" json:"id"`
	Name string `gorm:"column:brs_name" json:"name" validate:"required;lte=45"`

	// One-to-many relationship with Member (author)
	AuthorID string `gorm:"column:mem_id" json:"memId"`
	Author   Member `json:"author" gorm:"foreignKey:AuthorID;references:ID"`

	// One-to-many relationship with Project table (Not required)
	ProjectID *string  `gorm:"column:pro_id" json:"proId"`
	Project   *Project `json:"project" gorm:"foreignKey:ProjectID;references:ID"`

	Ideas []*Idea `json:"idea" gorm:"foreignKey:BrainStormID;references:ID"`
}

// Before creating the idea, create its UUID
func (bs *BrainStorm) BeforeCreate(tx *gorm.DB) (err error) {
	// Create UUID
	bs.ID = uuid.New().String()
	return
}

func (bs *BrainStorm) TableName() string {
	return "brainstorm"
}
