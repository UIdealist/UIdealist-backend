package models

type Role struct {
	ID   int    `db:"roleId" json:"id" validate:"required"`
	Name string `db:"roleName" json:"name" validate:"required"`
}
