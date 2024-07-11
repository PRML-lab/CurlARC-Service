package infra

import (
	"CurlARC/internal/domain/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() *SqlHandler {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// 環境変数から接続情報を取得
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")
	tz := os.Getenv("DATABASE_TZ")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s",
		host, user, password, dbname, port, tz)

	// データベースへの接続
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}

	// マイグレーション
	err = conn.AutoMigrate(&model.User{}, &model.Team{}, &model.Record{})
	if err != nil {
		panic(err.Error)
	}

	sqlHandler := &SqlHandler{
		Conn: conn,
	}
	return sqlHandler
}
