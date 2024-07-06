package model

import "github.com/lib/pq"

type User struct {
	Id      string         `json:"id"`
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	TeamIds pq.StringArray `gorm:"type:jsonb"` // 配列だとGORMの対応がめんどくさいので一旦teamIdは１つだけにする
	// TeamIds string `json:"team_ids"`
}

type Users []User
