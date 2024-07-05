package handler

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/usecase"
	"CurlARC/internal/utils"
	"context"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
	authClient  *auth.Client
}

func NewUserHandler(userUsecase usecase.UserUsecase, authClient *auth.Client) UserHandler {
	return UserHandler{userUsecase: userUsecase, authClient: authClient}
}

// 新規ユーザー登録
func (h *UserHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
			TeamIds  string `json:"team_ids"`
		}

		// リクエストのバインド
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		// Firebase Auth でユーザーを作成
		params := (&auth.UserToCreate{}).
			Email(req.Email).
			Password(req.Password)

		u, err := h.authClient.CreateUser(context.Background(), params)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// ユーザー情報をDBに保存
		err = h.userUsecase.SignUp(c.Request().Context(), u.UID, req.Name, req.Email, req.TeamIds)
		if err != nil {
			if err == repository.ErrEmailExists {
				return c.JSON(http.StatusConflict, map[string]string{"error": "email already exists"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, err)
	}
}

// ログイン
func (h *UserHandler) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req struct {
			IdToken string `json:"id_token"`
		}

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		// Verify the ID token
		token, err := h.authClient.VerifyIDToken(context.Background(), req.IdToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		// Retrieve user information from PostgreSQL
		_, err = h.userUsecase.GetUser(c.Request().Context(), token.UID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// Generate JWT
		jwt, err := utils.GenerateJWT(token.UID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]string{"jwt": jwt})
	}
}

// ユーザー一覧の取得
func (h *UserHandler) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := h.userUsecase.GetAllUsers(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, users)
	}
}

// ユーザー情報の取得
func (h *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("userID").(string)
		user, err := h.userUsecase.GetUser(c.Request().Context(), userID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}

// ユーザー情報の更新
func (h *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("userID").(string)
		var user model.User
		if err := c.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
		}
		user.Id = userId
		if err := h.userUsecase.UpdateUser(c.Request().Context(), &user); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}

// ユーザーの削除
func (h *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("userID").(string)
		if err := h.userUsecase.DeleteUser(c.Request().Context(), userID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}
