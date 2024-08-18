package response

import (
	"CurlARC/internal/domain/model"
	"time"
)

type RecordIndex struct {
	Id            string       `json:"id"`
	Result        model.Result `json:"result"`
	EnemyTeamName string       `json:"enemy_team_name"`
	Place         string       `json:"place"`
	Date          time.Time    `json:"date"`
}

type GetRecordIndicesByTeamIdRespone struct {
	Status string `json:"status"`
	Data   struct {
		RecordIndices []RecordIndex `json:"record_indices"`
	} `json:"data"`
}
