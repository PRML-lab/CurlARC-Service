package model

import (
	"time"
)

// ストーンの位置(ハウスの中心からの極座標で表現)
type Coordinate struct {
	StoneNumber int     `json:"stone_number"`
	R           float64 `json:"r"`
	Theta       float64 `json:"theta"`
}

type Stones struct {
	FriendStones []Coordinate `json:"friend_stones" gorm:"embedded"`
	EnemyStones  []Coordinate `json:"enemy_stones" gorm:"embedded"`
}

// ショットの情報
type Shot struct {
	Type        string  `json:"type"`
	SuccessRate float64 `json:"success_rate"`
	Shooter     string  `json:"shooter"`
	Stones      Stones  `json:"stones" gorm:"embedded"`
}

// 1エンド分のデータ
type DataPerEnd struct {
	Shots []Shot `json:"shots" gorm:"embedded"`
	Score int    `json:"score"` // エンド終了時の得点
}

// 1試合の記録
type Record struct {
	Id string `json:"id"`
	// Team     Team         `json:"team"`
	EndsData []DataPerEnd `json:"ends_data" gorm:"embedded"`
	Place    string       `json:"place"`
	Date     time.Time    `json:"date"`
}
