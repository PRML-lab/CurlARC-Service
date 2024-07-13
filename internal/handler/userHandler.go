package handler

import (
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/handler/request"
	"CurlARC/internal/handler/response"
	"CurlARC/internal/usecase"
	"CurlARC/internal/utils"
	"net/http"

	"github.com/labstack/echo"
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
		var req request.SignUpRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		// ユースケースにリクエストを渡す
		err := h.userUsecase.SignUp(c.Request().Context(), req.IdToken, req.Name, req.Email)
		if err != nil {
			if err == repository.ErrUnauthorized {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid id token"})
			} else if err == repository.ErrEmailExists {
				return c.JSON(http.StatusConflict, map[string]string{"error": "email already exists"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.NoContent(http.StatusCreated)
	}
}

// ログイン
func (h *UserHandler) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.SignInRequest
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

		res := response.SignInResponse{
			Jwt:     jwt,
			Id:      user.Id,
			Name:    user.Name,
			Email:   user.Email,
			TeamIds: user.TeamIds,
		}

		return c.JSON(http.StatusOK, res)
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
		id := c.Get("uid").(string)

		user, err := h.userUsecase.GetUser(c.Request().Context(), id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		res := response.GetUserResponse{
			Id:      user.Id,
			Name:    user.Name,
			Email:   user.Email,
			TeamIds: user.TeamIds,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// ユーザー情報の更新
func (h *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.UpdateUserRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
		}
		id := c.Get("uid").(string)

		if err := h.userUsecase.UpdateUser(c.Request().Context(), id, req.Name, req.Email, req.TeamIds); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusNoContent)
	}
}

// ユーザーの削除
func (h *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.DeleteUserRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
		}

		if err := h.userUsecase.DeleteUser(c.Request().Context(), req.Id); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusNoContent)
	}
}
