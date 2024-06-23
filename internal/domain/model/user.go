package model

type User struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	TeamIds []string `json:"team_ids"`
}

type Users []User
