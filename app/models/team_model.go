package models

import "gorm.io/gorm"

// SignIn struct to describe login user.
type Team struct {
	gorm.Model
	// One-to-one relationship with Member table
	ID string `gorm:"primaryKey" json:"id"`

	MemberID string `json:"memberId" gorm:"column:mem_id"`
	Member   Member `json:"member" gorm:"foreignKey:MemberID;references:ID"`
	Name     string `gorm:"column:team_name" json:"name" validate:"required;lte=45"`
}

func (t *Team) TableName() string {
	return "team"
}

type TeamRole struct {
	gorm.Model
	// One-to-one relationship with Member table
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `gorm:"column:team_role_name" json:"name" validate:"required;lte=45"`
}

func (t *TeamRole) TableName() string {
	return "teamrole"
}

// Intermediary table for many-to-many relationship between Team and Member
type TeamHasMember struct {
	gorm.Model

	ID string `gorm:"primaryKey;" json:"id"`

	// One-to-many relationship with Team table
	TeamID string `gorm:"column:team_id" json:"teamId"`
	Team   Team   `json:"team" gorm:"foreignKey:TeamID;references:ID"`

	// One-to-many relationship with TeamRole table
	TeamRoleID string   `gorm:"column:team_role_id" json:"teamRoleId"`
	TeamRole   TeamRole `json:"teamRole" gorm:"foreignKey:TeamRoleID;references:ID"`

	// One-to-many relationship with Member table
	MemberID string `gorm:"column:mem_id" json:"memId"`
	Member   Member `json:"member" gorm:"foreignKey:MemberID;references:ID"`

	MemberTeamName string `gorm:"column:mem_team_name" json:"memTeamName" validate:"lte=45"`
}

func (thm *TeamHasMember) TableName() string {
	return "team_has_member"
}
