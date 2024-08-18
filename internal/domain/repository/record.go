package repository

import (
	"time"

	"CurlARC/internal/domain/model"
)

type RecordRepository interface {
	Create(teamId, enemyTeamName, place string, result model.Result, date time.Time) (*model.Record, error)
	FindByRecordId(recordId string) (*model.Record, error)
	FindByTeamId(teamId string) (*[]model.Record, error)
	Update(recordId string, updates model.RecordUpdate) (*model.Record, error)
	Delete(recordId string) error
}
