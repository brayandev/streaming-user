package user

import (
	"context"
	"database/sql"
	"time"
)

var (
	insertUserQuery = RepositoryQuery{Name: "insertUser", Query: "INSERT INTO usr (name, email) values (?, ?)"}
)

// RepositoryQuery represents queries of database.
type RepositoryQuery struct {
	Name  string
	Query string
}

// Repository implements repository methods.
type Repository interface {
	insertUser(ctx context.Context, user *User) (sql.Result, error)
}

// RepositoryImpl repository dependecies.
type RepositoryImpl struct {
	db        Database
	dbTimeout time.Duration
}

// NewRepository repository constructor.
func NewRepository(db Database, dbTimeout time.Duration) *RepositoryImpl {
	return &RepositoryImpl{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r *RepositoryImpl) insertUser(ctx context.Context, user *User) (sql.Result, error) {
	ctxTimeout, ctxCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer ctxCancel()

	return r.db.ExecInsertItem(ctxTimeout, insertUserQuery)
}
