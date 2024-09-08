package repository

import (
	"CurlARC/internal/domain/entity"
)

// UserRepository interface
type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	FindById(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(id string) error
}
