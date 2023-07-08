package queries

import (
	"idealist/app/models"

	"gorm.io/gorm"
)

type ProjectQueries struct {
	*gorm.DB
}

func (q *ProjectQueries) GetAllProjectMemberUsers() []models.User {
	members := []models.User{}

	query := `SELECT user.* FROM project 
					NATURAL JOIN project_includes_member 
					NATURAL JOIN member 
					NATURAL JOIN user
					WHERE project.id = ?

				UNION SELECT * FROM project
					NATURAL JOIN project_includes_member
					NATURAL JOIN member
					NATURAL JOIN team
					NATURAL JOIN team_has_user
					NATURAL JOIN user
					WHERE project.id = ?
			`
	// We will ignore anonymous users for now

	q.Select(&members, query)

	return members
}
