package database

import (
	"management-helper/app/queries"
	"os"

	"github.com/jmoiron/sqlx"
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries   // load queries from User model
	*queries.MemberQueries // load queries from Member model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define Database connection variables.
	var (
		db  *sqlx.DB
		err error
	)

	// Get DB_TYPE value from .env file.
	dbType := os.Getenv("DB_TYPE")

	// Define a new Database connection with right DB type.
	switch dbType {
	case "pgx":
		db, err = PostgreSQLConnection()
	case "mysql":
		db, err = MysqlConnection()
	}

	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries:   &queries.UserQueries{DB: db},   // from User model
		MemberQueries: &queries.MemberQueries{DB: db}, // from Book model
	}, nil
}
