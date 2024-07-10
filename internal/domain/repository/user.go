package repository

import (
	"CurlARC/internal/domain/model"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrEmailExists  = errors.New("email already exists")
	ErrUnauthorized = errors.New("unauthorized")
)

// UserRepository interface
type UserRepository interface {
	Save(user *model.User) (*model.User, error)
	FindAll() ([]*model.User, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Delete(id string) error
}
