package infra

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	userRepository := UserRepository{DB: db}
	return &userRepository
}

func (userRepo *UserRepository) CreateUser(user *model.User) error {
	result := userRepo.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userRepo *UserRepository) AuthUser(email, token string) (*model.User, error) {
	user := new(model.User)
	result := userRepo.DB.Where("email = ? AND password = ?", email, token).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userRepo *UserRepository) FindById(id string) (*model.User, error) {
	user := new(model.User)
	result := userRepo.DB.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userRepo *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := new(model.User)
	result := userRepo.DB.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userRepo *UserRepository) Update(user *model.User) error {
	result := userRepo.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userRepo *UserRepository) Delete(id string) error {
	result := userRepo.DB.Delete(&model.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
