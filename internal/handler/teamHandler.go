package handler

import (
	"CurlARC/internal/handler/request"
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
		var req request.CreateTeamRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		err := h.teamUsecase.CreateTeam(req.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.NoContent(http.StatusCreated)
	}
}

func (h *TeamHandler) GetAllTeams() echo.HandlerFunc {
	return func(c echo.Context) error {
		teams, err := h.teamUsecase.GetAllTeams()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, teams)
	}
}

func (h *TeamHandler) GetTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		team, err := h.teamUsecase.GetTeam(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, team)
	}
}

func (h *TeamHandler) UpdateTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var req request.UpdateTeamRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		err := h.teamUsecase.UpdateTeam(id, req.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.NoContent(http.StatusOK)
	}
}

func (h *TeamHandler) DeleteTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		err := h.teamUsecase.DeleteTeam(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.NoContent(http.StatusOK)
	}
}

func (h *TeamHandler) AddMember() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")
		userId := c.Param("userId")

		err := h.teamUsecase.AddMember(teamId, userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.NoContent(http.StatusCreated)
	}
}

func (h *TeamHandler) RemoveMember() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")
		userId := c.Param("userId")

		err := h.teamUsecase.RemoveMember(teamId, userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.NoContent(http.StatusOK)
	}
}
