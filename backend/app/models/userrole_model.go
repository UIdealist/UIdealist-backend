package models

type UserRole struct {
	ID     int `db:"userRoleId" json:"id" validate:"required"`
	Member int `db:"memId" json:"member" validate:"required"`
	Role   int `db:"roleId" json:"role" validate:"required"`
}
