package queries

import (
	"management-helper/app/models"

	"github.com/jmoiron/sqlx"
)

// UserQueries struct for queries from User model.
type UserQueries struct {
	*sqlx.DB
}

// GetUserByEmail query for getting one User by given Email.
func (q *UserQueries) GetUserByUsername(usename string) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Define query string.
	query := `SELECT * FROM appuser WHERE usrusername = $1`

	// Send query to database.
	err := q.Get(&user, query, usename)
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// CreateUser query for creating a new user by given email and password hash.
func (q *UserQueries) CreateUser(u *models.User) error {
	// Define query string.
	query := `INSERT INTO appuser VALUES ($1, $2)`

	// Send query to database.
	_, err := q.Exec(
		query,
		u.Username, u.Password,
	)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
