package models

type Rank struct {
	ID    int    `db:"rankId" json:"id"`
	Name  string `db:"rankName" json:"name" validate:"required"`
	Color int    `db:"rankColor" json:"color" validate:"required"`
}

type RankAndMembers struct {
	Rank   Rank   `json:"rank"`
	Member Member `json:"member"`
}
