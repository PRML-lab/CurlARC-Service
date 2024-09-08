package infra

import (
	"CurlARC/internal/domain/entity"
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
type Record struct {
	Id            string         `gorm:"type:uuid;primaryKey"`
	TeamId        string         `gorm:"foreignKey:TeamId"`
	Result        string         `gorm:"type:varchar(10)"`
	EnemyTeamName string         `gorm:"type:varchar(255)"`
	Place         string         `gorm:"type:varchar(255)"`
	Date          time.Time      `gorm:"type:timestamp"`
	EndsDataJSON  datatypes.JSON `gorm:"type:json"`
	IsPublic      bool           `gorm:"type:boolean"`
}

func (r *Record) FromDomain(record *entity.Record) {
	r.Id = record.GetId().Value()
	r.TeamId = record.GetTeamId()
	r.Result = string(record.GetResult())
	r.EnemyTeamName = record.GetEnemyTeamName()
	r.Place = record.GetPlace()
	r.Date = record.GetDate()
	r.EndsDataJSON = convertToJSON(record.GetEndsData())
	r.IsPublic = record.IsPublic()
}

func (r *Record) ToDomain() *entity.Record {
	result := entity.Result(r.Result)           // convert string to Result
	endsData := convertFromJSON(r.EndsDataJSON) // convert JSON to []DataPerEnd

	record := entity.NewRecordFromDB(r.Id, r.TeamId, r.EnemyTeamName, r.Place, result, r.Date, endsData, r.IsPublic) // create a new Record

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

////////////////////////////////////////////////////////////////
// Record Repository Implementation
////////////////////////////////////////////////////////////////

func (r *RecordRepository) Save(record entity.Record) (*entity.Record, error) {
	var dbRecord Record
	dbRecord.FromDomain(&record)

	if err := r.Conn.Create(&dbRecord).Error; err != nil {
		return nil, err
	}

	return dbRecord.ToDomain(), nil
}

func (r *RecordRepository) FindByRecordId(recordId string) (*entity.Record, error) {
	var dbRecord Record
	if err := r.Conn.First(&dbRecord, "id = ?", recordId).Error; err != nil {
		return nil, err
	}
	return dbRecord.ToDomain(), nil
}

func (r *RecordRepository) FindIndicesByTeamId(teamId string) (*[]response.RecordIndex, error) {
	var dbRecords []Record
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

func (r *RecordRepository) FindByTeamId(teamId string) (*[]entity.Record, error) {
	var dbRecords []Record
	if err := r.Conn.Where("team_id = ?", teamId).Find(&dbRecords).Error; err != nil {
		return nil, err
	}

	var records []entity.Record
	for _, dbRecord := range dbRecords {
		records = append(records, *dbRecord.ToDomain())
	}

	return &records, nil
}

func (r *RecordRepository) Update(record entity.Record) (*entity.Record, error) {
	var dbRecord Record
	if err := r.Conn.First(&dbRecord, "id = ?", record.GetId().Value()).Error; err != nil {
		return nil, err
	}

	dbRecord.FromDomain(&record)

	if err := r.Conn.Save(&dbRecord).Error; err != nil {
		return nil, err
	}

	return dbRecord.ToDomain(), nil
}

func (r *RecordRepository) Delete(id string) error {
	if err := r.Conn.Delete(&Record{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
