package models

type Rank struct {
	ID    int    `db:"rankId" json:"id" validate:"required"`
	Name  string `db:"rankName" json:"name" validate:"required"`
	Color int    `db:"rankColor" json:"color" validate:"required"`
}
