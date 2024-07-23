package handler

import (
	"CurlARC/internal/handler/request"
	"CurlARC/internal/handler/response"
	"CurlARC/internal/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type TeamHandler struct {
	teamUsecase usecase.TeamUsecase
}

func NewTeamHandler(teamUsecase usecase.TeamUsecase) TeamHandler {
	return TeamHandler{teamUsecase: teamUsecase}
}

func (h *TeamHandler) CreateTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		// 認証済みユーザーのIDを取得
		userId := c.Get("uid").(string)
		var req request.CreateTeamRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			})
		}

		err := h.teamUsecase.CreateTeam(req.Name, userId)
		if err != nil {
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

func (h *TeamHandler) GetAllTeams() echo.HandlerFunc {
	return func(c echo.Context) error {
		teams, err := h.teamUsecase.GetAllTeams()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		responseTeams := make([]response.Team, 0, len(teams))
		for _, team := range teams {
			responseTeams = append(responseTeams, response.Team{
				Id:   team.Id,
				Name: team.Name,
			})
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   responseTeams,
		})
	}
}

func (h *TeamHandler) GetTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		team, err := h.teamUsecase.GetTeam(id)
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
			Data:   team,
		})
	}
}

func (h *TeamHandler) UpdateTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var req request.UpdateTeamRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			})
		}

		err := h.teamUsecase.UpdateTeam(id, req.Name)
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
			Data:   nil,
		})
	}
}

func (h *TeamHandler) DeleteTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		err := h.teamUsecase.DeleteTeam(id)
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
			Data:   nil,
		})
	}
}

func (h *TeamHandler) InviteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamID := c.Param("teamId")
		userID := c.Param("userId")

		err := h.teamUsecase.InviteUser(teamID, userID)
		if err != nil {
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

func (h *TeamHandler) AcceptInvitation() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamID := c.Param("teamId")
		userID := c.Param("userId")

		err := h.teamUsecase.AcceptInvitation(teamID, userID)
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
			Data:   nil,
		})
	}
}

func (h *TeamHandler) RemoveMember() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamID := c.Param("teamId")
		userID := c.Param("userId")

		err := h.teamUsecase.RemoveMember(teamID, userID)
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
			Data:   nil,
		})
	}
}

func (h *TeamHandler) GetMembers() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamID := c.Param("teamId")

		users, err := h.teamUsecase.GetMembersByTeamId(teamID)
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

func (h *TeamHandler) GetTeamsByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("uid").(string)

		teams, err := h.teamUsecase.GetTeamsByUserId(userID)
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
			Data:   teams,
		})
	}
}
