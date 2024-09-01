package infra

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/handler/response"
	"time"
)

type RecordRepository struct {
	SqlHandler
}

func NewRecordRepository(sqlHandler SqlHandler) repository.RecordRepository {
	recordRepository := RecordRepository{SqlHandler: sqlHandler}
	return &recordRepository
}

func (r *RecordRepository) Save(teamId, enemyTeamName, place string, result model.Result, date time.Time) (*model.Record, error) {

	record := model.Record{
		Result:        result,
		EnemyTeamName: enemyTeamName,
		Place:         place,
		Date:          date,
		TeamId:        teamId,
		EndsData:      nil,
	}

	if err := r.Conn.Create(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *RecordRepository) FindByRecordId(id string) (*model.Record, error) {
	var record model.Record
	if err := r.Conn.First(&record, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *RecordRepository) FindIndicesByTeamId(teamId string) (*[]response.RecordIndex, error) {
	var records []model.Record
	if err := r.Conn.Select("id", "result", "enemy_team_name", "place", "date").Find(&records, "team_id = ?", teamId).Error; err != nil {
		return nil, err
	}

	var recordIndices []response.RecordIndex

	for _, record := range records {
		recordIndex := response.RecordIndex{
			Id:            record.Id,
			Result:        record.Result,
			EnemyTeamName: record.EnemyTeamName,
			Place:         record.Place,
			Date:          record.Date,
		}
		recordIndices = append(recordIndices, recordIndex)
	}

	return &recordIndices, nil
}

func (r *RecordRepository) FindByTeamId(teamId string) (*[]model.Record, error) {
	var records []model.Record
	if err := r.Conn.Find(&records, "team_id = ?", teamId).Error; err != nil {
		return nil, err
	}
	return &records, nil
}

func (r *RecordRepository) Update(recordId string, updates model.RecordUpdate) (*model.Record, error) {
	// Find the existing record
	var existingRecord model.Record
	if err := r.Conn.Where("id = ?", recordId).First(&existingRecord).Error; err != nil {
		return nil, err
	}

	// Prepare the fields to be updated
	if updates.Place != nil {
		existingRecord.Place = *updates.Place
	}
	if updates.EnemyTeamName != nil {
		existingRecord.EnemyTeamName = *updates.EnemyTeamName
	}
	if updates.Result != nil {
		existingRecord.Result = *updates.Result
	}
	if updates.Date != nil {
		existingRecord.Date = *updates.Date
	}
	if updates.EndsData != nil {
		existingRecord.EndsData = *updates.EndsData
	}
	if updates.IsPublic != nil {
		existingRecord.IsPublic = *updates.IsPublic
	}

	// Update the record with only the fields provided
	if err := r.Conn.Save(&existingRecord).Error; err != nil {
		return nil, err
	}

	// Return the updated record
	return &existingRecord, nil
}

func (r *RecordRepository) Delete(id string) error {
	if err := r.Conn.Delete(&model.Record{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
