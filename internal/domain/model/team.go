package model

type Team struct {
	Id   string `gorm:"primaryKey"`
	Name string `gorm:"size:255 unique"`
}

type UserTeam struct {
	UserId string `gorm:"primaryKey"`
	TeamId string `gorm:"primaryKey"`
}
