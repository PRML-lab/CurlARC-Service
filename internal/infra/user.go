package infra

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"fmt"
)

type UserRepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	userRepository := UserRepository{SqlHandler: sqlHandler}
	return &userRepository
}

func (userRepo *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	fmt.Println("User @infra:", user)
	result := userRepo.SqlHandler.Conn.Create(user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (userRepo *UserRepository) AuthUser(email, token string) (*model.User, error) {
	user := new(model.User)
	result := userRepo.SqlHandler.Conn.Where("email = ? AND password = ?", email, token).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userRepo *UserRepository) FindById(id string) (*model.User, error) {
	user := new(model.User)
	result := userRepo.SqlHandler.Conn.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userRepo *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := new(model.User)
	result := userRepo.SqlHandler.Conn.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userRepo *UserRepository) Update(user *model.User) error {
	result := userRepo.SqlHandler.Conn.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userRepo *UserRepository) Delete(id string) error {
	result := userRepo.SqlHandler.Conn.Delete(&model.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
