package handler

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/handler/request"
	"CurlARC/internal/handler/response"
	"CurlARC/internal/usecase"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RecordHandler struct {
	recordUsecase usecase.RecordUsecase
}

func NewRecordHandler(recordHandler usecase.RecordUsecase) RecordHandler {
	return RecordHandler{recordUsecase: recordHandler}
}

// CreateRecord godoc
// @Summary Create a new record
// @Description Create a new record for a team by a user
// @Tags records
// @Accept  json
// @Produce  json
// @Param teamId path string true "Team ID"
// @Param userId path string true "User ID"
// @Param record body request.CreateRecordRequest true "Record Data"
// @Success 201 {object} model.Record
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /record/{teamId}/{userId} [post]
func (h *RecordHandler) CreateRecord() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")
		userId := c.Param("userId")

		var req request.CreateRecordRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			})
		}

		// Validate JSON format
		var ends []model.DataPerEnd
		if err := json.Unmarshal([]byte(req.EndsData), &ends); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid JSON format",
				},
			})
		}

		// ユースケースにリクエストを渡す
		record, err := h.recordUsecase.CreateRecord(userId, teamId, req.Place, req.Date, req.EndsData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		// 成功時のレスポンス形式も統一
		return c.JSON(http.StatusCreated, record)
	}
}

// GetRecordByTeamId godoc
// @Summary Get records by team ID
// @Description Get all records for a specific team
// @Tags records
// @Produce  json
// @Param teamId path string true "Team ID"
// @Success 200 {object} []model.Record
// @Failure 500 {object} response.ErrorResponse
// @Router /record/{teamId} [get]
func (h *RecordHandler) GetRecordByTeamId() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")

		// ユースケースにリクエストを渡す
		record, err := h.recordUsecase.GetRecordByTeamId(teamId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		// 成功時のレスポンス形式も統一
		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   record,
		})
	}
}

// UpdateRecord godoc
// @Summary Update a record
// @Description Update a record by its ID and user ID
// @Tags records
// @Accept  json
// @Produce  json
// @Param recordId path string true "Record ID"
// @Param userId path string true "User ID"
// @Param record body request.UpdateRecordRequest true "Updated Record Data"
// @Success 200 {object} model.Record
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /record/{recordId}/{userId} [patch]
func (h *RecordHandler) UpdateRecord() echo.HandlerFunc {
	return func(c echo.Context) error {
		recordId := c.Param("recordId")
		userId := c.Param("userId")

		var req request.UpdateRecordRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			})
		}

		// Validate JSON format
		var ends []model.DataPerEnd
		if err := json.Unmarshal([]byte(req.EndsData), &ends); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid JSON format",
				},
			})
		}

		// ユースケースにリクエストを渡す
		record, err := h.recordUsecase.UpdateRecord(recordId, userId, req.Place, req.Date, req.EndsData, req.IsPublic)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		// 成功時のレスポンス形式も統一
		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   record,
		})
	}
}

// DeleteRecord godoc
// @Summary Delete a record
// @Description Delete a record by its ID
// @Tags records
// @Produce  json
// @Param recordId path string true "Record ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /record/{recordId} [delete]
func (h *RecordHandler) DeleteRecord() echo.HandlerFunc {
	return func(c echo.Context) error {
		recordId := c.Param("recordId")

		// ユースケースにリクエストを渡す
		err := h.recordUsecase.DeleteRecord(recordId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		// 成功時のレスポンス形式も統一
		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   nil,
		})
	}
}

// SetVisibility godoc
// @Summary Set record visibility
// @Description Set the visibility of a record by its ID and user ID
// @Tags records
// @Accept  json
// @Produce  json
// @Param recordId path string true "Record ID"
// @Param userId path string true "User ID"
// @Param visibility body request.SetVisibilityRequest true "Visibility Data"
// @Success 200 {object} model.Record
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /record/{recordId}/{userId}/visibility [patch]
func (h *RecordHandler) SetVisibility() echo.HandlerFunc {
	return func(c echo.Context) error {
		recordId := c.Param("recordId")
		userId := c.Param("userId")

		var req request.SetVisibilityRequest
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
		record, err := h.recordUsecase.SetVisibility(recordId, userId, req.IsPublic)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		// 成功時のレスポンス形式も統一
		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   record,
		})
	}
}
