package usecase

import (
	"CurlARC/internal/domain/entity"
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/handler/response"
	"errors"
	"time"
)

type RecordUsecase interface {
	CreateRecord(userId, teamId, enemyTeamName, place string, result entity.Result, date time.Time) (*entity.Record, error) // Create a new record which has no endsData
	AppendEndData(recordId, userId string, endsData []entity.DataPerEnd) (*entity.Record, error)                            // Append endsData to an existing record
	GetRecordDetailsByRecordId(recordId string) (*entity.Record, error)
	GetRecordIndicesByTeamId(teamId string) (*[]response.RecordIndex, error)
	GetRecordsByTeamId(teamId string) (*[]entity.Record, error)
	UpdateRecord(recordId, userId string, result entity.Result, enemyTeamName, place string, endsData []entity.DataPerEnd, date time.Time, isRed bool, isFirst bool, isPublic bool) (*entity.Record, error)
	DeleteRecord(id string) error
	SetVisibility(recordId, userId string, isPublic bool) (*entity.Record, error)
}

type recordUsecase struct {
	recordRepo   repository.RecordRepository
	userTeamRepo repository.UserTeamRepository
	teamRepo     repository.TeamRepository
}

func NewRecordUsecase(recordRepo repository.RecordRepository, userTeamRepo repository.UserTeamRepository, teamRepo repository.TeamRepository) RecordUsecase {
	return &recordUsecase{recordRepo: recordRepo, userTeamRepo: userTeamRepo, teamRepo: teamRepo}
}

func (u *recordUsecase) CreateRecord(userId, teamId, enemyTeamName, place string, result entity.Result, date time.Time) (*entity.Record, error) {

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

	// Create a new record
	record, err := entity.NewRecord(
		teamId,
		entity.WithEnemyTeamName(enemyTeamName),
		entity.WithResult(result),
		entity.WithPlace(place),
		entity.WithDate(date),
	)
	if err != nil {
		return nil, err
	}

	// Save the record
	savedRecord, err := u.recordRepo.Save(*record)

	return savedRecord, err
}

func (u *recordUsecase) AppendEndData(recordId, userId string, endsData []entity.DataPerEnd) (*entity.Record, error) {

	// Get the record by ID
	currentRecord, err := u.recordRepo.FindByRecordId(recordId)
	if err != nil {
		return nil, err
	}

	// Check if the user is a member of the team
	isMember, err := u.userTeamRepo.IsMember(userId, currentRecord.GetTeamId())
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("appender is not a member of the team")
	}

	// Append the new endsData to the record
	newEndsData := append(currentRecord.GetEndsData(), endsData...)
	err = currentRecord.ValidateEndsData(newEndsData)
	if err != nil {
		return nil, err
	}
	newRecord := currentRecord
	err = newRecord.SetEndsData(newEndsData)
	if err != nil {
		return nil, err
	}

	// Update the record with the new endsData
	updatedRecord, err := u.recordRepo.Update(*newRecord)
	if err != nil {
		return nil, err
	}

	return updatedRecord, nil
}

func (u *recordUsecase) GetRecordDetailsByRecordId(recordId string) (*entity.Record, error) {
	return u.recordRepo.FindByRecordId(recordId)
}

func (u *recordUsecase) GetRecordIndicesByTeamId(teamId string) (*[]response.RecordIndex, error) {
	return u.recordRepo.FindIndicesByTeamId(teamId)
}

func (u *recordUsecase) GetRecordsByTeamId(teamId string) (*[]entity.Record, error) {
	return u.recordRepo.FindByTeamId(teamId)
}

func (u *recordUsecase) UpdateRecord(recordId, userId string, result entity.Result, enemyTeamName, place string, endsData []entity.DataPerEnd, date time.Time, isRed bool, isFirst, isPublic bool) (*entity.Record, error) {

	// Get the record by ID
	record, err := u.recordRepo.FindByRecordId(recordId)
	if err != nil {
		return nil, err
	}

	// Check if the user is a member of the team
	isMember, err := u.userTeamRepo.IsMember(userId, record.GetTeamId())
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("updater is not a member of the team")
	}

	// Prepare the update struct
	newRecord := record

  if result != "" {
    newRecord.SetResult(result)
  }
	if enemyTeamName != "" {
		newRecord.SetEnemyTeamName(enemyTeamName)
	}
	if place != "" {
		newRecord.SetPlace(place)
	}
	if !date.IsZero() {
		newRecord.SetDate(date)
	}
	if len(endsData) > 0 {
		err = newRecord.ValidateEndsData(endsData)
		if err != nil {
			return nil, err
		}
		err = newRecord.SetEndsData(endsData)
		if err != nil {
			return nil, err
		}
	}
	newRecord.SetIsRed(isRed)
	newRecord.SetIsFirst(isFirst)
	newRecord.SetVisibility(isPublic)

	// Update the record with only the fields provided in the updates
	updatedRecord, err := u.recordRepo.Update(*newRecord)
	if err != nil {
		return nil, err
	}

	return updatedRecord, nil
}

func (u *recordUsecase) DeleteRecord(id string) error {
	return u.recordRepo.Delete(id)
}

func (u *recordUsecase) SetVisibility(recordId, userId string, isPublic bool) (*entity.Record, error) {

	// check if the record exists
	record, err := u.recordRepo.FindByRecordId(recordId)
	if err != nil {
		return nil, err
	}

	// check if the user is the member of the record
	isMember, err := u.userTeamRepo.IsMember(userId, record.GetTeamId())
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("inviter is not a member of the team")
	}

	// update the record
	newRecord := record
	newRecord.SetVisibility(isPublic)

	updatedRecord, err := u.recordRepo.Update(*newRecord)
	if err != nil {
		return nil, err
	}

	return updatedRecord, nil
}
