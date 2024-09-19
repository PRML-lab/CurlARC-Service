package usecase_test

import (
	"context"
	"errors"
	"testing"

	"CurlARC/internal/domain/entity"
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
	user := entity.NewUser(*entity.NewUserId("firebase_uid"), "John Doe", "JohnDoe@gmail.com")
	token := &firebaseAuth.Token{
		UID: "firebase_uid",
	}

	t.Run("正常系: ユーザーが正常に認証される", func(t *testing.T) {

		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(token, nil)
		mockRepo.EXPECT().FindById(token.UID).Return(user, nil)

		authorizedUser, _, err := usecase.SignIn(ctx, idToken)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user.GetId(), authorizedUser.GetId())
	})

	t.Run("異常系: IDトークンが無効である", func(t *testing.T) {
		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(nil, errors.New("invalid token"))

		_, _, err := usecase.SignIn(ctx, idToken)
		assert.Error(t, err)
	})

	t.Run("異常系: ユーザーが見つからない", func(t *testing.T) {
		token := &firebaseAuth.Token{
			UID: "firebase_uid",
		}

		mockAuth.EXPECT().VerifyIDToken(ctx, idToken).Return(token, nil)
		mockRepo.EXPECT().FindById(token.UID).Return(nil, gorm.ErrRecordNotFound)

		_, _, err := usecase.SignIn(ctx, idToken)
		assert.Error(t, err)
	})
}

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockUserRepository(ctrl)
	mockAuth := mock.NewMockAuthClient(ctrl)

	usecase := usecase.NewUserUsecase(mockRepo, mockAuth)

	ctx := context.Background()
	users := []*entity.User{
		entity.NewUser(*entity.NewUserId("1"), "John Doe", "JohnDoe@gmail.com"),
		entity.NewUser(*entity.NewUserId("2"), "Jane Doe", "JaneDoe@gmail.com"),
	}

	t.Run("正常系: ユーザーが正常に取得される", func(t *testing.T) {

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
	user := entity.NewUser(*entity.NewUserId("1"), "John Doe", "JohnDoe@gmail.com")

	t.Run("正常系: ユーザーが正常に取得される", func(t *testing.T) {
		mockRepo.EXPECT().FindById(userId).Return(user, nil)

		result, err := usecase.GetUser(ctx, userId)
		if err != nil {

			t.Errorf("unexpected error: %v", err)
		}
		if result.GetId().Value() != userId {
			t.Errorf("expected user id: %s, got: %s", userId, result.GetId().Value())
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
	newName := "Jane Doe"
	email := "John-Doe@gmail.com"
	newEmail := "Jane-Doe@gmail.com"

	user := entity.NewUser(*entity.NewUserId(userId), name, email)
	toUpdateUser := entity.NewUser(*entity.NewUserId(userId), newName, newEmail)

	updatedAuthUserRecord := &firebaseAuth.UserRecord{
		ProviderUserInfo: []*firebaseAuth.UserInfo{
			{
				DisplayName: name,
				Email:       email,
			},
		},
	}

	t.Run("正常系: ユーザーが正常に更新される", func(t *testing.T) {
		mockRepo.EXPECT().FindById(userId).Return(user, nil)
		mockAuth.EXPECT().UpdateUser(ctx, userId, gomock.Any()).Return(updatedAuthUserRecord, nil)
		mockRepo.EXPECT().Update(gomock.Any()).Return(toUpdateUser, nil)

		updatedUser, err := usecase.UpdateUser(ctx, userId, newName, newEmail)
		assert.NoError(t, err)
		assert.NotNil(t, updatedUser)
		assert.Equal(t, updatedUser.GetName(), newName)
		assert.Equal(t, updatedUser.GetEmail(), newEmail)
	})

	t.Run("異常系: Firebase上のユーザー情報の更新に失敗する", func(t *testing.T) {
		mockRepo.EXPECT().FindById(userId).Return(user, nil)
		mockAuth.EXPECT().UpdateUser(ctx, userId, gomock.Any()).Return(nil, errors.New("firebase error"))

		updatedUser, err := usecase.UpdateUser(ctx, userId, name, email)
		assert.Error(t, err)
		assert.Nil(t, updatedUser)
	})

	t.Run("異常系: データベースのユーザー情報の更新に失敗する", func(t *testing.T) {
		mockRepo.EXPECT().FindById(userId).Return(user, nil)
		mockAuth.EXPECT().UpdateUser(ctx, userId, gomock.Any()).Return(updatedAuthUserRecord, nil)
		mockRepo.EXPECT().Update(gomock.Any()).Return(nil, errors.New("db error"))

		updatedUser, err := usecase.UpdateUser(ctx, userId, name, email)
		assert.Error(t, err)
		assert.Nil(t, updatedUser)
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
