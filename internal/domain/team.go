package domain

type Team struct {
	Id      string
	Name    string
	Members Users
	Records Records
}
