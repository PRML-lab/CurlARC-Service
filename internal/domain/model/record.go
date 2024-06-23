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
	FriendStones []Coordinate
	EnemyStones  []Coordinate
}

// ショットの情報
type Shot struct {
	Type        string
	SuccessRate float64
	Shooter     User
	Stones      Stones
}

// 1エンド分のデータ
type DataPerEnd struct {
	Shots []Shot
	Score int // エンド終了時の得点
}

// 1試合の記録
type Record struct {
	Id       string
	Team     Team
	EndsData []DataPerEnd
	Place    string
	Date     time.Time
}

type Records []Record
