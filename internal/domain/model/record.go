package model

import (
	"time"

	"gorm.io/datatypes"
)

type Record struct {
	Id       string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Place    string `gorm:"size:255"`
	Date     time.Time
	TeamId   string
	Team     Team           `gorm:"foreignKey:TeamId"`
	EndsData datatypes.JSON `gorm:"type:json"`
}

type DataPerEnd struct {
	Id       string
	RecordId string
	Score    int
	Shots    []Shot
}

type Shot struct {
	Id          string
	EndId       string
	Type        string
	SuccessRate float64
	Shooter     string
	Coordinates []Coordinate
}

type Coordinate struct {
	Id          string
	ShotId      string
	StoneNumber int
	R           float64
	Theta       float64
}
