package handler

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	userHandler := UserHandler{userUsecase: userUsecase}
	return userHandler
}

func (handler *UserHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.User
		c.Bind(&user)
		err := handler.userUsecase.SignUp(&user)
		return c.JSON(http.StatusOK, err)
	}
}

func (h *UserHandler) SignIn(c echo.Context) error {
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

func (h *UserHandler) GetUser(c echo.Context) error {
	userID := c.Get("userID").(string) // userIDは認証ミドルウェアから取得する想定
	user, err := h.userUsecase.GetUser(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	userId := c.Get("userID").(string)
	var user model.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	user.Id = userId // 確実に認証されたユーザーのIDを使用
	if err := h.userUsecase.UpdateUser(c.Request().Context(), &user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	userID := c.Get("userID").(string)
	if err := h.userUsecase.DeleteUser(c.Request().Context(), userID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
