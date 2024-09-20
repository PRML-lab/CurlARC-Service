package infra

import (
	"CurlARC/internal/domain/entity"
	"CurlARC/internal/domain/repository"
)

type UserRepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	rsitory := UserRepository{SqlHandler: sqlHandler}
	return &rsitory
}

func (u *User) FromDomain(user *entity.User) {
	u.Id = user.GetId().Value()
	u.Name = user.GetName()
	u.Email = user.GetEmail()
}

func (u *User) ToDomain() *entity.User {
	user := entity.NewUser(*entity.NewUserId(u.Id), u.Name, u.Email)
	return user
}

////////////////////////////////////////
// User Repository Inplementation
////////////////////////////////////////

func (r *UserRepository) Save(user *entity.User) (*entity.User, error) {
	var User User
	User.FromDomain(user)

	if err := r.Conn.Create(&User).Error; err != nil {
		return nil, err
	}

	return User.ToDomain(), nil
}

func (r *UserRepository) FindAll() ([]*entity.User, error) {
	var users []User
	if err := r.Conn.Find(&users).Error; err != nil {
		return nil, err
	}

	var usersEntity []*entity.User
	for _, user := range users {
		usersEntity = append(usersEntity, user.ToDomain())
	}

	return usersEntity, nil
}

func (r *UserRepository) FindById(id string) (*entity.User, error) {
	var user User
	if err := r.Conn.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user User
	if err := r.Conn.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

func (r *UserRepository) Update(user *entity.User) (*entity.User, error) {
	var User User
	User.FromDomain(user)

	if err := r.Conn.Save(&User).Error; err != nil {
		return nil, err
	}

	return User.ToDomain(), nil
}

func (r *UserRepository) Delete(id string) error {
	if err := r.Conn.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}

	return nil
}
