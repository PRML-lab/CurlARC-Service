package model

import "github.com/lib/pq"

type User struct {
	Id      string         `json:"id"`
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	TeamIds pq.StringArray `json:"team_ids" gorm:"type:text[]"`
	// TeamIds string `json:"team_ids"`
}

type Users []User
