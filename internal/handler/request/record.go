package request

import (
	"time"

	"gorm.io/datatypes"
)

type CreateRecordRequest struct {
	UserId   string         `json:"user_id"`
	TeamId   string         `json:"team_id"`
	Place    string         `json:"place"`
	Date     time.Time      `json:"date"`
	EndsData datatypes.JSON `json:"ends_data"`
}
