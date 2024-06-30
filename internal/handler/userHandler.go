package handler

import (
	"CurlARC/internal/domain/model"
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

func (h *UserHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.User
		c.Bind(&user)
		err := h.userUsecase.SignUp(&user)
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
	userID := c.Get("userID").(string)
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
	user.Id = userId
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
