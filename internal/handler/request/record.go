package request

import (
	"time"

	"gorm.io/datatypes"
)

type CreateRecordRequest struct {
	TeamId   string         `json:"team_id"`
	Place    string         `json:"place"`
	Date     time.Time      `json:"date"`
	EndsData datatypes.JSON `json:"ends_data"`
}
