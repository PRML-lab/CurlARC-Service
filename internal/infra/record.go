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

func (r *DBRecord) FromDomain(record *entity.Record) {
	r.id = record.GetId().Value()
	r.teamId = record.GetTeamId()
	r.result = string(record.GetResult())
	r.enemyTeamName = record.GetEnemyTeamName()
	r.place = record.GetPlace()
	r.date = record.GetDate()
	r.endsDataJSON = convertToJSON(record.GetEndsData())
	r.isPublic = record.IsPublic()
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

func (r *RecordRepository) Save(record entity.Record) error {
	var dbRecord DBRecord
	dbRecord.FromDomain(&record)

	if err := r.Conn.Create(&dbRecord).Error; err != nil {
		return err
	}

	return nil
}

func (r *RecordRepository) FindByRecordId(recordId string) (*entity.Record, error) {
	var dbRecord DBRecord
	if err := r.Conn.First(&dbRecord, "id = ?", recordId).Error; err != nil {
		return nil, err
	}
	return dbRecord.ToDomain(), nil
}

func (r *RecordRepository) FindIndicesByTeamId(teamId string) (*[]response.RecordIndex, error) {
	var dbRecords []DBRecord
	if err := r.Conn.Select("id", "result", "enemy_team_name", "place", "date").Where("team_id = ?", teamId).Find(&dbRecords).Error; err != nil {
		return nil, err
	}

	var recordIndices []response.RecordIndex
	for _, dbRecord := range dbRecords {
		recordIndex := response.RecordIndex{
			Id:            dbRecord.id,
			Result:        entity.Result(dbRecord.result),
			EnemyTeamName: dbRecord.enemyTeamName,
			Place:         dbRecord.place,
			Date:          dbRecord.date,
		}
		recordIndices = append(recordIndices, recordIndex)
	}

	return &recordIndices, nil
}

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

func (r *RecordRepository) Update(record entity.Record) error {
	var dbRecord DBRecord
	if err := r.Conn.First(&dbRecord, "id = ?", record.GetId().Value()).Error; err != nil {
		return err
	}

	dbRecord.FromDomain(&record)

	if err := r.Conn.Save(&dbRecord).Error; err != nil {
		return err
	}

	return nil
}

func (r *RecordRepository) Delete(id string) error {
	if err := r.Conn.Delete(&DBRecord{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
