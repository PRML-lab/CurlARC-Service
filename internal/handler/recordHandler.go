package handler

import (
	"CurlARC/internal/domain/entity"
	"CurlARC/internal/handler/request"
	"CurlARC/internal/handler/response"
	"CurlARC/internal/usecase"
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
// @Success 201 {object} entity.Record
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/record/{teamId}/{userId} [post]
func (h *RecordHandler) CreateRecord() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")
		userId := c.Get("uid").(string)

		// validate request
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

		// call usecase
		createdRecord, err := h.recordUsecase.CreateRecord(
			userId,
			teamId,
			req.EnemyTeamName,
			req.Place,
			req.Result,
			req.Date,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		// return response
		res := response.Record{
			Id:            createdRecord.GetId().Value(),
			TeamId:        createdRecord.GetTeamId(),
			Result:        createdRecord.GetResult(),
			EnemyTeamName: createdRecord.GetEnemyTeamName(),
			Place:         createdRecord.GetPlace(),
			Date:          createdRecord.GetDate(),
			EndsData:      createdRecord.GetEndsDataAsJSON(),
			IsPublic:      createdRecord.IsPublic(),
		}

		return c.JSON(http.StatusCreated, response.SuccessResponse{
			Status: "success",
			Data: struct {
				Record response.Record `json:"record"`
			}{
				Record: res,
			},
		})
	}
}

// AppendEndData godoc
// @Summary Append end data
// @Description Append end data to a record by its ID and user ID
// @Tags records
// @Accept  json
// @Produce  json
// @Param recordId path string true "Record ID"
// @Param userId path string true "User ID"
// @Param endsData body request.AppendEndDataRequest true "End Data"
// @Success 201 {object} entity.Record
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/record/{recordId}/{userId}/end [post]
func (h *RecordHandler) AppendEndData() echo.HandlerFunc {
	return func(c echo.Context) error {
		recordId := c.Param("recordId")
		userId := c.Get("uid").(string)

		// validate request
		var req request.AppendEndDataRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			})
		}

		// call usecase
		updatedRecord, err := h.recordUsecase.AppendEndData(
			recordId,
			userId,
			req.EndsData,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		// return response
		return c.JSON(http.StatusCreated, response.SuccessResponse{
			Status: "success",
			Data: struct {
				Record entity.Record `json:"record"`
			}{
				Record: *updatedRecord,
			},
		})
	}
}

func (h *RecordHandler) GetRecordDetailsByRecordId() echo.HandlerFunc {
	return func(c echo.Context) error {
		recordId := c.Param("recordId")

		record, err := h.recordUsecase.GetRecordDetailsByRecordId(recordId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		res := response.Record{
			Id:            record.GetId().Value(),
			TeamId:        record.GetTeamId(),
			Result:        record.GetResult(),
			EnemyTeamName: record.GetEnemyTeamName(),
			Place:         record.GetPlace(),
			Date:          record.GetDate(),
			EndsData:      record.GetEndsDataAsJSON(),
			IsFirst:       record.GetIsFirst(),
			IsPublic:      record.IsPublic(),
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data: struct {
				Record response.Record `json:"record"`
			}{
				Record: res,
			},
		})
	}
}

// GetRecordByTeamId godoc
// @Summary Get records by team ID
// @Description Get all records for a specific team
// @Tags records
// @Produce  json
// @Param teamId path string true "Team ID"
// @Success 200 {object} []entity.Record
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/record/{teamId} [get]
func (h *RecordHandler) GetRecordsByTeamId() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")

		// ユースケースにリクエストを渡す
		RecordIndices, err := h.recordUsecase.GetRecordIndicesByTeamId(teamId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		return c.JSON(http.StatusOK, response.GetRecordIndicesByTeamIdResponse{
			Status: "success",
			Data: struct {
				RecordIndices []response.RecordIndex `json:"record_indices"`
			}{
				RecordIndices: *RecordIndices,
			},
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
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/record/{recordId}/{userId} [patch]
func (h *RecordHandler) UpdateRecord() echo.HandlerFunc {
	return func(c echo.Context) error {
		recordId := c.Param("recordId")
		userId := c.Get("uid").(string)

		// validate request
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

		// call usecase
		updatedRecord, err := h.recordUsecase.UpdateRecord(
			recordId,
			userId,
      *req.Result,
			*req.EnemyTeamName,
			*req.Place,
			*req.EndsData,
			*req.Date,
			*req.IsFirst,
			*req.IsPublic,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		// return response
		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data: struct {
				Record entity.Record `json:"record"`
			}{
				Record: *updatedRecord,
			},
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
// @Router /auth/record/{recordId} [delete]
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
// @Success 200 {object} entity.Record
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/record/{recordId}/{userId}/visibility [patch]
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
