package request

import (
	"CurlARC/internal/domain/model"
	"time"

	"gorm.io/datatypes"
)

type CreateRecordRequest struct {
	Result        model.Result `json:"result"`
	EnemyTeamName string       `json:"enemy_team_name"`
	Place         string       `json:"place"`
	Date          time.Time    `json:"date"`
}

type AppendEndDataRequest struct {
	EndsData datatypes.JSON `json:"ends_data"`
}

type UpdateRecordRequest = model.RecordUpdate

type SetVisibilityRequest struct {
	IsPublic bool `json:"is_public"`
}
