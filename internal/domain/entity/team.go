package entity

import (
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

// getter

func (t *Team) GetId() TeamId {
	return t.id
}

func (t *Team) GetName() string {
	return t.name
}

// setter

func (t *Team) SetName(name string) {
	t.name = name
}
