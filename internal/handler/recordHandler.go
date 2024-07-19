package handler

import (
	"CurlARC/internal/handler/request"
	"CurlARC/internal/usecase"
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

		// ユースケースにリクエストを渡す
		record, err := h.recordUsecase.CreateRecord(c.Request().Context(), req.TeamId, req.Place, req.Date, req.EndsData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, record)
	}
}
