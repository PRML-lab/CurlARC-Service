package usecase

import (
	"CurlARC/internal/domain"
	"CurlARC/internal/repository"
	"context"
)

// UseCase interface
type IUserUseCase interface {
	CreateUser(ctx context.Context, req CreateUserRequest) error
	AuthUser(ctx context.Context, req AuthUserRequest) (AuthUserResponse, error)
	GetUser(ctx context.Context, id string) (UserResponse, error)
}

// DTOs
type CreateUserRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUserResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

// Implementations
type userUseCase struct {
	userRepository repository.UserRepository
	passwordHasher PasswordHasher
	jwtService     JWTService
}

func (u *userUseCase) CreateUser(ctx context.Context, req CreateUserRequest) error {
	hashedPassword, err := u.passwordHasher.Hash(req.Password)
	if err != nil {
		return err
	}
	user := &domain.User{
		Name:     req.UserName,
		Email:    req.Email,
		Password: hashedPassword,
	}
	return u.userRepository.CreateUser(ctx, user)
}
