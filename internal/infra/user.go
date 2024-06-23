package infra

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
)

type UserRepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	userRepository := UserRepository{sqlHandler}
	return &userRepository
}

func (userRepo *UserRepository) CreateUser(user *model.User) error {
	_, err := userRepo.Conn.Exec("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", user.Id, user.Name, user.Email, user.TeamIds)
	if err != nil {
		return err
	}
	return nil
}

func (userRepo *UserRepository) AuthUser(email, token string) (*model.User, error) {
	row := userRepo.Conn.QueryRow("SELECT * FROM users WHERE email = ? AND password = ?", email, token)
	user := new(model.User)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.TeamIds)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepo *UserRepository) FindById(id string) (*model.User, error) {
	row := userRepo.Conn.QueryRow("SELECT * FROM users WHERE id = ?", id)
	user := new(model.User)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.TeamIds)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepo *UserRepository) FindByEmail(email string) (*model.User, error) {
	row := userRepo.Conn.QueryRow("SELECT * FROM users WHERE email = ?", email)
	user := new(model.User)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.TeamIds)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepo *UserRepository) Update(user *model.User) error {
	_, err := userRepo.Conn.Exec("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?", user.Name, user.Email, user.TeamIds, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (userRepo *UserRepository) Delete(id string) error {
	_, err := userRepo.Conn.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
