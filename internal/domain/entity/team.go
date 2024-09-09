package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Team struct {
	id      TeamId
	name    string
	records []Record
	users   []User
}

func NewTeam(name string) *Team {
	teamId := NewTeamId(uuid.New().String())
	return &Team{
		id:   *teamId,
		name: name,
	}
}

func NewTeamFromDB(id string, name string) *Team {
	teamId := NewTeamId(id)
	return &Team{
		id:   *teamId,
		name: name,
	}
}

// AddRecord with business rule
func (t *Team) AddRecord(record Record) error {
	// Example rule: Limit the number of records to 100
	if len(t.records) >= 100 {
		return errors.New("cannot add more than 100 records")
	}
	t.records = append(t.records, record)
	return nil
}

// RemoveRecord removes a record from the team's records.
func (t *Team) RemoveRecord(recordId RecordId) error {
	for i, record := range t.records {
		if record.GetId().Equals(&recordId) {
			t.records = append(t.records[:i], t.records[i+1:]...)
			return nil
		}
	}
	return errors.New("record not found")
}

// AddUser with business rule
func (t *Team) AddUser(user User) error {
	for _, u := range t.users {
		if u.GetId().Equals(user.GetId()) {
			return errors.New("user already in team")
		}
	}
	t.users = append(t.users, user)
	return nil
}

// RemoveUser removes a user from the team.
func (t *Team) RemoveUser(userId UserId) error {
	for i, user := range t.users {
		if user.GetId().Equals(&userId) {
			t.users = append(t.users[:i], t.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

// getter

func (t *Team) GetId() *TeamId {
	return &t.id
}

func (t *Team) GetName() string {
	return t.name
}

func (t *Team) GetRecords() []Record {
	return t.records
}

func (t *Team) GetUsers() []User {
	return t.users
}

// setter

func (t *Team) SetName(name string) {
	t.name = name
}
