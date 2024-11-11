package response

import (
	"CurlARC/internal/domain/entity"
	"time"

	"gorm.io/datatypes"
)

type RecordIndex struct {
	Id            string        `json:"id"`
	Result        entity.Result `json:"result"`
	EnemyTeamName string        `json:"enemy_team_name"`
	Place         string        `json:"place"`
	Date          time.Time     `json:"date"`
}

type Record struct {
	Id            string         `json:"id"`
	TeamId        string         `json:"team_id"`
	Result        entity.Result  `json:"result"`
	EnemyTeamName string         `json:"enemy_team_name"`
	Place         string         `json:"place"`
	Date          time.Time      `json:"date"`
	EndsData      datatypes.JSON `json:"ends_data"`
	IsFirst       bool           `json:"is_first"`
	IsPublic      bool           `json:"is_public"`
}

type GetRecordIndicesByTeamIdResponse struct {
	Status string `json:"status"`
	Data   struct {
		RecordIndices []RecordIndex `json:"record_indices"`
	} `json:"data"`
}
