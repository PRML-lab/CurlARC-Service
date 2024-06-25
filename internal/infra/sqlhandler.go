package infra

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	dsn := "host=db user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := &SqlHandler{
		Conn: conn,
	}
	return sqlHandler
}
