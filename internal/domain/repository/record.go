package repository

import (
	"context"

	"CurlARC/internal/domain/model"
)

type RecordRepository interface {
	Create(ctx context.Context, record *model.Record) error
	GetById(ctx context.Context, Id string) (*model.Record, error)
	Update(ctx context.Context, record *model.Record) error
	Delete(ctx context.Context, Id string) error
}

type EndRepository interface {
	Create(ctx context.Context, end *model.End) error
	GetById(ctx context.Context, Id string) (*model.End, error)
	Update(ctx context.Context, end *model.End) error
	Delete(ctx context.Context, Id string) error
}

type ShotRepository interface {
	Create(ctx context.Context, shot *model.Shot) error
	GetById(ctx context.Context, Id string) (*model.Shot, error)
	Update(ctx context.Context, shot *model.Shot) error
	Delete(ctx context.Context, Id string) error
}
