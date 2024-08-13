package handler

import (
	"CurlARC/internal/handler/request"
	"CurlARC/internal/handler/response"
	"CurlARC/internal/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TeamHandler handles requests related to teams.
type TeamHandler struct {
	teamUsecase usecase.TeamUsecase
}

// NewTeamHandler creates a new TeamHandler instance.
func NewTeamHandler(teamUsecase usecase.TeamUsecase) TeamHandler {
	return TeamHandler{teamUsecase: teamUsecase}
}

// CreateTeam creates a new team.
// @Summary Create a new team
// @Description Creates a new team with the specified name
// @Tags Teams
// @Accept json
// @Produce json
// @Param team body request.CreateTeamRequest true "Team information"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /teams [post]
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

// GetTeamsByUserId retrieves all teams for a specific user.
// @Summary Get all teams for a user
// @Description Retrieves a list of all teams associated with a specific user
// @Tags Teams
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]response.Team}
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/teams/{userId} [get]
func (h *TeamHandler) GetTeamsByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("uid").(string)
		fmt.Print(userId)

		teams, err := h.teamUsecase.GetTeamsByUserId(userId)
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

// GetAllTeams retrieves all teams.
// @Summary Get all teams
// @Description Retrieves a list of all teams
// @Tags Teams
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]response.Team}
// @Failure 500 {object} response.ErrorResponse
// @Router /teams [get]
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

// UpdateTeam updates an existing team.
// @Summary Update a team
// @Description Updates the name of an existing team
// @Tags Teams
// @Accept json
// @Produce json
// @Param id path string true "Team ID"
// @Param team body request.UpdateTeamRequest true "Updated team information"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /teams/{teamId} [PATCH]
func (h *TeamHandler) UpdateTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")
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

		err := h.teamUsecase.UpdateTeam(teamId, req.Name)
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

// DeleteTeam deletes a specific team by ID.
// @Summary Delete a team
// @Description Deletes a team by its ID
// @Tags Teams
// @Param id path string true "Team ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /teams/{teamId} [delete]
func (h *TeamHandler) DeleteTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")
		err := h.teamUsecase.DeleteTeam(teamId)
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

// InviteUser invites a user to a team.
// @Summary Invite a user to a team
// @Description Invites a user to a specific team
// @Tags Teams
// @Param teamId path string true "Team ID"
// @Param targetId path string true "Target User ID"
// @Success 201 {object} response.SuccessResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /teams/{teamId}/invite/{targetId} [post]
func (h *TeamHandler) InviteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")
		userId := c.Get("uid").(string)
		targetId := c.Param("targetId")

		err := h.teamUsecase.InviteUser(teamId, userId, targetId)
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

// AcceptInvitation accepts an invitation to join a team.
// @Summary Accept a team invitation
// @Description Accepts an invitation to join a specific team
// @Tags Teams
// @Param teamId path string true "Team ID"
// @Param userId path string true "User ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /teams/{teamId}/accept/{userId} [post]
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

// RemoveMember removes a member from a team.
// @Summary Remove a member from a team
// @Description Removes a member from a specific team
// @Tags Teams
// @Param teamId path string true "Team ID"
// @Param userId path string true "User ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /teams/{teamId}/remove/{userId} [post]
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

// GetMembers retrieves all members of a team.
// @Summary Get all members of a team
// @Description Retrieves a list of all members of a specific team
// @Tags Teams
// @Param teamId path string true "Team ID"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]model.User}
// @Failure 500 {object} response.ErrorResponse
// @Router /teams/{teamId}/members [get]
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