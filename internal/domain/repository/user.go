package repository

import (
	"CurlARC/internal/domain/model"
)

// UserRepository interface
type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	//
	AuthUser(email, token string) (*model.User, error)
	FindAll() ([]*model.User, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Delete(id string) error
}
