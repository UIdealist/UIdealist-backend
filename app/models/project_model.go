package models

import "gorm.io/gorm"

// SignIn struct to describe login user.
type Project struct {
	gorm.Model
	// One-to-one relationship with Member table
	ID          string `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"column:pro_name" json:"name" validate:"required;lte=45"`
	Description string `gorm:"column:pro_description" json:"description" validate:"required;lte=255"`
	Icon        string `gorm:"column:pro_icon" json:"icon" validate:"required;lte=45"`
}

func (project *Project) TableName() string {
	return "project"
}

type ProjectMemberRole struct {
	gorm.Model
	// One-to-one relationship with Member table
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `gorm:"column:pro_mem_role_name" json:"name" validate:"required;lte=45"`
}

func (pmr *ProjectMemberRole) TableName() string {
	return "project_member_role"
}

// Intermediary table for many-to-many relationship between Team and Member
type ProjectIncludesMember struct {
	gorm.Model

	ID string `gorm:"primaryKey" json:"id"`

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
