package model

import (
	"time"

	"gorm.io/datatypes"
)

type Result string

const (
	Win  Result = "WIN"
	Loss Result = "LOSE"
	Draw Result = "DRAW"
)

type Record struct {
	Id            string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TeamId        string         `json:"team_id"`
	Team          Team           `gorm:"foreignKey:TeamId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team"`
	Result        Result         `json:"result"`
	EnemyTeamName string         `json:"enemy_team_name"`
	Place         string         `gorm:"size:255" json:"place"`
	Date          time.Time      `json:"date"`
	EndsData      datatypes.JSON `gorm:"type:json" json:"ends_data"`
	IsPublic      bool           `gorm:"default:false" json:"is_public"`
}

type RecordUpdate struct {
	Place         *string         `json:"place,omitempty"`
	EnemyTeamName *string         `json:"enemy_team_name,omitempty"`
	Result        *Result         `json:"result,omitempty"`
	Date          *time.Time      `json:"date,omitempty"`
	EndsData      *datatypes.JSON `json:"ends_data,omitempty"`
	IsPublic      *bool           `json:"is_public,omitempty"`
}

type DataPerEnd struct {
	Score int    `json:"score"`
	Shots []Shot `json:"shots"`
}

type Shot struct {
	Type        string  `json:"type"`
	SuccessRate float64 `json:"success_rate"`
	Shooter     string  `json:"shooter"`
	Stones      Stones  `json:"stones"`
}

type Stones struct {
	FriendStones []Coordinate `json:"friend_stones"`
	EnemyStones  []Coordinate `json:"enemy_stones"`
}

type Coordinate struct {
	Index int     `json:"index"`
	R     float64 `json:"r"`
	Theta float64 `json:"theta"`
}
