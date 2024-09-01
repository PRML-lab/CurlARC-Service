package repository

import (
	"time"

	"CurlARC/internal/domain/model"
	"CurlARC/internal/handler/response"
)

type RecordRepository interface {
	Save(teamId, enemyTeamName, place string, result model.Result, date time.Time) (*model.Record, error)
	FindByRecordId(recordId string) (*model.Record, error)
	FindIndicesByTeamId(teamId string) (*[]response.RecordIndex, error)
	FindByTeamId(teamId string) (*[]model.Record, error)
	Update(recordId string, updates model.RecordUpdate) (*model.Record, error)
	Delete(recordId string) error
}
