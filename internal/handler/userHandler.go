package handler

import (
	"CurlARC/internal/handler/request"
	"CurlARC/internal/handler/response"
	"CurlARC/internal/usecase"
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

// Authorize handles user login.
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
func (h *UserHandler) Authorize() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.AuthorizeRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			})
		}

		user, err := h.userUsecase.Authorize(c, req.Name, req.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		res := response.User{
			Id:    user.GetId().Value(),
			Name:  user.GetName(),
			Email: user.GetEmail(),
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data: struct {
				User response.User `json:"user"`
			}{
				User: res,
			},
		})
	}
}

// GetAllUsers retrieves all users.
// @Summary Get all users
// @Description Retrieves a list of all registered users
// @Tags Users
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]entity.User}
// @Failure 500 {object} response.ErrorResponse
// @Router /users [get]
func (h *UserHandler) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := h.userUsecase.GetAllUsers(c)
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
// @Success 200 {object} response.SuccessResponse{data=response.User}
// @Failure 500 {object} response.ErrorResponse
// @Router /users/me [get]
func (h *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Get("uid").(string)

		user, err := h.userUsecase.GetUser(c, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		res := response.User{
			Id:    user.GetId().Value(),
			Name:  user.GetName(),
			Email: user.GetEmail(),
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

		if _, err := h.userUsecase.UpdateUser(c, id, req.Name, req.Email); err != nil {
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

		if err := h.userUsecase.DeleteUser(c, req.Id); err != nil {
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
