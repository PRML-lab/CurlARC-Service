package infra

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() *SqlHandler {

	// 環境変数から接続情報を取得
	host := os.Getenv("DATABASE_HOST")
	dbname := os.Getenv("DATABASE_NAME")
	password := os.Getenv("DATABASE_PASSWORD")
	user := os.Getenv("DATABASE_USER")
	// port := os.Getenv("DATABASE_PORT")
	// tz := os.Getenv("DATABASE_TZ")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		host, user, password, dbname)

	// データベースへの接続
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}

	conn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// マイグレーション
	err = conn.AutoMigrate(
		&User{},
		&Team{},
		&UserTeam{},
		&Record{},
	)
	if err != nil {
		panic(err.Error)
	}

	sqlHandler := &SqlHandler{
		Conn: conn,
	}
	return sqlHandler
}
