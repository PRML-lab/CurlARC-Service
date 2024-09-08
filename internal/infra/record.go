package infra

import (
	entity "CurlARC/internal/domain/entity/record"
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
	id            string         `gorm:"type:uuid;primaryKey"`
	teamId        string         `gorm:"index"`
	result        string         `gorm:"type:varchar(10)"`
	enemyTeamName string         `gorm:"type:varchar(255)"`
	place         string         `gorm:"type:varchar(255)"`
	date          time.Time      `gorm:"type:timestamp"`
	endsDataJSON  datatypes.JSON `gorm:"type:json"`
	isPublic      bool           `gorm:"type:boolean"`
}

func (r *DBRecord) ToDomain() *entity.Record {
	result := entity.Result(r.result)           // convert string to Result
	endsData := convertFromJSON(r.endsDataJSON) // convert JSON to []DataPerEnd

	record := entity.NewRecordFromDB(r.id, r.teamId, r.enemyTeamName, r.place, result, r.date, endsData, r.isPublic) // create a new Record

	return record
}

// convert DataPerEnd to JSON
func convertToJSON(data []entity.DataPerEnd) datatypes.JSON {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return jsonData
}

// convert JSON to DataPerEnd
func convertFromJSON(data datatypes.JSON) []entity.DataPerEnd {
	var result []entity.DataPerEnd
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil
	}
	return result
}

func (r *RecordRepository) Save(record entity.Record) (*entity.Record, error) {
	dbRecord := DBRecord{
		id:            record.GetId().Value(),
		teamId:        record.GetTeamId(),
		result:        string(record.GetResult()),
		enemyTeamName: record.GetEnemyTeamName(),
		place:         record.GetPlace(),
		date:          record.GetDate(),
		endsDataJSON:  convertToJSON(record.GetEndsData()),
		isPublic:      record.IsPublic(),
	}

	if err := r.Conn.Create(&dbRecord).Error; err != nil {
		return nil, err
	}

	return dbRecord.ToDomain(), nil
}

func (r *RecordRepository) FindByRecordId(recordId string) (*entity.Record, error) {
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
			Result:        entity.Result(dbRecord.Result),
			EnemyTeamName: dbRecord.EnemyTeamName,
			Place:         dbRecord.Place,
			Date:          dbRecord.Date,
		}
		recordIndices = append(recordIndices, recordIndex)
	}

	return &recordIndices, nil
}

// FindByTeamIdはチームIDでレコードを検索します
func (r *RecordRepository) FindByTeamId(teamId string) (*[]entity.Record, error) {
	var dbRecords []DBRecord
	if err := r.Conn.Where("team_id = ?", teamId).Find(&dbRecords).Error; err != nil {
		return nil, err
	}

	var records []entity.Record
	for _, dbRecord := range dbRecords {
		records = append(records, *dbRecord.ToDomain())
	}

	return &records, nil
}

// Updateはレコードを更新します
func (r *RecordRepository) Update(updatedRecord entity.Record) (*entity.Record, error) {
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
