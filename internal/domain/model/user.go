package model

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Teams []Team `gorm:"many2many:user_teams;constraint:OnDelete:CASCADE;"`
}

type UserTeam struct {
	UserId string `gorm:"primaryKey;constraint:OnDelete:CASCADE;" json:"user_id"`
	TeamId string `gorm:"type:uuid;primaryKey;constraint:OnDelete:CASCADE;" json:"team_id"`
	State  string `gorm:"size:255" json:"state"` // "INVITED" or "MEMBER"
}