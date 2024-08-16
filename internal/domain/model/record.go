package model

import (
	"time"

	"gorm.io/datatypes"
)

type Record struct {
	Id       string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TeamId   string `json:"team_id"`
	Team     Team   `gorm:"foreignKey:TeamId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Place    string `gorm:"size:255" json:"place"`
	Date     time.Time `json:"date"`
	EndsData datatypes.JSON `gorm:"type:json" json:"ends_data"`
	IsPublic bool           `gorm:"default:false" json:"is_public"`
}

type DataPerEnd struct {
	Index int
	Score int
	Shots []Shot
}

type Shot struct {
	Index       int
	Type        string
	SuccessRate float64
	Shooter     string
	Stones      Stones
}

type Stones struct {
	FriendStones []Coordinate
	EnemyStones  []Coordinate
}

type Coordinate struct {
	Index int
	R     float64
	Theta float64
}
