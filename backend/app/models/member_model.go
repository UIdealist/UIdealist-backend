package models

type Member struct {
	ID   int    `db:"memId" json:"id"`
	Name string `db:"memName" json:"name" validate:"required"`
	Rank int    `db:"rankId" json:"rank" validate:"required"`
}
