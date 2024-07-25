package usecase

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"context"
	"time"

	"gorm.io/datatypes"
)

type RecordUsecase interface {
	CreateRecord(ctx context.Context, userId, teamId, place string, date time.Time, endsData datatypes.JSON) (*model.Record, error)
	GetRecordByTeamId(ctx context.Context, id string) (*model.Record, error)
	UpdateRecord(ctx context.Context, id, teamId, place string, date time.Time, endsData datatypes.JSON) (*model.Record, error)
	DeleteRecord(ctx context.Context, id string) error
}

type recordUsecase struct {
	recordRepo   repository.RecordRepository
	userTeamRepo repository.UserTeamRepository
	teamRepo     repository.TeamRepository
}

func NewRecordUsecase(recordRepo repository.RecordRepository, userTeamRepo repository.UserTeamRepository, teamRepo repository.TeamRepository) RecordUsecase {
	return &recordUsecase{recordRepo: recordRepo, userTeamRepo: userTeamRepo, teamRepo: teamRepo}
}

func (u *recordUsecase) CreateRecord(ctx context.Context, userId, teamId, place string, date time.Time, endsData datatypes.JSON) (*model.Record, error) {

	// check if the user is a member of the team
	if _, err := u.userTeamRepo.IsMember(userId, teamId); err != nil {
		return nil, err
	}

	// check if the team exists
	if _, err := u.teamRepo.FindById(teamId); err != nil {
		return nil, err
	}

	return u.recordRepo.Create(ctx, teamId, place, date, endsData)
}

func (u *recordUsecase) GetRecordByTeamId(ctx context.Context, teamId string) (*model.Record, error) {
	return u.recordRepo.GetByTeamId(ctx, teamId)
}

func (u *recordUsecase) UpdateRecord(ctx context.Context, id, teamId, place string, date time.Time, endsData datatypes.JSON) (*model.Record, error) {
	return u.recordRepo.Update(ctx, id, teamId, place, date, endsData)
}

func (u *recordUsecase) DeleteRecord(ctx context.Context, id string) error {
	return u.recordRepo.Delete(ctx, id)
}
