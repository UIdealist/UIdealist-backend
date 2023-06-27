package models

type MemberRole struct {
	ID     int `db:"userRoleId" json:"id"`
	Member int `db:"memId" json:"member" validate:"required"`
	Role   int `db:"roleId" json:"role" validate:"required"`
}
