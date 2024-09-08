package request

import (
	entity "CurlARC/internal/domain/entity/record"
	"time"
)

type CreateRecordRequest struct {
	Result        entity.Result `json:"result"`
	EnemyTeamName string        `json:"enemy_team_name"`
	Place         string        `json:"place"`
	Date          time.Time     `json:"date"`
}

type AppendEndDataRequest struct {
	EndsData []entity.DataPerEnd `json:"ends_data"`
}

type UpdateRecordRequest = struct {
	Result        *entity.Result       `json:"result"`
	EnemyTeamName *string              `json:"enemy_team_name"`
	Place         *string              `json:"place"`
	Date          *time.Time           `json:"date"`
	EndsData      *[]entity.DataPerEnd `json:"ends_data"`
	IsPublic      *bool                `json:"is_public"`
}

type SetVisibilityRequest struct {
	IsPublic bool `json:"is_public"`
}
