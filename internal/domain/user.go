package domain

type User struct {
	Id      string
	Name    string
	Email   string
	TeamIds []string
}

type Users []User
