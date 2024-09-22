package usecase

import (
	"CurlARC/internal/domain/entity"
	"CurlARC/internal/domain/repository"

	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	// CRUD
	Authorize(c echo.Context, name, email string) (*entity.User, error)
	GetAllUsers(c echo.Context) ([]*entity.User, error)
	GetUser(c echo.Context, id string) (*entity.User, error)
	UpdateUser(c echo.Context, id, name, email string) (*entity.User, error)
	DeleteUser(c echo.Context, id string) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (usecase *userUsecase) Authorize(c echo.Context, name, email string) (*entity.User, error) {
	// Find the user by email
	user, err := usecase.userRepo.FindByEmail(email)
	if err != nil && err.Error() == "record not found" {
		// If the user does not exist, create and save a new user
		user = entity.NewUser(name, email)
		user, err = usecase.userRepo.Save(user)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (usecase *userUsecase) GetAllUsers(c echo.Context) ([]*entity.User, error) {
	return usecase.userRepo.FindAll()
}

func (usecase *userUsecase) GetUser(c echo.Context, id string) (*entity.User, error) {
	return usecase.userRepo.FindById(id)
}

func (usecase *userUsecase) UpdateUser(c echo.Context, id, name, email string) (*entity.User, error) {
	// Check if the user exists
	_, err := usecase.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	// update user info in database
	user := entity.NewUserFromDB(id, name, email)
	updatedUser, err := usecase.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, err
}

func (usecase *userUsecase) DeleteUser(c echo.Context, id string) error {

	return usecase.userRepo.Delete(id)
}
