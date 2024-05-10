package domain

import (
	"time"
)

// ストーンの位置(ハウスの中心からの極座標で表現)
type StoneCoordinate struct {
	StoneNumber int
	R           float64
	Theta       float64
}

// ショットの情報
type Shot struct {
	Type        string
	SuccessRate float64
	Shooter     User
}

// 1エンド分のデータ
type DataPerEnd struct {
	StonesCoordinate []StoneCoordinate // 全ストーンの位置
	Shot             Shot
	Score            int // エンド終了時の得点
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
