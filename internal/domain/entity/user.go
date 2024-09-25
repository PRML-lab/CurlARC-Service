package entity

import "github.com/google/uuid"

type User struct {
	id    UserId
	name  string
	email string
	teams []Team
}

func NewUser(name string, email string) *User {
	userId := NewUserId(uuid.New().String())
	return &User{
		id:    *userId,
		name:  name,
		email: email,
	}
}

func NewUserFromDB(id string, name string, email string) *User {
	userId := NewUserId(id)
	return &User{
		id:    *userId,
		name:  name,
		email: email,
	}
}

// getter

func (u *User) GetId() *UserId {
	return &u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetTeams() []Team {
	return u.teams
}

// setter

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) AddTeam(team Team) {
	u.teams = append(u.teams, team)
}

func (u *User) RemoveTeam(team Team) {
	for i, t := range u.teams {
		if t.GetId().Equals(team.GetId()) {
			u.teams = append(u.teams[:i], u.teams[i+1:]...)
		}
	}
}
