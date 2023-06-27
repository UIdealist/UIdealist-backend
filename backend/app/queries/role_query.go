package queries

import (
	"management-helper/app/models"

	"github.com/jmoiron/sqlx"
)

type RoleQueries struct {
	*sqlx.DB
}

func (q *RoleQueries) GetRoles() ([]models.Role, error) {
	roles := []models.Role{}

	query := `SELECT roleName FROM role`

	err := q.Select(&roles, query)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (q *RoleQueries) CreateRole(role models.Role) error {
	query := "INSERT INTO role(roleName) VALUES ($1)"

	_, err := q.Exec(query, role.Name)
	if err != nil {
		return err
	}

	return nil
}

func (q *RoleQueries) AssignRole(role int, member int) error {
	query := "INSERT INTO memberrole(memId, roleId) VALUES ($1, $2)"

	_, err := q.Exec(query, member, role)
	if err != nil {
		return err
	}

	return nil
}

func (q *RoleQueries) RemoveRole(role int, member int) error {
	query := "DELETE FROM memberrole WHERE memId = $1 AND roleId = $2"

	_, err := q.Exec(query, member, role)
	if err != nil {
		return err
	}

	return nil
}
