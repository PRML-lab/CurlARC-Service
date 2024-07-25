package model

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Teams []Team `gorm:"many2many:user_teams;constraint:OnDelete:CASCADE;"`
}
