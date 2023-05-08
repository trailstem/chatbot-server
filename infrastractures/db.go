package infrastractures

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/joho/godotenv"
)

// DB DSN取得処理
func getDSN() string {
	// //local用相対パス
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	dbURL := os.Getenv("DATABASE_URL")
	// dbURL := os.Getenv("LOCAL_MYSQL")

	parsedURL, err := url.Parse(dbURL)
	if err != nil {
		log.Fatal("Error parsing JAWSDB_URL:", err)
	}
	//フォーマット出力した文字列を返す
	return fmt.Sprintf("%v", parsedURL)
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
