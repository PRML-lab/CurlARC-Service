package request

import (
	"CurlARC/internal/domain/model"
	"time"
)

type CreateRecordRequest struct {
	Result        model.Result `json:"result"`
	EnemyTeamName string       `json:"enemy_team_name"`
	Place         string       `json:"place"`
	Date          time.Time    `json:"date"`
}

type AppendEndDataRequest struct {
	EndsData []model.DataPerEnd `json:"ends_data"`
}

type UpdateRecordRequest = struct {
	Result        *model.Result       `json:"result"`
	EnemyTeamName *string             `json:"enemy_team_name"`
	Place         *string             `json:"place"`
	Date          *time.Time          `json:"date"`
	EndsData      *[]model.DataPerEnd `json:"ends_data"`
	IsPublic      *bool               `json:"is_public"`
}

type SetVisibilityRequest struct {
	IsPublic bool `json:"is_public"`
}
