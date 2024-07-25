package handler

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/handler/request"
	"CurlARC/internal/handler/response"
	"CurlARC/internal/usecase"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type RecordHandler struct {
	recordUsecase usecase.RecordUsecase
}

func NewRecordHandler(recordHandler usecase.RecordUsecase) RecordHandler {
	return RecordHandler{recordUsecase: recordHandler}
}

// レコード作成
func (h *RecordHandler) CreateRecord() echo.HandlerFunc {
	return func(c echo.Context) error {
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
		record, err := h.recordUsecase.CreateRecord(c.Request().Context(), req.UserId, req.TeamId, req.Place, req.Date, req.EndsData)
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

// レコード取得
func (h *RecordHandler) GetRecordByTeamId() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("teamId")

		// ユースケースにリクエストを渡す
		record, err := h.recordUsecase.GetRecordByTeamId(c.Request().Context(), teamId)
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
