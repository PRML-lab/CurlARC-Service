package model

import (
	"encoding/json"
	"fmt"
)

func (c *Coordinates) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Coordinates: %v", value)
	}
	return json.Unmarshal(bytes, c)
}
func (s *Stones) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Stones: %v", value)
	}
	return json.Unmarshal(bytes, s)
}
func (s *Shot) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Shot: %v", value)
	}
	return json.Unmarshal(bytes, s)
}
func (s *Shots) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Shots: %v", value)
	}
	return json.Unmarshal(bytes, s)
}
func (d *DataPerEnd) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan DataPerEnd: %v", value)
	}
	return json.Unmarshal(bytes, d)
}
func (d *DataPerEnds) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan DataPerEnds: %v", value)
	}
	return json.Unmarshal(bytes, d)
}
func (r *Records) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Records: %v", value)
	}
	return json.Unmarshal(bytes, r)
}

func (t *Team) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Team: %v", value)
	}
	return json.Unmarshal(bytes, t)
}

func (u *Users) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Users: %v", value)
	}
	return json.Unmarshal(bytes, u)
}
