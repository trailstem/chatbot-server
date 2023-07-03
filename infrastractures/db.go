package infrastractures

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB DSN取得処理
func getDSN() (string, string) {
	var dsn string
	var dbType string
	// 本番環境の場合
	if os.Getenv("APP_ENV") == "PROD" {

		err := godotenv.Load()
		if err != nil {
			log.Fatal("サーバ設定ファイル又は .envファイルが読み込めませんでした")
		}

		dbName := os.Getenv("DB_NAME")
		dbUser := os.Getenv("DB_USER")
		dbHost := os.Getenv("DB_HOST")
		dbPortStr := os.Getenv("DB_PORT")
		dbPort, err := strconv.Atoi(dbPortStr)
		dbEndpoint := fmt.Sprintf("%s:%d", dbHost, dbPort)
		region := os.Getenv("DB_REGION")

		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic("configuration error: " + err.Error())
		}

		dsn = dbUser + ":" + "" + "@tcp(" + dbHost + ":" + dbPortStr + ")" + "/" + dbName

		authenticationToken, err := auth.BuildAuthToken(
			context.TODO(), dbEndpoint, region, dbUser, cfg.Credentials)
		if err != nil {
			panic("failed to create authentication token: " + err.Error())
		}

		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true&allowCleartextPasswords=true",
			dbUser, authenticationToken, dbEndpoint, dbName,
		)
		dbType = os.Getenv("DATABASE_TYPE")

	} else {

		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
		if err != nil {
			log.Fatal("開発環境：サーバ設定ファイル又は .envファイルが読み込めませんでした")
		}
		// 開発環境だった場合
		dsn = os.Getenv("DATABASE_URL")
		dbType = os.Getenv("DATABASE_TYPE")
	}
	return dsn, dbType
}

// データベース接続処理
func ConnectDB() (*sql.DB, error) {
	dsn, dbType := getDSN()
	//mysqlへの接続処理
	db, err := sql.Open(dbType, dsn)

	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
