package user

import (
	"context"
	"database/sql"
)

// Service methods.
type Service interface {
	InsertUser(ctx context.Context, user User) (sql.Result, error)
	JSONValidator(content interface{}) (bool, error)
}

// ServiceImpl service dependecies.
type ServiceImpl struct {
	repository Repository
	validators Validators
}

// NewService service constructor.
func NewService(repository Repository, validators Validators) *ServiceImpl {
	return &ServiceImpl{
		repository: repository,
		validators: validators,
	}
}

// InsertUser insert user on db.
func (s ServiceImpl) InsertUser(ctx context.Context, user User) (sql.Result, error) {
	return s.repository.insertUser(ctx, user)
}

// JSONValidator validate json input.
func (s ServiceImpl) JSONValidator(content interface{}) (bool, error) {
	return s.validators.JSONValidator(content)
}
