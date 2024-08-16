package model

type Team struct {
	Id    string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name  string `gorm:"size:255 unique" json:"name"`
}


