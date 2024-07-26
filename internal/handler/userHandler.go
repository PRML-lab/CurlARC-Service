package handler

import (
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/handler/request"
	"CurlARC/internal/handler/response"
	"CurlARC/internal/usecase"
	"CurlARC/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserHandler handles requests related to users.
type UserHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler creates a new UserHandler instance.
func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return UserHandler{userUsecase: userUsecase}
}

// SignUp handles user registration.
// @Summary Register a new user
// @Description Registers a new user with the provided ID token, name, and email
// @Tags Users
// @Accept json
// @Produce json
// @Param user body request.SignUpRequest true "User registration information"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /signup [post]
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

// SignIn handles user login.
// @Summary Log in a user
// @Description Logs in a user with the provided ID token and returns a JWT
// @Tags Users
// @Accept json
// @Produce json
// @Param user body request.SignInRequest true "User login information"
// @Success 200 {object} response.SuccessResponse{data=response.SignInResponse}
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /signin [post]
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

// GetAllUsers retrieves all users.
// @Summary Get all users
// @Description Retrieves a list of all registered users
// @Tags Users
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]model.User}
// @Failure 500 {object} response.ErrorResponse
// @Router /users [get]
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

// GetUser retrieves information about a specific user.
// @Summary Get user information
// @Description Retrieves information about the currently authenticated user
// @Tags Users
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=response.GetUserResponse}
// @Failure 500 {object} response.ErrorResponse
// @Router /users/me [get]
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

// UpdateUser updates user information.
// @Summary Update user information
// @Description Updates the name and email of the currently authenticated user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body request.UpdateUserRequest true "Updated user information"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users [PATCH]
func (h *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.UpdateUserRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
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

// DeleteUser deletes a specific user.
// @Summary Delete a user
// @Description Deletes a user with the provided ID
// @Tags Users
// @Accept json
// @Produce json
// @Param user body request.DeleteUserRequest true "User ID to delete"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users [delete]
func (h *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.DeleteUserRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
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
