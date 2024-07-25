package handler

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/handler/request"
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
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}
		// Validate JSON format
		var ends []model.DataPerEnd
		if err := json.Unmarshal([]byte(req.EndsData), &ends); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid JSON format"})
		}

		// ユースケースにリクエストを渡す
		record, err := h.recordUsecase.CreateRecord(c.Request().Context(), req.UserId, req.TeamId, req.Place, req.Date, req.EndsData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

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
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, record)
	}
}
