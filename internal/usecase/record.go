package usecase

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"context"
	"time"
)

type RecordUsecase interface {
	CreateRecord(ctx context.Context, teamId, place string, date time.Time, endsData []model.DataPerEnd) (*model.Record, error)
	GetRecord(ctx context.Context, id string) (*model.Record, error)
	UpdateRecord(ctx context.Context, id, teamId, place string, date time.Time, endsData []model.DataPerEnd) error
	DeleteRecord(ctx context.Context, id string) error
}

type recordUsecase struct {
	recordRepo repository.RecordRepository
}

func NewRecordUsecase(recordRepo repository.RecordRepository) RecordUsecase {
	return &recordUsecase{recordRepo: recordRepo}
}

func (u *recordUsecase) CreateRecord(ctx context.Context, teamId, place string, date time.Time, endsData []model.DataPerEnd) (*model.Record, error) {
	return u.recordRepo.Create(ctx, teamId, place, date, endsData)
}

func (u *recordUsecase) GetRecord(ctx context.Context, id string) (*model.Record, error) {
	return u.recordRepo.GetById(ctx, id)
}

func (u *recordUsecase) UpdateRecord(ctx context.Context, id, teamId, place string, date time.Time, endsData []model.DataPerEnd) error {
	return u.recordRepo.Update(ctx, id, teamId, place, date, endsData)
}

func (u *recordUsecase) DeleteRecord(ctx context.Context, id string) error {
	return u.recordRepo.Delete(ctx, id)
}
