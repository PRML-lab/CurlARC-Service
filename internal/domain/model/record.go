package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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

func (c Coordinates) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Coordinates) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Coordinates: %v", value)
	}
	return json.Unmarshal(bytes, c)
}

// Stones ストーンの集合
type Stones struct {
	FriendStones Coordinates `json:"friend_stones"`
	EnemyStones  Coordinates `json:"enemy_stones"`
}

func (s Stones) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Stones) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Stones: %v", value)
	}
	return json.Unmarshal(bytes, s)
}

// Shot ショットの情報
type Shot struct {
	Type        string  `json:"type"`
	SuccessRate float64 `json:"success_rate"`
	Shooter     string  `json:"shooter"`
	Stones      Stones  `json:"stones"`
}

func (s Shot) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Shot) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Shot: %v", value)
	}
	return json.Unmarshal(bytes, s)
}

// Shots カスタムタイプとして実装
type Shots []Shot

func (s Shots) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Shots) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Shots: %v", value)
	}
	return json.Unmarshal(bytes, s)
}

// DataPerEnd 1エンド分のデータ
type DataPerEnd struct {
	Shots Shots `json:"shots"`
	Score int   `json:"score"`
}

func (d DataPerEnd) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *DataPerEnd) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan DataPerEnd: %v", value)
	}
	return json.Unmarshal(bytes, d)
}

// DataPerEnds カスタムタイプとして実装
type DataPerEnds []DataPerEnd

func (d DataPerEnds) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *DataPerEnds) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan DataPerEnds: %v", value)
	}
	return json.Unmarshal(bytes, d)
}

// Record 1試合の記録
type Record struct {
	ID       uint        `gorm:"primaryKey"`
	EndsData DataPerEnds `json:"ends_data"`
	Place    string      `json:"place"`
	Date     time.Time   `json:"date"`
}

type Records []Record

func (r Records) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Records) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Records: %v", value)
	}
	return json.Unmarshal(bytes, r)
}
