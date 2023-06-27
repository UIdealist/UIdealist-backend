package models

type MemberRole struct {
	ID     int `db:"memroleid" json:"id"`
	Member int `db:"memid" json:"member" validate:"required"`
	Role   int `db:"roleid" json:"role" validate:"required"`
}
