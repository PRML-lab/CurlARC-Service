package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Team struct {
	Id      string
	Name    string
	Members Users
	Records Records
}

func (t Team) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *Team) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Team: %v", value)
	}
	return json.Unmarshal(bytes, t)
}
