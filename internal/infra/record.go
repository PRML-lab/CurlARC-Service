package infra

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/handler/response"
	"encoding/json"
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

// define the struct for the database
type DBRecord struct {
	Id            string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TeamId        string         `gorm:"index"`
	Result        string         `gorm:"type:varchar(10)"`
	EnemyTeamName string         `gorm:"type:varchar(255)"`
	Place         string         `gorm:"type:varchar(255)"`
	Date          time.Time      `gorm:"type:timestamp"`
	EndsDataJSON  datatypes.JSON `gorm:"type:json"`
	IsPublic      bool           `gorm:"default:false"`
}

func (r *DBRecord) ToDomain() *model.Record {
	result := model.Result(r.Result)
	endsData := convertFromJSON(r.EndsDataJSON)
	return &model.Record{
		Id:            r.Id,
		TeamId:        r.TeamId,
		Result:        &result,
		EnemyTeamName: &r.EnemyTeamName,
		Place:         &r.Place,
		Date:          &r.Date,
		EndsData:      &endsData,
		IsPublic:      &r.IsPublic,
	}
}

// convertToJSONはドメインモデルのデータをJSONに変換します
func convertToJSON(data []model.DataPerEnd) datatypes.JSON {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return jsonData
}

// convertFromJSONはJSONデータをドメインモデルに変換します
func convertFromJSON(data datatypes.JSON) []model.DataPerEnd {
	var result []model.DataPerEnd
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil
	}
	return result
}

// Saveはレコードをデータベースに保存します
func (r *RecordRepository) Save(teamId, enemyTeamName, place string, result model.Result, date time.Time) (*model.Record, error) {
	dbRecord := DBRecord{
		TeamId:        teamId,
		Result:        string(result),
		EnemyTeamName: enemyTeamName,
		Place:         place,
		Date:          date,
		IsPublic:      false, // デフォルト値
	}

	if err := r.Conn.Create(&dbRecord).Error; err != nil {
		return nil, err
	}

	return dbRecord.ToDomain(), nil
}

// FindByRecordIdはIDでレコードを検索します
func (r *RecordRepository) FindByRecordId(id string) (*model.Record, error) {
	var dbRecord DBRecord
	if err := r.Conn.First(&dbRecord, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return dbRecord.ToDomain(), nil
}

// FindIndicesByTeamIdはチームIDでレコードのインデックスを検索します
func (r *RecordRepository) FindIndicesByTeamId(teamId string) (*[]response.RecordIndex, error) {
	var dbRecords []DBRecord
	if err := r.Conn.Select("id", "result", "enemy_team_name", "place", "date").Where("team_id = ?", teamId).Find(&dbRecords).Error; err != nil {
		return nil, err
	}

	var recordIndices []response.RecordIndex
	for _, dbRecord := range dbRecords {
		recordIndex := response.RecordIndex{
			Id:            dbRecord.Id,
			Result:        model.Result(dbRecord.Result),
			EnemyTeamName: dbRecord.EnemyTeamName,
			Place:         dbRecord.Place,
			Date:          dbRecord.Date,
		}
		recordIndices = append(recordIndices, recordIndex)
	}

	return &recordIndices, nil
}

// FindByTeamIdはチームIDでレコードを検索します
func (r *RecordRepository) FindByTeamId(teamId string) (*[]model.Record, error) {
	var dbRecords []DBRecord
	if err := r.Conn.Where("team_id = ?", teamId).Find(&dbRecords).Error; err != nil {
		return nil, err
	}

	var records []model.Record
	for _, dbRecord := range dbRecords {
		records = append(records, *dbRecord.ToDomain())
	}

	return &records, nil
}

// Updateはレコードを更新します
func (r *RecordRepository) Update(updatedRecord model.Record) (*model.Record, error) {
	var dbRecord DBRecord
	if err := r.Conn.First(&dbRecord, "id = ?", recordId).Error; err != nil {
		return nil, err
	}

	if updates.Place != nil {
		dbRecord.Place = *updates.Place
	}
	if updates.EnemyTeamName != nil {
		dbRecord.EnemyTeamName = *updates.EnemyTeamName
	}
	if updates.Result != nil {
		dbRecord.Result = string(*updates.Result)
	}
	if updates.Date != nil {
		dbRecord.Date = *updates.Date
	}
	if updates.EndsData != nil {
		dbRecord.EndsDataJSON = convertToJSON(*updates.EndsData)
	}
	if updates.IsPublic != nil {
		dbRecord.IsPublic = *updates.IsPublic
	}

	if err := r.Conn.Save(&dbRecord).Error; err != nil {
		return nil, err
	}

	return dbRecord.ToDomain(), nil
}

// Deleteはレコードを削除します
func (r *RecordRepository) Delete(id string) error {
	if err := r.Conn.Delete(&DBRecord{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
