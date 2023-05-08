package infrastractures

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/joho/godotenv"
)

// DB DSN取得処理
func getDSN() string {
	// //local用相対パス
	// err := godotenv.Load(".env")
	dbURL := os.Getenv("DATABASE_URL")
	// dbURL := os.Getenv("LOCAL_MYSQL")

	u, err := url.Parse(dbURL)
	if err != nil {
		log.Fatal("Error parsing JAWSDB_URL:", err)
	}
	//フォーマット出力した文字列を返す
	user := u.User.Username()
	password, _ := u.User.Password()
	host := u.Hostname()
	port := u.Port()
	database := strings.TrimPrefix(u.Path, "/")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

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
