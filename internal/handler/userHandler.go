package handler

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/usecase"
	"CurlARC/internal/utils"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lib/pq"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return UserHandler{userUsecase: userUsecase}
}

// 新規ユーザー登録
func (h *UserHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req struct {
			Name     string         `json:"name"`
			Email    string         `json:"email"`
			Password string         `json:"password"`
			TeamIds  pq.StringArray `json:"team_ids"`
		}

		// リクエストのバインド
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		// ユースケースにリクエストを渡す
		err := h.userUsecase.SignUp(c.Request().Context(), req.Name, req.Email, req.Password, req.TeamIds)
		if err != nil {
			if err == repository.ErrEmailExists {
				return c.JSON(http.StatusConflict, map[string]string{"error": "email already exists"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "success"})
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

		// リクエストをユースケースに渡す
		user, err := h.userUsecase.AuthUser(c.Request().Context(), req.IdToken)
		if err != nil {
			if err == repository.ErrUserNotFound {
				return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// JWT 発行
		jwt, err := utils.GenerateJWT(user.Id)
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
