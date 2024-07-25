package repository

import (
	"context"
	"time"

	"CurlARC/internal/domain/model"

	"gorm.io/datatypes"
)

type RecordRepository interface {
	Create(ctx context.Context, teamId, place string, date time.Time, endsData datatypes.JSON) (*model.Record, error)
	GetById(ctx context.Context, id string) (*model.Record, error)
	Update(ctx context.Context, id, teamId, place string, date time.Time, endsData datatypes.JSON) (*model.Record, error)
	Delete(ctx context.Context, id string) error
}
