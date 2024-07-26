package repository

import (
	"time"

	"CurlARC/internal/domain/model"

	"gorm.io/datatypes"
)

type RecordRepository interface {
	Create(teamId, place string, date time.Time, endsData datatypes.JSON) (*model.Record, error)
	FindById(recordId string) (*model.Record, error)
	FindByTeamId(teamId string) (*model.Record, error)
	Update(recordId, place string, date time.Time, endsData datatypes.JSON, isPulbic bool) (*model.Record, error)
	Delete(recordId string) error
}
