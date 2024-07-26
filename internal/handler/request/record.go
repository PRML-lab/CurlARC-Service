package request

import (
	"time"

	"gorm.io/datatypes"
)

type CreateRecordRequest struct {
	Place    string         `json:"place"`
	Date     time.Time      `json:"date"`
	EndsData datatypes.JSON `json:"ends_data"`
}

type UpdateRecordRequest struct {
	RecordId string         `json:"user_id"`
	Place    string         `json:"place"`
	Date     time.Time      `json:"date"`
	EndsData datatypes.JSON `json:"ends_data"`
	IsPublic bool           `json:"is_public"`
}

type SetVisibilityRequest struct {
	IsPublic bool `json:"is_public"`
}
