package models

type Member struct {
	ID   int    `db:"memid" json:"id"`
	Name string `db:"memname" json:"name" validate:"required"`
	Rank int    `db:"rankid" json:"rank" validate:"required"`
}
