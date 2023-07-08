package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"idealist/pkg/repository"
)

// SignIn struct to describe login user.
type User struct {
	gorm.Model
	// One-to-one relationship with Member table
	ID string `json:"id" gorm:"primaryKey"`

	MemberID string `json:"memberId" gorm:"column:mem_id"`
	Member   Member `json:"member" gorm:"foreignKey:MemberID;references:ID"`

	Username string `gorm:"column:usr_username" json:"username" validate:"required"`
	Password string `gorm:"column:usr_password" json:"password" validate:"required"`
	Email    string `gorm:"column:usr_email" json:"email" validate:"required"`
	Verified bool   `gorm:"column:usr_verified" json:"verified"`
}

func (user *User) TableName() string {
	return "user"
}

type AnonymousUser struct {
	gorm.Model
	// One-to-one relationship with Member table
	ID string `gorm:"primaryKey" json:"id"`

	MemberID string `json:"memberId" gorm:"column:mem_id"`
	Member   Member `json:"member" gorm:"foreignKey:MemberID;references:ID"`

	TempName string `gorm:"column:auser_temp_name" json:"tempName" validate:"required"`
}

func (au *AnonymousUser) TableName() string {
	return "anonymoususer"
}

// Before creating the user, create its UUID and member
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Create UUID
	user.ID = uuid.New().String()

	// Create member
	member := Member{
		ID:   uuid.New().String(),
		Type: repository.MemberIsUser,
	}

	// Reference member to user
	user.MemberID = member.ID

	tx.Create(&member)

	return
}

// Before creating the anonymous user, create its UUID and member
func (anonymousUser *AnonymousUser) BeforeCreate(tx *gorm.DB) (err error) {
	// Create UUID
	anonymousUser.ID = uuid.New().String()

	// Create member
	member := Member{
		ID:   uuid.New().String(),
		Type: repository.MemberIsAnonymousUser,
	}

	// Reference member to anonymous user
	anonymousUser.MemberID = member.ID

	return
}
