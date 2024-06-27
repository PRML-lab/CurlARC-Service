package infra

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() *SqlHandler {
	dsn := "host=db user=app password=password dbname=app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := &SqlHandler{
		Conn: conn,
	}
	return sqlHandler
}
