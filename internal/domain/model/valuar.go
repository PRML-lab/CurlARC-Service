package model

import (
	"database/sql/driver"
	"encoding/json"
)

func (c Coordinates) Value() (driver.Value, error) {
	return json.Marshal(c)
}
func (s Stones) Value() (driver.Value, error) {
	return json.Marshal(s)
}
func (s Shot) Value() (driver.Value, error) {
	return json.Marshal(s)
}
func (s Shots) Value() (driver.Value, error) {
	return json.Marshal(s)
}
func (d DataPerEnd) Value() (driver.Value, error) {
	return json.Marshal(d)
}
func (d DataPerEnds) Value() (driver.Value, error) {
	return json.Marshal(d)
}
func (r Records) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (t Team) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (u Users) Value() (driver.Value, error) {
	return json.Marshal(u)
}
