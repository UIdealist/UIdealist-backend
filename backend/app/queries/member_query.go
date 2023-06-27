package queries

import (
	"management-helper/app/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// BookQueries struct for queries from Book model.
type MemberQueries struct {
	*sqlx.DB
}

func (q *MemberQueries) GetMembers() ([]models.Member, error) {
	members := []models.Member{}

	query := "SELECT * FROM member"

	err := q.Select(&members, query)

	if err != nil {
		return members, err
	}

	return members, nil
}

func (q *MemberQueries) GetMember(id uuid.UUID) (models.Member, error) {
	member := models.Member{}

	query := "SELECT * FROM member WHERE memId = $1"

	err := q.Get(&member, query, id)

	if err != nil {
		return member, err
	}

	return member, nil
}

func (q *MemberQueries) CreateMember(member *models.Member) error {
	query := "INSERT INTO member(memName, rankId) VALUES ($1, $2)"

	_, err := q.Exec(query, member.Name, member.Rank)
	if err != nil {
		return err
	}

	return nil
}

func (q *MemberQueries) UpdateMember(member *models.Member) error {
	query := "UPDATE member SET memName = $2, rankId = $3 WHERE memId = $1"

	_, err := q.Exec(query, member.ID, member.Name, member.Rank)
	if err != nil {
		return err
	}

	return nil
}

func (q *MemberQueries) DeleteMember(id uuid.UUID) error {
	query := "DELETE FROM member WHERE memId = $1"

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
