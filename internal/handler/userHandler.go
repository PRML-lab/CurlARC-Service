package handler

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/usecase"
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

func (h *UserHandler) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		// Firebaseトークンの検証
		decodedToken, err := h.authClient.VerifyIDToken(context.Background(), token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// ユーザーIDをコンテキストに追加
		c.Set("userID", decodedToken.UID)
		return next(c)
	}
}

// 新規ユーザー登録
func (h *UserHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req struct {
			Name    string `json:"name"`
			Email   string `json:"email"`
			TeamIds string `json:"team_ids"`
		}

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		err := h.userUsecase.SignUp(c.Request().Context(), req.Name, req.Email, req.TeamIds)
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
		var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.Bind(&credentials); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
		}
		user, err := h.userUsecase.SignIn(c.Request().Context(), credentials.Email, credentials.Password)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
		}
		return c.JSON(http.StatusOK, user)
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
