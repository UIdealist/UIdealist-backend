package models

type Role struct {
	ID   int    `db:"roleId" json:"id"`
	Name string `db:"roleName" json:"name" validate:"required"`
}
