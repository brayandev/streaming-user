package user

import "time"

// RepositoryQuery represents queries of database.
type RepositoryQuery struct {
	Name  string
	Query string
}

// Repository implements repository methods.
type Repository interface{}

// RepositoryImpl repository dependecies.
type RepositoryImpl struct {
	db        Database
	dbTimeout time.Duration
}

// NewRepositoryImpl repository constructor.
func NewRepositoryImpl(db Database, dbTimeout time.Duration) *RepositoryImpl {
	return &RepositoryImpl{
		db:        db,
		dbTimeout: dbTimeout,
	}
}
