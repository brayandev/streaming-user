package user

import (
	"context"
	"database/sql"
)

// Service methods.
type Service interface {
	InsertUser(ctx context.Context, user User) (sql.Result, error)
	Validate(content interface{}) (bool, error)
	GetUser(ctx context.Context, id int64) (User, error)
}

// ServiceImpl service dependecies.
type ServiceImpl struct {
	repository Repository
	validator  Validator
}

// NewService service constructor.
func NewService(repository Repository, validator Validator) *ServiceImpl {
	return &ServiceImpl{
		repository: repository,
		validator:  validator,
	}
}

// InsertUser insert user on db.
func (s ServiceImpl) InsertUser(ctx context.Context, user User) (sql.Result, error) {
	return s.repository.insertUser(ctx, user)
}

// Validate validate input using json schema.
func (s ServiceImpl) Validate(content interface{}) (bool, error) {
	return s.validator.Validate(content)
}

// GetUser get one user by id.
func (s ServiceImpl) GetUser(ctx context.Context, id int64) (User, error) {
	return s.GetUser(ctx, id)
}
