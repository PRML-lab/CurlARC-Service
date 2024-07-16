package repository

import (
	"context"
	"time"

	"CurlARC/internal/domain/model"
)

type RecordRepository interface {
	Create(ctx context.Context, teamId, place string, date time.Time, endsData []model.DataPerEnd) (*model.Record, error)
	GetById(ctx context.Context, id string) (*model.Record, error)
	Update(ctx context.Context, id, teamId, place string, date time.Time, endsData []model.DataPerEnd) error
	Delete(ctx context.Context, id string) error
}
