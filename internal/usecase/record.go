package usecase

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/handler/response"
	"errors"
	"fmt"
	"time"
)

type RecordUsecase interface {
	CreateRecord(userId, teamId, enemyTeamName, place string, result model.Result, date time.Time) (*model.Record, error) // Create a new record which has no endsData
	AppendEndData(recordId, userId string, endsData []model.DataPerEnd) (*model.Record, error)                            // Append endsData to an existing record
	GetRecordDetailsByRecordId(recordId string) (*model.Record, error)
	GetRecordIndicesByTeamId(teamId string) (*[]response.RecordIndex, error)
	GetRecordsByTeamId(teamId string) (*[]model.Record, error)
	UpdateRecord(recordId, userId string, updates model.RecordUpdate) (*model.Record, error)
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

func (u *recordUsecase) CreateRecord(userId, teamId, enemyTeamName, place string, result model.Result, date time.Time) (*model.Record, error) {

	// check if the user is a member of the team
	isMember, err := u.userTeamRepo.IsMember(userId, teamId)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("user is not a member of the team")
	}

	// check if the team exists
	if _, err := u.teamRepo.FindById(teamId); err != nil {
		return nil, err
	}

	fmt.Print("usecase", enemyTeamName)

	return u.recordRepo.Save(teamId, enemyTeamName, place, result, date)
}

func (u *recordUsecase) AppendEndData(recordId, userId string, endsData []model.DataPerEnd) (*model.Record, error) {
	// Get the record by ID
	currentRecord, err := u.recordRepo.FindByRecordId(recordId)
	if err != nil {
		return nil, err
	}

	// Check if the user is a member of the team
	isMember, err := u.userTeamRepo.IsMember(userId, currentRecord.TeamId)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("appender is not a member of the team")
	}

	updatedRecord := currentRecord
	updatedRecord.AppendEndsData(endsData)

	// Update the record with the new endsData
	updatedRecord, err := u.recordRepo.Update()
	if err != nil {
		return nil, err
	}

	return updatedRecord, nil
}

func (u *recordUsecase) GetRecordDetailsByRecordId(recordId string) (*model.Record, error) {
	return u.recordRepo.FindByRecordId(recordId)
}

func (u *recordUsecase) GetRecordIndicesByTeamId(teamId string) (*[]response.RecordIndex, error) {
	return u.recordRepo.FindIndicesByTeamId(teamId)
}

func (u *recordUsecase) GetRecordsByTeamId(teamId string) (*[]model.Record, error) {
	return u.recordRepo.FindByTeamId(teamId)
}

func (u *recordUsecase) UpdateRecord(recordId, userId string, updates model.RecordUpdate) (*model.Record, error) {
	// Get the record by ID
	record, err := u.recordRepo.FindByRecordId(recordId)
	if err != nil {
		return nil, err
	}

	// Check if the user is a member of the team
	isMember, err := u.userTeamRepo.IsMember(userId, record.TeamId)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("updater is not a member of the team")
	}

	// Update the record with only the fields provided in the updates
	updatedRecord, err := u.recordRepo.Update(recordId, updates)
	if err != nil {
		return nil, err
	}

	return updatedRecord, nil
}

func (u *recordUsecase) DeleteRecord(id string) error {
	return u.recordRepo.Delete(id)
}

func (u *recordUsecase) SetVisibility(recordId, userId string, isPublic bool) (*model.Record, error) {

	// check if the record exists
	record, err := u.recordRepo.FindByRecordId(recordId)
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

	// Prepare the update struct
	updateFields := model.RecordUpdate{
		IsPublic: &isPublic,
	}

	return u.recordRepo.Update(recordId, updateFields)
}
