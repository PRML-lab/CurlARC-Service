package infra

import (
	"CurlARC/internal/domain/entity"
	"CurlARC/internal/domain/repository"
)

type UserRepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	userRepository := UserRepository{SqlHandler: sqlHandler}
	return &userRepository
}

func (userRepo *UserRepository) Save(user *entity.User) (*entity.User, error) {
	result := userRepo.SqlHandler.Conn.Create(user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (userRepo *UserRepository) FindAll() ([]*entity.User, error) {
	users := []*entity.User{}
	result := userRepo.SqlHandler.Conn.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (userRepo *UserRepository) FindById(id string) (*entity.User, error) {
	user := new(entity.User)
	result := userRepo.SqlHandler.Conn.Where("id = ?", id).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userRepo *UserRepository) FindByEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	result := userRepo.SqlHandler.Conn.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userRepo *UserRepository) Update(user *entity.User) error {
	result := userRepo.SqlHandler.Conn.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userRepo *UserRepository) Delete(id string) error {
	result := userRepo.SqlHandler.Conn.Delete(&entity.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
