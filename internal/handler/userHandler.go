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
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			})
		}

		// ユースケースにリクエストを渡す
		err := h.userUsecase.SignUp(c.Request().Context(), req.IdToken, req.Name, req.Email)
		if err != nil {
			if err == repository.ErrUnauthorized {
				return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
					Status: "error",
					Error: response.ErrorDetail{
						Code:    http.StatusUnauthorized,
						Message: "invalid id token",
					},
				})
			} else if err == repository.ErrEmailExists {
				return c.JSON(http.StatusConflict, response.ErrorResponse{
					Status: "error",
					Error: response.ErrorDetail{
						Code:    http.StatusConflict,
						Message: "email already exists",
					},
				})
			}
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		return c.JSON(http.StatusCreated, response.SuccessResponse{
			Status: "success",
			Data:   nil,
		})
	}
}

// ログイン
func (h *UserHandler) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.SignInRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			})
		}

		// リクエストをユースケースに渡す
		user, err := h.userUsecase.AuthUser(c.Request().Context(), req.IdToken)
		if err != nil {
			if err == repository.ErrUserNotFound {
				return c.JSON(http.StatusNotFound, response.ErrorResponse{
					Status: "error",
					Error: response.ErrorDetail{
						Code:    http.StatusNotFound,
						Message: "user not found",
					},
				})
			}
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		// JWT 発行
		jwt, err := utils.GenerateJWT(user.Id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		res := response.SignInResponse{
			Jwt:   jwt,
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   res,
		})
	}
}

// ユーザー一覧の取得
func (h *UserHandler) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := h.userUsecase.GetAllUsers(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   users,
		})
	}
}

// ユーザー情報の取得
func (h *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Get("uid").(string)

		user, err := h.userUsecase.GetUser(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		res := response.GetUserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   res,
		})
	}
}

// ユーザー情報の更新
func (h *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.UpdateUserRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "Invalid input",
				},
			})
		}
		id := c.Get("uid").(string)

		if err := h.userUsecase.UpdateUser(c.Request().Context(), id, req.Name, req.Email); err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}
		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   nil,
		})
	}
}

// ユーザーの削除
func (h *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.DeleteUserRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "Invalid input",
				},
			})
		}

		if err := h.userUsecase.DeleteUser(c.Request().Context(), req.Id); err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}
		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   nil,
		})
	}
}
