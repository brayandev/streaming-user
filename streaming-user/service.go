package user

import (
	"context"
	"database/sql"
)

// Service methods.
type Service interface {
	InsertUser(ctx context.Context, user *User) (sql.Result, error)
}

// ServiceImpl service dependecies.
type ServiceImpl struct {
	repository Repository
}

// NewService service constructor.
func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{
		repository: repository,
	}
}

// InsertUser insert user on db.
func (s ServiceImpl) InsertUser(ctx context.Context, user *User) (sql.Result, error) {
	return s.repository.insertUser(ctx, user)
}
