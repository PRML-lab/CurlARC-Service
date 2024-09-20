package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Coordinate struct {
	Index int     `json:"index"`
	R     float64 `json:"r"`
	Theta float64 `json:"theta"`
}

type Stones struct {
	FriendStones []Coordinate `json:"friend_stones"`
	EnemyStones  []Coordinate `json:"enemy_stones"`
}

type Shot struct {
	Type        string  `json:"type"`
	SuccessRate float64 `json:"success_rate"`
	Shooter     string  `json:"shooter"`
	Stones      Stones  `json:"stones"`
}

type DataPerEnd struct {
	Score int    `json:"score"`
	Shots []Shot `json:"shots"`
}

type Result string

const (
	Win  Result = "WIN"
	Loss Result = "LOSE"
	Draw Result = "DRAW"
)

//////////////////////////////////////////////////////////////////////////////////////////
// Record domain model
//////////////////////////////////////////////////////////////////////////////////////////

type Record struct {
	id            RecordId
	teamId        string
	result        Result
	enemyTeamName string
	place         string
	date          time.Time
	endsData      []DataPerEnd
	isPublic      bool
}

// RecordOption is a functional option for creating a new Record.
type RecordOption func(*Record) error

func WithEnemyTeamName(name string) RecordOption {
	return func(r *Record) error {
		return r.SetEnemyTeamName(name)
	}
}

func WithPlace(place string) RecordOption {
	return func(r *Record) error {
		return r.SetPlace(place)
	}
}

func WithDate(date time.Time) RecordOption {
	return func(r *Record) error {
		return r.SetDate(date)
	}
}

func NewRecord(teamId string, options ...RecordOption) (*Record, error) {
	recordId := NewRecordId(uuid.New().String())
	record := &Record{
		id:     *recordId,
		teamId: teamId,
	}

	for _, opt := range options {
		if err := opt(record); err != nil {
			return nil, err
		}
	}

	return record, nil
}

func NewRecordFromDB(id, teamId, enemyTeamName, place string, result Result, date time.Time, endsData []DataPerEnd, isPublic bool) *Record {
	return &Record{
		id:            *NewRecordId(id),
		teamId:        teamId,
		result:        result,
		enemyTeamName: enemyTeamName,
		place:         place,
		date:          date,
		endsData:      endsData,
		isPublic:      isPublic,
	}
}

func (r *Record) ValidateEndsData(endsData []DataPerEnd) error {
	for _, end := range endsData {
		if len(end.Shots) != 8 {
			return errors.New("each end must contains 8 shots")
		}
	}
	return nil
}

// getter

func (r *Record) GetId() *RecordId {
	return &r.id
}

func (r *Record) GetTeamId() string {
	return r.teamId
}

func (r *Record) GetResult() Result {
	return r.result
}

func (r *Record) GetEnemyTeamName() string {
	return r.enemyTeamName
}

func (r *Record) GetPlace() string {
	return r.place
}

func (r *Record) GetDate() time.Time {
	return r.date
}

func (r *Record) GetEndsData() []DataPerEnd {
	return r.endsData
}

func (r *Record) IsPublic() bool {
	return r.isPublic
}

// setter

func (r *Record) SetEnemyTeamName(name string) error {
	r.enemyTeamName = name
	return nil
}

func (r *Record) SetPlace(place string) error {
	r.place = place
	return nil
}

// SetDate sets the date of the match. Future dates are not allowed.
func (r *Record) SetDate(date time.Time) error {
	if date.After(time.Now()) {
		return errors.New("the match date cannot be in the future")
	}
	r.date = date
	return nil
}

// SetEndsData sets the ends data of the match and performs basic validation based on curling rules.
func (r *Record) SetEndsData(endsData []DataPerEnd) error {
	if err := r.ValidateEndsData(endsData); err != nil {
		return err
	}
	r.endsData = endsData
	return nil
}

func (r *Record) SetVisibility(isPublic bool) {
	r.isPublic = isPublic
}
