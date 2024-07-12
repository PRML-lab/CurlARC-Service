package model

import (
	"time"
)

// Coordinate ストーンの位置
type Coordinate struct {
	StoneNumber int     `json:"stone_number"`
	R           float64 `json:"r"`
	Theta       float64 `json:"theta"`
}

// Coordinates カスタムタイプとして実装
type Coordinates []Coordinate

// Stones ストーンの集合
type Stones struct {
	FriendStones Coordinates `json:"friend_stones"`
	EnemyStones  Coordinates `json:"enemy_stones"`
}

// Shot ショットの情報
type Shot struct {
	Type        string  `json:"type"`
	SuccessRate float64 `json:"success_rate"`
	Shooter     string  `json:"shooter"`
	Stones      Stones  `json:"stones"`
}

// Shots カスタムタイプとして実装
type Shots []Shot

// DataPerEnd 1エンド分のデータ
type DataPerEnd struct {
	Shots Shots `json:"shots"`
	Score int   `json:"score"`
}

// DataPerEnds カスタムタイプとして実装
type DataPerEnds []DataPerEnd

// Record 1試合の記録
type Record struct {
	ID       uint        `gorm:"primaryKey"`
	EndsData DataPerEnds `json:"ends_data"`
	Place    string      `json:"place"`
	Date     time.Time   `json:"date"`
}

type Records []Record
