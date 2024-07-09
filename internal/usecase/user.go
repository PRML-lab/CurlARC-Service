package usecase

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"context"
	"errors"
	"fmt"

	"firebase.google.com/go/v4/auth"
	"gorm.io/gorm"
)

// UserUsecase はユーザー関連のユースケースを定義するインターフェースです。
type UserUsecase interface {
	// SignUp は新しいユーザーを登録します。
	SignUp(ctx context.Context, idToken, name, email string, teamIds []string) error
	// AuthUser はユーザーを認証します。
	AuthUser(ctx context.Context, id_token string) (*model.User, error)
	// GetAllUsers は全てのユーザー情報を取得します。
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	// GetUser はログイン中のユーザー情報を取得します。
	GetUser(ctx context.Context, id string) (*model.User, error)
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
	userRepo   repository.UserRepository
	authClient *auth.Client
}

func NewUserUsecase(userRepo repository.UserRepository, authCli *auth.Client) UserUsecase {
	return &userUsecase{userRepo: userRepo, authClient: authCli}
}

func (usecase *userUsecase) SignUp(ctx context.Context, idToken, name, email string, teamIds []string) (err error) {
	// idTokenを検証
	token, err := usecase.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		fmt.Print(err)
		return repository.ErrUnauthorized
	}

	// ユーザーをdbに保存
	user := &model.User{
		Id:      token.UID,
		Name:    name,
		Email:   email,
		TeamIds: teamIds,
	}

	if _, err := usecase.userRepo.Save(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return repository.ErrEmailExists
		}
		usecase.authClient.DeleteUser(ctx, token.UID) // dbへの保存が失敗したらfirebase上のユーザーも削除
		return err
	}

	return nil
}

func (usecase *userUsecase) AuthUser(ctx context.Context, id_token string) (*model.User, error) {
	// Verify the ID token
	authToken, err := usecase.authClient.VerifyIDToken(context.Background(), id_token)
	if err != nil {
		return nil, repository.ErrUnauthorized
	}

	// Find the user by UID
	user, err := usecase.userRepo.FindById(authToken.UID)
	if err != nil {
		return nil, repository.ErrUserNotFound
	}

	return user, nil
}

func (usecase *userUsecase) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return usecase.userRepo.FindAll()
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
