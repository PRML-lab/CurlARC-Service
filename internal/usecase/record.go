package usecase

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"context"
	"time"
)

type RecordUsecase interface {
	// CRUD
	CreateRecord(
		ctx context.Context,
		teamId string,
		place string,
		date time.Time,
	) (*model.Record, error)

	GetRecord(
		ctx context.Context,
		Id string,
	) (*model.Record, error)

	GetRecordsByTeamId(
		ctx context.Context,
		teamId string,
	) ([]*model.Record, error)

	UpdateRecord(
		ctx context.Context,
		Id string,
		teamId string,
		place string,
		date time.Time,
	) error

	DeleteRecord(
		ctx context.Context,
		Id string,
	) error
}

type recordUsecase struct {
	recordRepo repository.RecordRepository
	endRepo    repository.EndRepository
	shotRepo   repository.ShotRepository
	coordRepo  repository.CoordinateRepository
}

func NewRecordUsecase(
	recordRepo repository.RecordRepository,
	endRepo repository.EndRepository,
	shotRepo repository.ShotRepository,
	coordRepo repository.CoordinateRepository,
) RecordUsecase {
	return &recordUsecase{
		recordRepo: recordRepo,
		endRepo:    endRepo,
		shotRepo:   shotRepo,
		coordRepo:  coordRepo,
	}
}

func (r *recordUsecase) CreateRecord(ctx context.Context, teamId string, place string, date time.Time) (*model.Record, error) {
	record := &model.Record{
		TeamId: teamId,
		Place:  place,
		Date:   date,
	}
	if err := r.recordRepo.Create(ctx, record); err != nil {
		return nil, err
	}
	return record, nil
}

func (r *recordUsecase) GetRecord(ctx context.Context, recordId string) (*model.Record, error) {
	record, err := r.recordRepo.GetById(ctx, recordId)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (r *recordUsecase) GetRecordsByTeamId(ctx context.Context, teamId string) ([]*model.Record, error) {
	return r.recordRepo.GetByTeamId(ctx, teamId)
}

func (r *recordUsecase) UpdateRecord(ctx context.Context, recordId string, teamId string, place string, date time.Time) error {
	record, err := r.recordRepo.GetById(ctx, recordId)
	if err != nil {
		return err
	}
	record.TeamId = teamId
	record.Place = place
	record.Date = date
	if err := r.recordRepo.Update(ctx, record); err != nil {
		return err
	}
	return nil
}

func (r *recordUsecase) DeleteRecord(ctx context.Context, recordId string) error {
	return r.recordRepo.Delete(ctx, recordId)
}
