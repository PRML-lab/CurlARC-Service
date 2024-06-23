package usecase

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"context"
)

// UserUsecase はユーザー関連のユースケースを定義するインターフェースです。
type UserUsecase interface {
	// SignUp は新しいユーザーを登録します。
	SignUp(ctx context.Context, user *model.User) error
	// SignIn はユーザーのログインを処理します。
	SignIn(ctx context.Context, email, password string) (*model.User, error)
	// GetUser はログイン中のユーザー情報を取得します。
	GetUser(ctx context.Context, userID string) (*model.User, error)
	// UpdateUser はユーザー情報を更新します。
	UpdateUser(ctx context.Context, user *model.User) error
	// DeleteUser はユーザーを削除します。
	DeleteUser(ctx context.Context, userID string) error
	// // AcceptTeamInvitation はチームへの招待を承認します。
	// AcceptTeamInvitation(ctx context.Context, userID, teamID string) error
	// // RejectTeamInvitation はチームへの招待を拒否します。
	// RejectTeamInvitation(ctx context.Context, userID, teamID string) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	userUsecase := userUsecase{userRepo: userRepo}
	return &userUsecase
}

func (usecase *userUsecase) SignUp(ctx context.Context, user *model.User) error {
	return usecase.userRepo.CreateUser(user)
}

func (usecase *userUsecase) SignIn(ctx context.Context, email, token string) (*model.User, error) {
	return usecase.userRepo.AuthUser(email, token)
}

func (usecase *userUsecase) GetUser(ctx context.Context, userID string) (*model.User, error) {
	return usecase.userRepo.FindById(userID)
}

func (usecase *userUsecase) UpdateUser(ctx context.Context, user *model.User) error {
	return usecase.userRepo.Update(user)
}

func (usecase *userUsecase) DeleteUser(ctx context.Context, userID string) error {
	return usecase.userRepo.Delete(userID)
}
