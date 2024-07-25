package infra

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"context"
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

func (r *RecordRepository) Create(ctx context.Context, teamId, place string, date time.Time, endsData datatypes.JSON) (*model.Record, error) {

	record := &model.Record{
		Place:    place,
		Date:     date,
		TeamId:   teamId,
		EndsData: endsData,
	}

	if err := r.Conn.WithContext(ctx).Create(record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r *RecordRepository) GetByTeamId(ctx context.Context, teamId string) (*model.Record, error) {
	var record model.Record
	if err := r.Conn.WithContext(ctx).First(&record, "team_id = ?", teamId).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *RecordRepository) Update(ctx context.Context, id, teamId, place string, date time.Time, endsData datatypes.JSON) (*model.Record, error) {

	updateRecord := &model.Record{
		Place:    place,
		Date:     date,
		TeamId:   teamId,
		EndsData: endsData,
	}

	if err := r.Conn.WithContext(ctx).Model(&model.Record{}).Where("id = ?", id).Updates(updateRecord).Error; err != nil {
		return nil, err
	}
	return updateRecord, nil
}

func (r *RecordRepository) Delete(ctx context.Context, id string) error {
	if err := r.Conn.WithContext(ctx).Delete(&model.Record{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
