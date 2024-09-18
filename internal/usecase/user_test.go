package usecase_test

import (
	"context"
	"errors"
	"testing"

	"CurlARC/internal/domain/entity"
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/usecase"
	"CurlARC/mock"

	firebaseAuth "firebase.google.com/go/v4/auth"
	"gorm.io/gorm"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
	user := entity.NewUser(*entity.NewUserId("firebase_uid"), name, email)

	t.Run("正常系: ユーザーが正常にサインアップされる", func(t *testing.T) {
		token := &firebaseAuth.Token{
			UID: "firebase_uid",
		}

		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(token, nil)
		mockRepo.EXPECT().Save(gomock.Any()).Return(user, nil)

		_, err := usecase.SignUp(ctx, idToken, name, email)
		assert.NoError(t, err)
	})

	t.Run("異常系: IDトークンが無効である", func(t *testing.T) {
		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(nil, errors.New("invalid token"))

		signUpedUser, err := usecase.SignUp(ctx, idToken, name, email)
		assert.Error(t, err)
		assert.Nil(t, signUpedUser)
		assert.Equal(t, errors.New("invalid token"), err)
	})

	t.Run("異常系: メールアドレスが既に存在する", func(t *testing.T) {
		token := &firebaseAuth.Token{
			UID: "firebase_uid",
		}

		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(token, nil)
		mockRepo.EXPECT().Save(gomock.Any()).Return(nil, gorm.ErrDuplicatedKey)
		mockAuth.EXPECT().DeleteUser(ctx, token.UID).Return(nil)

		signUpedUser, err := usecase.SignUp(ctx, idToken, name, email)
		assert.Error(t, err)
		assert.Nil(t, signUpedUser)
	})

	t.Run("異常系: データベースへの保存に失敗する", func(t *testing.T) {
		token := &firebaseAuth.Token{
			UID: "firebase_uid",
		}

		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(token, nil)
		mockRepo.EXPECT().Save(gomock.Any()).Return(nil, errors.New("db error"))
		mockAuth.EXPECT().DeleteUser(ctx, token.UID).Return(nil)

		signUpedUser, err := usecase.SignUp(ctx, idToken, name, email)
		assert.Error(t, err)
		assert.Nil(t, signUpedUser)
	})
}

func TestAuthUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockUserRepository(ctrl)
	mockAuth := mock.NewMockAuthClient(ctrl)

	usecase := usecase.NewUserUsecase(mockRepo, mockAuth)

	ctx := context.Background()
	idToken := "valid_id_token"

	t.Run("正常系: ユーザーが正常に認証される", func(t *testing.T) {
		token := &firebaseAuth.Token{
			UID: "firebase_uid",
		}

		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(token, nil)
		mockRepo.EXPECT().FindById(token.UID).Return(&entity.User{Id: "firebase_uid"}, nil)

		user, _, err := usecase.AuthUser(ctx, idToken)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if user.Id != "firebase_uid" {
			t.Errorf("expected user id: %s, got: %s", "firebase_uid", user.Id)
		}
	})

	t.Run("異常系: IDトークンが無効である", func(t *testing.T) {
		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(nil, errors.New("invalid token"))

		_, _, err := usecase.AuthUser(ctx, idToken)
		if !errors.Is(err, repository.ErrUnauthorized) {
			t.Errorf("expected error: %v, got: %v", repository.ErrUnauthorized, err)
		}
	})

	t.Run("異常系: ユーザーが見つからない", func(t *testing.T) {
		token := &firebaseAuth.Token{
			UID: "firebase_uid",
		}

		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(token, nil)
		mockRepo.EXPECT().FindById(token.UID).Return(nil, gorm.ErrRecordNotFound)

		_, _, err := usecase.AuthUser(ctx, idToken)
		if !errors.Is(err, repository.ErrUserNotFound) {
			t.Errorf("expected error: %v, got: %v", repository.ErrUserNotFound, err)
		}
	})
}

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockUserRepository(ctrl)
	mockAuth := mock.NewMockAuthClient(ctrl)

	usecase := usecase.NewUserUsecase(mockRepo, mockAuth)

	ctx := context.Background()

	t.Run("正常系: ユーザーが正常に取得される", func(t *testing.T) {
		users := []*entity.User{
			{Id: "1", Name: "John Doe", Email: "John-Doe@gmail.com"},
			{Id: "2", Name: "Jane Doe", Email: "Jane-Doe@gmail.com"},
		}

		mockRepo.EXPECT().FindAll().Return(users, nil)

		result, err := usecase.GetAllUsers(ctx)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected length: %d, got: %d", 2, len(result))
		}
	})
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockUserRepository(ctrl)
	mockAuth := mock.NewMockAuthClient(ctrl)

	usecase := usecase.NewUserUsecase(mockRepo, mockAuth)

	ctx := context.Background()
	userId := "1"

	t.Run("正常系: ユーザーが正常に取得される", func(t *testing.T) {
		user := &entity.User{Id: userId, Name: "John Doe", Email: "example@co.jp"}
		mockRepo.EXPECT().FindById(userId).Return(user, nil)

		result, err := usecase.GetUser(ctx, userId)
		if err != nil {

			t.Errorf("unexpected error: %v", err)
		}
		if result.Id != userId {
			t.Errorf("expected user id: %s, got: %s", userId, result.Id)
		}
	})
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockUserRepository(ctrl)
	mockAuth := mock.NewMockAuthClient(ctrl)

	usecase := usecase.NewUserUsecase(mockRepo, mockAuth)

	ctx := context.Background()
	userId := "1"
	name := "John Doe"
	email := "John-Doe@gmail.com"

	updatedAuthUserRecord := &firebaseAuth.UserRecord{
		ProviderUserInfo: []*firebaseAuth.UserInfo{
			{
				DisplayName: name,
				Email:       email,
			},
		},
	}

	t.Run("正常系: ユーザーが正常に更新される", func(t *testing.T) {
		mockAuth.EXPECT().UpdateUser(ctx, userId, gomock.Any()).Return(updatedAuthUserRecord, nil)
		mockRepo.EXPECT().Update(gomock.Any()).Return(nil)

		err := usecase.UpdateUser(ctx, userId, name, email)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("異常系: Firebase上のユーザー情報の更新に失敗する", func(t *testing.T) {
		mockAuth.EXPECT().UpdateUser(ctx, userId, gomock.Any()).Return(nil, errors.New("firebase error"))

		err := usecase.UpdateUser(ctx, userId, name, email)
		if err == nil {
			t.Errorf("expected error, got: %v", err)
		}
	})

	t.Run("異常系: データベースのユーザー情報の更新に失敗する", func(t *testing.T) {
		mockAuth.EXPECT().UpdateUser(ctx, userId, gomock.Any()).Return(updatedAuthUserRecord, nil)
		mockRepo.EXPECT().Update(gomock.Any()).Return(errors.New("db error"))

		err := usecase.UpdateUser(ctx, userId, name, email)
		if err == nil {
			t.Errorf("expected error, got: %v", err)
		}
	})
}

func TestDeleteUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockUserRepository(ctrl)
	mockAuth := mock.NewMockAuthClient(ctrl)

	usecase := usecase.NewUserUsecase(mockRepo, mockAuth)

	ctx := context.Background()
	userId := "1"

	t.Run("正常系: ユーザーが正常に削除される", func(t *testing.T) {
		mockAuth.EXPECT().DeleteUser(ctx, userId).Return(nil)
		mockRepo.EXPECT().Delete(userId).Return(nil)

		err := usecase.DeleteUser(ctx, userId)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("異常系: Firebase上のユーザー情報の削除に失敗する", func(t *testing.T) {
		mockAuth.EXPECT().DeleteUser(ctx, userId).Return(errors.New("firebase error"))

		err := usecase.DeleteUser(ctx, userId)
		if err == nil {
			t.Errorf("expected error, got: %v", err)
		}
	})

	t.Run("異常系: データベースのユーザー情報の削除に失敗する", func(t *testing.T) {
		mockAuth.EXPECT().DeleteUser(ctx, userId).Return(nil)
		mockRepo.EXPECT().Delete(userId).Return(errors.New("db error"))

		err := usecase.DeleteUser(ctx, userId)
		if err == nil {
			t.Errorf("expected error, got: %v", err)
		}
	})
}
