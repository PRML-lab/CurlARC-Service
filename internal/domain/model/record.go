package model

import (
	"errors"
	"time"
)

type Result string

const (
	Win  Result = "WIN"
	Loss Result = "LOSE"
	Draw Result = "DRAW"
)

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

//////////////////////////////////////////////////////////////////////////////////////////
// Record
//////////////////////////////////////////////////////////////////////////////////////////

type Record struct {
	Id            string        `json:"id"`
	TeamId        string        `json:"team_id"`
	Team          Team          `json:"team"`
	Result        *Result       `json:"result"`
	EnemyTeamName *string       `json:"enemy_team_name"`
	Place         *string       `json:"place"`
	Date          *time.Time    `json:"date"`
	EndsData      *[]DataPerEnd `json:"ends_data"`
	IsPublic      *bool         `json:"is_public"`
}

func (r *Record) ValidateEndsData(endsData []DataPerEnd) error {
	for _, end := range endsData {
		if len(end.Shots) != 8 {
			return errors.New("each end must contain 8 shots")
		}
	}
	return nil
}

// SetDate sets the date of the match. Future dates are not allowed.
func (r *Record) SetDate(date time.Time) error {
	if date.After(time.Now()) {
		return errors.New("the match date cannot be in the future")
	}
	r.Date = &date
	return nil
}

// SetEndsData sets the ends data of the match and performs basic validation based on curling rules.
func (r *Record) SetEndsData(endsData []DataPerEnd) error {
	if err := r.ValidateEndsData(endsData); err != nil {
		return err
	}
	r.EndsData = &endsData
	return nil
}

// AppendEndsData appends ends data to the match and performs basic validation based on curling rules.
func (r *Record) AppendEndsData(endsData []DataPerEnd) error {
	if err := r.ValidateEndsData(endsData); err != nil {
		return err
	}
	*r.EndsData = append(*r.EndsData, endsData...)
	return nil
}
