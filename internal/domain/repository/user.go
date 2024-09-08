package repository

import (
	entity "CurlARC/internal/domain/entity/user"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrEmailExists  = errors.New("email already exists")
	ErrUnauthorized = errors.New("unauthorized")
)

// UserRepository interface
type UserRepository interface {
	Save(user *entity.User) error
	FindAll() ([]*entity.User, error)
	FindById(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id string) error
}
