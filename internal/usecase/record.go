package usecase

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"errors"
	"time"

	"gorm.io/datatypes"
)

type RecordUsecase interface {
	CreateRecord(userId, teamId, place string, date time.Time, endsData datatypes.JSON) (*model.Record, error)
	GetRecordByTeamId(teamId string) (*model.Record, error)
	UpdateRecord(recordId, userId, place string, date time.Time, endsData datatypes.JSON, isPublic bool) (*model.Record, error)
	DeleteRecord(id string) error

	SetVisibility(recordId, userId string, isPublic bool) (*model.Record, error)
}

type recordUsecase struct {
	recordRepo   repository.RecordRepository
	userTeamRepo repository.UserTeamRepository
	teamRepo     repository.TeamRepository
}

func NewRecordUsecase(recordRepo repository.RecordRepository, userTeamRepo repository.UserTeamRepository, teamRepo repository.TeamRepository) RecordUsecase {
	return &recordUsecase{recordRepo: recordRepo, userTeamRepo: userTeamRepo, teamRepo: teamRepo}
}

func (u *recordUsecase) CreateRecord(userId, teamId, place string, date time.Time, endsData datatypes.JSON) (*model.Record, error) {

	// check if the user is a member of the team
	if _, err := u.userTeamRepo.IsMember(userId, teamId); err != nil {
		return nil, err
	}

	// check if the team exists
	if _, err := u.teamRepo.FindById(teamId); err != nil {
		return nil, err
	}

	return u.recordRepo.Create(teamId, place, date, endsData)
}

func (u *recordUsecase) GetRecordByTeamId(teamId string) (*model.Record, error) {
	return u.recordRepo.FindByTeamId(teamId)
}

func (u *recordUsecase) UpdateRecord(recordId, userId, place string, date time.Time, endsData datatypes.JSON, isPublic bool) (*model.Record, error) {
	// get the team id of the record
	record, err := u.recordRepo.FindById(recordId)
	if err != nil {
		return nil, err
	}

	// check if the user is the member of the record
	isMember, err := u.userTeamRepo.IsMember(userId, record.TeamId)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("inviter is not a member of the team")
	}

	return u.recordRepo.Update(recordId, place, date, endsData, isPublic)
}

func (u *recordUsecase) DeleteRecord(id string) error {
	return u.recordRepo.Delete(id)
}

func (u *recordUsecase) SetVisibility(recordId, userId string, isPublic bool) (*model.Record, error) {

	// check if the record exists
	record, err := u.recordRepo.FindById(recordId)
	if err != nil {
		return nil, err
	}

	// check if the user is the member of the record
	isMember, err := u.userTeamRepo.IsMember(userId, record.TeamId)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("inviter is not a member of the team")
	}

	return u.recordRepo.Update(recordId, record.Place, record.Date, record.EndsData, isPublic)
}
