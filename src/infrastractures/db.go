package infrastractures

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

// DB DSN取得処理
func getDSN() string {
	// //local用相対パス
	// err := godotenv.Load(".env")
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")

	//フォーマット出力した文字列を返す
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
}

// データベース接続処理
func ConnectDB() (*sql.DB, error) {
	dsn := getDSN()
	//mysqlへの接続処理
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
