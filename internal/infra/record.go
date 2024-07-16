package infra

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"context"
	"encoding/json"
	"time"
)

type RecordRepository struct {
	SqlHandler
}

func NewRecordRepository(sqlHandler SqlHandler) repository.RecordRepository {
	recordRepository := RecordRepository{SqlHandler: sqlHandler}
	return &recordRepository
}

func (r *RecordRepository) Create(ctx context.Context, teamId, place string, date time.Time, endsData []model.DataPerEnd) (*model.Record, error) {
	endsDataJSON, err := json.Marshal(endsData)
	if err != nil {
		return nil, err
	}

	record := &model.Record{
		Place:    place,
		Date:     date,
		TeamId:   teamId,
		EndsData: endsDataJSON,
	}

	if err := r.Conn.WithContext(ctx).Create(record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r *RecordRepository) GetById(ctx context.Context, id string) (*model.Record, error) {
	var record model.Record
	if err := r.Conn.WithContext(ctx).First(&record, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *RecordRepository) Update(ctx context.Context, id, teamId, place string, date time.Time, endsData []model.DataPerEnd) error {
	endsDataJSON, err := json.Marshal(endsData)
	if err != nil {
		return err
	}

	updateData := map[string]interface{}{
		"team_id":   teamId,
		"place":     place,
		"date":      date,
		"ends_data": endsDataJSON,
	}

	if err := r.Conn.WithContext(ctx).Model(&model.Record{}).Where("id = ?", id).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}

func (r *RecordRepository) Delete(ctx context.Context, id string) error {
	if err := r.Conn.WithContext(ctx).Delete(&model.Record{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
