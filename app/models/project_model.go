package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Project model and relations definition
type Project struct {
	// One-to-one relationship with Member table
	ID          string `gorm:"primaryKey;column:pro_id" json:"id"`
	Name        string `gorm:"column:pro_name" json:"name" validate:"required;lte=45"`
	Description string `gorm:"column:pro_description" json:"description" validate:"required;lte=255"`
	Icon        string `gorm:"column:pro_icon" json:"icon" validate:"required;lte=45"`

	OwnerID string `gorm:"column:mem_id" json:"authorId"`
	Owner   Member `json:"author" gorm:"foreignKey:OwnerID;references:ID"`

	// Ideas and brainstorms can be null or related in a rich way
	BrainStorms []*BrainStorm `json:"brainstorms" gorm:"foreignKey:ProjectID;references:ID"`
	Ideas       []*Idea       `json:"ideas" gorm:"foreignKey:ProjectID;references:ID"`
}

func (project *Project) TableName() string {
	return "project"
}

// Before creating the user, create its UUID and member
func (project *Project) BeforeCreate(tx *gorm.DB) (err error) {
	// Create UUID
	project.ID = uuid.New().String()
	return
}

type ProjectMemberRole struct {
	// One-to-one relationship with Member table
	ID   string `gorm:"primaryKey;column:pro_mem_role_id" json:"id"`
	Name string `gorm:"column:pro_mem_role_name" json:"name" validate:"required;lte=45"`
}

func (pmr *ProjectMemberRole) TableName() string {
	return "project_member_role"
}

// Before creating the user, create its UUID and member
func (pmr *ProjectMemberRole) BeforeCreate(tx *gorm.DB) (err error) {
	// Create UUID
	pmr.ID = uuid.New().String()
	return
}

// Intermediary table for many-to-many relationship between Team and Member
type ProjectIncludesMember struct {
	ID string `gorm:"primaryKey;column:pro_inc_mem_id" json:"id"`

	// One-to-many relationship with Team table
	ProjectID string  `gorm:"column:pro_id" json:"proId"`
	Project   Project `json:"project" gorm:"foreignKey:ProjectID;references:ID"`

	// One-to-many relationship with TeamRole table
	ProjectMemberRoleID string            `gorm:"column:pro_mem_role_id" json:"proMemRoleId"`
	ProjectMemberRole   ProjectMemberRole `json:"projectMemberRole" gorm:"foreignKey:ProjectMemberRoleID;references:ID"`

	// One-to-many relationship with Member table
	MemberID string `gorm:"column:mem_id" json:"memId"`
	Member   Member `json:"member" gorm:"foreignKey:MemberID;references:ID"`
}

func (pim *ProjectIncludesMember) TableName() string {
	return "project_includes_member"
}

// Before creating the user, create its UUID and member
func (pim *ProjectIncludesMember) BeforeCreate(tx *gorm.DB) (err error) {
	// Create UUID
	pim.ID = uuid.New().String()
	return
}
