package models

type Role struct {
	ID   int    `db:"roleid" json:"id"`
	Name string `db:"rolename" json:"name" validate:"required"`
}
