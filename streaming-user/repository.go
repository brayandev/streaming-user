package user

import (
	"context"
	"time"
)

var (
	insertUserQuery = RepositoryQuery{}
)

// RepositoryQuery represents queries of database.
type RepositoryQuery struct {
	Name  string
	Query string
}

// Repository implements repository methods.
type Repository interface {
	insertUser(ctx context.Context, user *User) error
}

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

func (r *RepositoryImpl) insertUser(ctx context.Context, user *User) error {
	ctxTimeout, ctxCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer ctxCancel()

	_, err := r.db.ExecInsertItem(ctxTimeout, insertUserQuery)
	return err
}
