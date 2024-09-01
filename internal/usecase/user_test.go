package usecase_test

import (
	"context"
	"errors"
	"testing"

	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/usecase"
	"CurlARC/mock"

	firebaseAuth "firebase.google.com/go/v4/auth"
	"gorm.io/gorm"

	"github.com/golang/mock/gomock"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockUserRepository(ctrl)
	mockAuth := mock.NewMockAuthClient(ctrl)

	usecase := usecase.NewUserUsecase(mockRepo, mockAuth)

	ctx := context.Background()
	idToken := "valid_id_token"
	name := "John Doe"
	email := "john.doe@example.com"

	t.Run("正常系: ユーザーが正常にサインアップされる", func(t *testing.T) {
		token := &firebaseAuth.Token{
			UID: "firebase_uid",
		}

		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(token, nil)
		mockRepo.EXPECT().Save(gomock.Any()).Return(&model.User{Id: "firebase_uid", Name: name, Email: email}, nil)

		err := usecase.SignUp(ctx, idToken, name, email)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("異常系: IDトークンが無効である", func(t *testing.T) {
		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(nil, errors.New("invalid token"))

		err := usecase.SignUp(ctx, idToken, name, email)
		if !errors.Is(err, repository.ErrUnauthorized) {
			t.Errorf("expected error: %v, got: %v", repository.ErrUnauthorized, err)
		}
	})

	t.Run("異常系: メールアドレスが既に存在する", func(t *testing.T) {
		token := &firebaseAuth.Token{
			UID: "firebase_uid",
		}

		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(token, nil)
		mockRepo.EXPECT().Save(gomock.Any()).Return(nil, gorm.ErrDuplicatedKey)
		mockAuth.EXPECT().DeleteUser(ctx, token.UID).Return(nil)

		err := usecase.SignUp(ctx, idToken, name, email)
		if !errors.Is(err, repository.ErrEmailExists) {
			t.Errorf("expected error: %v, got: %v", repository.ErrEmailExists, err)
		}
	})

	t.Run("異常系: データベースへの保存に失敗する", func(t *testing.T) {
		token := &firebaseAuth.Token{
			UID: "firebase_uid",
		}

		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(token, nil)
		mockRepo.EXPECT().Save(gomock.Any()).Return(nil, errors.New("db error"))
		mockAuth.EXPECT().DeleteUser(ctx, token.UID).Return(nil)

		err := usecase.SignUp(ctx, idToken, name, email)
		if err == nil || errors.Is(err, repository.ErrEmailExists) {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
