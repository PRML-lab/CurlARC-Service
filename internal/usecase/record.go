package usecase

import (
	entity "CurlARC/internal/domain/entity/record"
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/handler/response"
	"errors"
	"time"
)

type RecordUsecase interface {
	CreateRecord(userId, teamId, enemyTeamName, place string, result entity.Result, date time.Time) error // Create a new record which has no endsData
	AppendEndData(recordId, userId string, endsData []entity.DataPerEnd) error                            // Append endsData to an existing record
	GetRecordDetailsByRecordId(recordId string) (*entity.Record, error)
	GetRecordIndicesByTeamId(teamId string) (*[]response.RecordIndex, error)
	GetRecordsByTeamId(teamId string) (*[]entity.Record, error)
	UpdateRecord(recordId, userId, enemyTeamName, place string, endsData []entity.DataPerEnd, date time.Time) error
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

func (u *recordUsecase) CreateRecord(userId, teamId, enemyTeamName, place string, result entity.Result, date time.Time) error {

	// check if the user is a member of the team
	isMember, err := u.userTeamRepo.IsMember(userId, teamId)
	if err != nil {
		return err
	}
	if !isMember {
		return errors.New("user is not a member of the team")
	}

	// check if the team exists
	if _, err := u.teamRepo.FindById(teamId); err != nil {
		return err
	}

	// Create a new record
	record, err := entity.NewRecord(
		teamId,
		entity.WithEnemyTeamName(enemyTeamName),
		entity.WithPlace(place),
		entity.WithDate(date),
	)
	if err != nil {
		return err
	}

	// Save the record
	err = u.recordRepo.Save(*record)

	return err
}

func (u *recordUsecase) AppendEndData(recordId, userId string, endsData []entity.DataPerEnd) error {

	// Get the record by ID
	currentRecord, err := u.recordRepo.FindByRecordId(recordId)
	if err != nil {
		return err
	}

	// Check if the user is a member of the team
	isMember, err := u.userTeamRepo.IsMember(userId, currentRecord.GetTeamId())
	if err != nil {
		return err
	}
	if !isMember {
		return errors.New("appender is not a member of the team")
	}

	// Append the new endsData to the record
	newEndsData := append(currentRecord.GetEndsData(), endsData...)
	err = currentRecord.ValidateEndsData(newEndsData)
	if err != nil {
		return err
	}
	newRecord := currentRecord
	err = newRecord.SetEndsData(newEndsData)
	if err != nil {
		return err
	}

	// Update the record with the new endsData
	err = u.recordRepo.Update(*newRecord)
	if err != nil {
		return err
	}

	return nil
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

func (u *recordUsecase) UpdateRecord(recordId, userId, enemyTeamName, place string, endsData []entity.DataPerEnd, date time.Time) error {

	// Get the record by ID
	record, err := u.recordRepo.FindByRecordId(recordId)
	if err != nil {
		return err
	}

	// Check if the user is a member of the team
	isMember, err := u.userTeamRepo.IsMember(userId, record.GetTeamId())
	if err != nil {
		return err
	}
	if !isMember {
		return errors.New("updater is not a member of the team")
	}

	// Prepare the update struct
	newRecord := record

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
			return err
		}
		err = newRecord.SetEndsData(endsData)
		if err != nil {
			return err
		}
	}

	// Update the record with only the fields provided in the updates
	err = u.recordRepo.Update(*newRecord)
	if err != nil {
		return err
	}

	return nil
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

	err = u.recordRepo.Update(*newRecord)
	if err != nil {
		return nil, err
	}

	return newRecord, nil
}
