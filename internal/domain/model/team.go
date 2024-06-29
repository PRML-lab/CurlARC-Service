package model

type Team struct {
	Id      string
	Name    string
	Members []User   `gorm:"many2many:team_members;"`
	Records []Record `gorm:"foreignKey:Id"`
}
