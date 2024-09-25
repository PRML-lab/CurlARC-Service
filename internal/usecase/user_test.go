package usecase_test

import (
	"errors"
	"testing"

	"CurlARC/internal/domain/entity"
	"CurlARC/internal/usecase"
	"CurlARC/mock"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// func TestAuthorize(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockUserRepository(ctrl)

// 	usecase := usecase.NewUserUsecase(mockRepo)

// 	ctx := echo.New().NewContext(nil, nil)
// 	idToken := "idToken"
// 	userName := "John Doe"
// 	userEmail := "JohnDoe@gmail.com"
// 	user := entity.NewUser(userName, userEmail)

// 	t.Run("正常系: ユーザーが正常に認証される", func(t *testing.T) {
// 		mockRepo.EXPECT().FindByEmail(userEmail).Return(user, nil)

// 		authorizedUser, _, err := usecase.Authorize(ctx, idToken)
// 		assert.NoError(t, err)
// 		assert.NotNil(t, user)
// 		assert.Equal(t, user.GetId(), authorizedUser.GetId())
// 	})
// }

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockUserRepository(ctrl)

	usecase := usecase.NewUserUsecase(mockRepo)

	ctx := echo.New().NewContext(nil, nil)
	users := []*entity.User{
		entity.NewUser("John Doe", "JohnDoe@gmail.com"),
		entity.NewUser("Jane Doe", "JaneDoe@gmail.com"),
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

	usecase := usecase.NewUserUsecase(mockRepo)

	ctx := echo.New().NewContext(nil, nil)
	user := entity.NewUser("John Doe", "JohnDoe@gmail.com")
	userId := user.GetId().Value()

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

	usecase := usecase.NewUserUsecase(mockRepo)

	ctx := echo.New().NewContext(nil, nil)
	userId := "1"
	name := "John Doe"
	newName := "Jane Doe"
	email := "John-Doe@gmail.com"
	newEmail := "Jane-Doe@gmail.com"

	user := entity.NewUser(name, email)
	toUpdateUser := entity.NewUserFromDB(user.GetId().Value(), newName, newEmail)

	t.Run("正常系: ユーザーが正常に更新される", func(t *testing.T) {
		mockRepo.EXPECT().FindById(userId).Return(user, nil)
		mockRepo.EXPECT().Update(gomock.Any()).Return(toUpdateUser, nil)

		updatedUser, err := usecase.UpdateUser(ctx, userId, newName, newEmail)
		assert.NoError(t, err)
		assert.NotNil(t, updatedUser)
		assert.Equal(t, updatedUser.GetName(), newName)
		assert.Equal(t, updatedUser.GetEmail(), newEmail)
	})

	t.Run("異常系: データベースのユーザー情報の更新に失敗する", func(t *testing.T) {
		mockRepo.EXPECT().FindById(userId).Return(user, nil)
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

	usecase := usecase.NewUserUsecase(mockRepo)

	ctx := echo.New().NewContext(nil, nil)
	userId := "1"

	t.Run("正常系: ユーザーが正常に削除される", func(t *testing.T) {
		mockRepo.EXPECT().Delete(userId).Return(nil)

		err := usecase.DeleteUser(ctx, userId)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("異常系: データベースのユーザー情報の削除に失敗する", func(t *testing.T) {
		mockRepo.EXPECT().Delete(userId).Return(errors.New("db error"))

		err := usecase.DeleteUser(ctx, userId)
		if err == nil {
			t.Errorf("expected error, got: %v", err)
		}
	})
}
