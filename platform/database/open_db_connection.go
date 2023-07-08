package database

import (
	"idealist/app/models"
	"idealist/app/queries"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

// Queries struct for collect all app queries.
// // These queries which aren't handled by GORM.
type Queries struct {
	// Example:
	// *queries.UserQueries   // load queries from User model
	// ...
	*queries.ProjectQueries
}

var queryList *Queries

// OpenDBConnection func for opening database connection.
func OpenDBConnection() {
	// Define Database connection variables.
	var (
		db  *gorm.DB
		err error
	)

	// Get DB_TYPE value from .env file.
	dbType := os.Getenv("DB_TYPE")

	// Define a new Database connection with right DB type.
	switch dbType {
	case "pgx":
		db, err = PostgresSQLConnection()
	case "mysql":
		db, err = MySQLConnection()
	}

	if err != nil {
		return
	}

	// Set database connection and configure advanced queries
	// for future operations.
	DB = db

	// AutoMigrate all models.
	db.AutoMigrate(
		&models.Member{}, // Member model

		&models.User{},          // User model
		&models.AnonymousUser{}, // AnonymousUser model
		&models.Team{},          // Team model

		&models.Project{}, // Project model

		&models.TeamRole{},      // TeamRole model
		&models.TeamHasMember{}, // TeamHasMember model

		&models.Project{},               // Project model
		&models.ProjectMemberRole{},     // ProjectMemberRole model
		&models.ProjectIncludesMember{}, // ProjectIncludesMember model

		&models.Idea{}, // Idea model
	)

	queryList = &Queries{
		// Set queries from models:
		// Example:
		// UserQueries: &queries.UserQueries{DB: db}, // from User model
		// ...
		ProjectQueries: &queries.ProjectQueries{DB: db},
	}

}
