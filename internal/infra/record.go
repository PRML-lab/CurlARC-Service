package infra

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"context"
)

type RecordRepository struct {
	SqlHandler
}

func NewRecordRepository(sqlHandler SqlHandler) repository.RecordRepository {
	recordRepository := RecordRepository{SqlHandler: sqlHandler}
	return &recordRepository
}

func (recordRepo *RecordRepository) Create(ctx context.Context, record *model.Record) error {
	result := recordRepo.SqlHandler.Conn.Create(record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (recordRepo *RecordRepository) GetById(ctx context.Context, Id string) (*model.Record, error) {
	record := new(model.Record)
	result := recordRepo.SqlHandler.Conn.Where("id = ?", Id).First(record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (recordRepo *RecordRepository) GetByTeamId(ctx context.Context, teamId string) ([]*model.Record, error) {
	var records []*model.Record
	result := recordRepo.SqlHandler.Conn.Where("team_id = ?", teamId).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}

func (recordRepo *RecordRepository) Update(ctx context.Context, record *model.Record) error {
	result := recordRepo.SqlHandler.Conn.Save(record)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (recordRepo *RecordRepository) Delete(ctx context.Context, Id string) error {
	result := recordRepo.SqlHandler.Conn.Delete(&model.Record{}, Id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type EndRepository struct {
	SqlHandler
}

func NewEndRepository(sqlHandler SqlHandler) repository.EndRepository {
	endRepository := EndRepository{SqlHandler: sqlHandler}
	return &endRepository
}

func (endRepo *EndRepository) Create(ctx context.Context, end *model.End) error {
	result := endRepo.SqlHandler.Conn.Create(end)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (endRepo *EndRepository) GetById(ctx context.Context, Id string) (*model.End, error) {
	end := new(model.End)
	result := endRepo.SqlHandler.Conn.Where("id = ?", Id).First(end)
	if result.Error != nil {
		return nil, result.Error
	}
	return end, nil
}

func (endRepo *EndRepository) Update(ctx context.Context, end *model.End) error {
	result := endRepo.SqlHandler.Conn.Save(end)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (endRepo *EndRepository) Delete(ctx context.Context, Id string) error {
	result := endRepo.SqlHandler.Conn.Delete(&model.End{}, Id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type ShotRepository struct {
	SqlHandler
}

func NewShotRepository(sqlHandler SqlHandler) repository.ShotRepository {
	shotRepository := ShotRepository{SqlHandler: sqlHandler}
	return &shotRepository
}

func (shotRepo *ShotRepository) Create(ctx context.Context, shot *model.Shot) error {
	result := shotRepo.SqlHandler.Conn.Create(shot)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (shotRepo *ShotRepository) GetById(ctx context.Context, Id string) (*model.Shot, error) {

	shot := new(model.Shot)
	result := shotRepo.SqlHandler.Conn.Where("id = ?", Id).First(shot)
	if result.Error != nil {
		return nil, result.Error
	}
	return shot, nil
}

func (shotRepo *ShotRepository) Update(ctx context.Context, shot *model.Shot) error {
	result := shotRepo.SqlHandler.Conn.Save(shot)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (shotRepo *ShotRepository) Delete(ctx context.Context, Id string) error {
	result := shotRepo.SqlHandler.Conn.Delete(&model.Shot{}, Id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
