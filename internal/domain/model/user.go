package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// TeamIds []string `json:"team_ids"` // 配列だとGORMの対応がめんどくさいので一旦teamIdは１つだけにする
	TeamIds string `json:"team_ids"`
}

type Users []User

func (u Users) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u *Users) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Users: %v", value)
	}
	return json.Unmarshal(bytes, u)
}
