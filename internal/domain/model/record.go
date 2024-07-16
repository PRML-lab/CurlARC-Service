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
	Id       string `json:"id"`
	RecordId string `json:"record_id"`
	Score    int    `json:"score"`
	Shots    []Shot `json:"shots"`
}

type Shot struct {
	Id          string       `json:"id"`
	EndId       string       `json:"end_id"`
	Type        string       `json:"type"`
	SuccessRate float64      `json:"success_rate"`
	Shooter     string       `json:"shooter"`
	Coordinates []Coordinate `json:"coordinates"`
}

type Coordinate struct {
	Id          string  `json:"id"`
	ShotId      string  `json:"shot_id"`
	StoneNumber int     `json:"stone_number"`
	R           float64 `json:"r"`
	Theta       float64 `json:"theta"`
}
