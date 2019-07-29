package user

import (
	"context"
	"database/sql"
	"time"
)

var (
	insertUserQuery = RepositoryQuery{Name: "insertUser", Query: "INSERT testUser SET name=?, email=?"}
	getUserQuery    = RepositoryQuery{Name: "getUser", Query: ""}
)

// RepositoryQuery represents queries of database.
type RepositoryQuery struct {
	Name  string
	Query string
}

// Repository implements repository methods.
type Repository interface {
	insertUser(ctx context.Context, user User) (sql.Result, error)
	getUser(ctx context.Context, id int64) (User, error)
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

func (r *RepositoryImpl) insertUser(ctx context.Context, user User) (sql.Result, error) {
	ctxTimeout, ctxCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer ctxCancel()

	return r.db.ExecInsertItem(ctxTimeout, insertUserQuery, user)
}

func (r *RepositoryImpl) getUser(ctx context.Context, id int64) (User, error) {
	ctxTimeout, ctxCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer ctxCancel()

	rows, err := r.db.QuerySingleResult(ctxTimeout, userDataMapper, getUserQuery, id)
	if err != nil {
		return User{}, err
	}

	user := rows.(User)

	return user, nil
}

func userDataMapper(scanner RowScanner) (interface{}, error) {
	user := User{}

	err := scanner.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Creation,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
