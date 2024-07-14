package model

import (
	"time"
)

type Record struct {
	Id     string `gorm:"primaryKey"`
	Place  string `gorm:"size:255"`
	Date   time.Time
	TeamId string
	Team   Team `gorm:"foreignKey:TeamId"`
}

type End struct {
	Id       string `gorm:"primaryKey"`
	RecordId string
	Record   Record `gorm:"foreignKey:RecordId"`
	Score    int
}

type Shot struct {
	Id          string `gorm:"primaryKey"`
	EndId       string
	End         End    `gorm:"foreignKey:EndId"`
	Type        string `gorm:"size:255"`
	SuccessRate float64
	Shooter     string `gorm:"size:255"`
}

type Coordinate struct {
	Id          string `gorm:"primaryKey"`
	ShotId      string
	Shot        Shot `gorm:"foreignKey:ShotId"`
	StoneNumber int
	R           float64
	Theta       float64
}
