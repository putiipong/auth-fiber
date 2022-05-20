package db

import (
	"crud/domain"
)

// Queries struct for collect all app queries.
type Queries struct {
	*domain.UserQueries // load queries from User model
	// *domain.BookQueries // load queries from Book model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries: &domain.UserQueries{DB: db}, // from User model
		// BookQueries: &domain.BookQueries{DB: db}, // from Book model
	}, nil
}
