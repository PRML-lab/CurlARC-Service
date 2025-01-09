package handler

import (
	"CurlARC/internal/handler/request"
	"CurlARC/internal/handler/response"
	"CurlARC/internal/usecase"
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
// @Router /auth/teams [post]
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

		createdTeam, err := h.teamUsecase.CreateTeam(req.Name, userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		responseTeam := response.Team{
			Id:   createdTeam.GetId().Value(),
			Name: createdTeam.GetName(),
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data: struct {
				Team response.Team `json:"team"`
			}{
				Team: responseTeam,
			},
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
// @Router /auth/users/me/teams [get]
func (h *TeamHandler) GetTeamsByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("uid").(string)

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

		responseTeams := make([]response.Team, 0, len(teams))
		for _, team := range teams {
			responseTeams = append(responseTeams, response.Team{
				Id:   team.GetId().Value(),
				Name: team.GetName(),
			})
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data: struct {
				Teams []response.Team `json:"teams"`
			}{
				Teams: responseTeams,
			},
		})
	}
}

// GetInvitedTeams retrieves all teams that a user has been invited to.
// @Summary Get all invited teams
// @Description Retrieves a list of all teams that a user has been invited to
// @Tags Teams
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]response.Team}
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/users/me/teams/invited [get]
func (h *TeamHandler) GetInvitedTeams() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("uid").(string)

		teams, err := h.teamUsecase.GetInvitedTeams(userId)
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
				Id:   team.GetId().Value(),
				Name: team.GetName(),
			})
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data: struct {
				Teams []response.Team `json:"teams"`
			}{
				Teams: responseTeams,
			},
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
// @Router /auth/teams [get]
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
				Id:   team.GetId().Value(),
				Name: team.GetName(),
			})
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data: struct {
				Teams []response.Team
			}{
				Teams: responseTeams,
			},
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
// @Router /auth/teams/{teamId} [PATCH]
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

		updatedTeam, err := h.teamUsecase.UpdateTeam(teamId, req.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		responseTeam := response.Team{
			Id:   updatedTeam.GetId().Value(),
			Name: updatedTeam.GetName(),
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data: struct {
				Team response.Team `json:"team"`
			}{
				Team: responseTeam,
			},
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
// @Router /auth/teams/{teamId} [delete]
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
// @Router /auth/teams/{teamId}/invite/{targetId} [post]
func (h *TeamHandler) InviteUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")
		userId := c.Get("uid").(string)

		var req request.InviteUsersRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			})
		}

		err := h.teamUsecase.InviteUsers(teamId, userId, req.TargetUserEmails)
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
// @Router /auth/teams/{teamId}/accept/{userId} [post]
func (h *TeamHandler) AcceptInvitation() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")
		userId := c.Get("uid").(string)

		err := h.teamUsecase.AcceptInvitation(teamId, userId)
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
// @Router /auth/teams/{teamId}/remove/{userId} [post]
func (h *TeamHandler) RemoveMember() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")
		userId := c.Param("userId")

		err := h.teamUsecase.RemoveMember(teamId, userId)

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
// @Success 200 {object} response.SuccessResponse{data=[]entity.User}
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/teams/{teamId}/members [get]
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

		responseUsers := make([]response.User, 0, len(users))
		for _, user := range users {
			responseUsers = append(responseUsers, response.User{
				Id:    user.GetId().Value(),
				Name:  user.GetName(),
				Email: user.GetEmail(),
			})
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data: struct {
				Users []response.User `json:"users"`
			}{
				Users: responseUsers,
			},
		})
	}
}

// GetInvitedUsers retrieves all users who have been invited to a team.
// @Summary Get all invited users
// @Description Retrieves a list of all users who have been invited to a specific team
// @Tags Teams
// @Param teamId path string true "Team ID"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]entity.User}
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/teams/{teamId}/invited [get]
func (h *TeamHandler) GetInvitedUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamID := c.Param("teamId")

		users, err := h.teamUsecase.GetInvitedUsersByTeamId(teamID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		responseUsers := make([]response.User, 0, len(users))
		for _, user := range users {
			responseUsers = append(responseUsers, response.User{
				Id:    user.GetId().Value(),
				Name:  user.GetName(),
				Email: user.GetEmail(),
			})
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data: struct {
				Users []response.User `json:"users"`
			}{
				Users: responseUsers,
			},
		})
	}
}

// GetTeamDetails retrieves detailed information about a team.
// @Summary Get team details
// @Description Retrieves detailed information about a specific team
// @Tags Teams
// @Param teamId path string true "Team ID"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=response.Team}
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/teams/{teamId}/detail [get]
func (h *TeamHandler) GetTeamDetails() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")

		team, err := h.teamUsecase.GetDetailsByTeamId(teamId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		responseTeam := response.Team{
			Id:   team.GetId().Value(),
			Name: team.GetName(),
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data: struct {
				Team response.Team `json:"team"`
			}{
				Team: responseTeam,
			},
		})
	}
}
