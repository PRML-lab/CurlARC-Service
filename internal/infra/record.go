package infra

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"time"

	"gorm.io/datatypes"
)

type RecordRepository struct {
	SqlHandler
}

func NewRecordRepository(sqlHandler SqlHandler) repository.RecordRepository {
	recordRepository := RecordRepository{SqlHandler: sqlHandler}
	return &recordRepository
}

func (r *RecordRepository) Create(teamId, place string, date time.Time, endsData datatypes.JSON) (*model.Record, error) {

	record := &model.Record{
		Place:    place,
		Date:     date,
		TeamId:   teamId,
		EndsData: endsData,
	}

	if err := r.Conn.Create(record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r *RecordRepository) FindById(id string) (*model.Record, error) {
	var record model.Record
	if err := r.Conn.First(&record, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *RecordRepository) FindByTeamId(teamId string) (*model.Record, error) {
	var record model.Record
	if err := r.Conn.First(&record, "team_id = ?", teamId).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *RecordRepository) Update(recordId, place string, date time.Time, endsData datatypes.JSON, isPublic bool) (*model.Record, error) {

	updateRecord := &model.Record{
		Place:    place,
		Date:     date,
		EndsData: endsData,
		IsPublic: isPublic,
	}

	if err := r.Conn.Model(&model.Record{}).Where("id = ?", recordId).Updates(updateRecord).Error; err != nil {
		return nil, err
	}
	return updateRecord, nil
}

func (r *RecordRepository) Delete(id string) error {
	if err := r.Conn.Delete(&model.Record{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
