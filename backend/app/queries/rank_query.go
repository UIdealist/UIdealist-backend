package queries

import (
	"management-helper/app/models"

	"github.com/jmoiron/sqlx"
)

type RankQueries struct {
	*sqlx.DB
}

func (q *RankQueries) GetRanks() ([]models.Rank, error) {
	ranks := []models.Rank{}

	query := `SELECT * FROM rank`

	err := q.Select(&ranks, query)
	if err != nil {
		return nil, err
	}

	return ranks, nil
}

func (q *RankQueries) CreateRank(rank models.Rank) error {
	query := "INSERT INTO rank(rankName, rankColor) VALUES ($1, $2)"

	_, err := q.Exec(query, rank.Name, rank.Color)
	if err != nil {
		return err
	}

	return nil
}

// Special query for getting ranks alongside with members associated with them.

func (q *RankQueries) GetRanksAndMembers() ([]models.RankAndMembers, error) {
	ranks := []models.RankAndMembers{}

	query := `SELECT * FROM rank r
		INNER JOIN member m ON r.rankId = m.rankId`

	err := q.Select(&ranks, query)
	if err != nil {
		return nil, err
	}

	return ranks, nil
}
