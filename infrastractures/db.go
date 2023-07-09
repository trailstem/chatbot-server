package infrastractures

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB DSN取得処理
func getDSN() (string, string) {

	var dsn string
	var dbType string
	// 本番環境の場合
	if os.Getenv("APP_ENV") == "PROD" {

		// 環境変数からRDSのDB情報を取得
		dbName := os.Getenv("DB_NAME")
		dbUser := os.Getenv("DB_USER")
		dbHost := os.Getenv("DB_HOST")
		//// IAMデータベース認証用
		// dbPort := os.Getenv("DB_PORT")
		// dbEndpoint := fmt.Sprintf("%s:%s", dbHost, dbPort)
		// region := os.Getenv("DB_REGION")
		dbType = os.Getenv("DATABASE_TYPE")

		user := dbUser
		password := os.Getenv("DB_PASSWORD")
		hostname := dbHost
		dbname := dbName
		connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, hostname, dbname)

		//// IAMデータベース認証用
		// RDSの認証トークンを取得
		// cfg, err := config.LoadDefaultConfig(context.TODO())
		// if err != nil {
		// 	panic("configuration error: " + err.Error())
		// }
		// // 認証トークンを取得
		// authenticationToken, err := auth.BuildAuthToken(
		// 	context.TODO(), dbEndpoint, region, dbUser, cfg.Credentials)
		// if err != nil {
		// 	panic("failed to create authentication token: " + err.Error())
		// }
		// // DSNを作成
		// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true&allowCleartextPasswords=true",
		// 	dbUser, authenticationToken, dbEndpoint, dbName,
		// )
		// return dsn, dbType

		return connStr, dbType

	} else {

		// 開発環境の場合
		err := godotenv.Load()
		if err != nil {
			panic(err)
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
	fmt.Println(dbType)
	fmt.Println(dsn)
	//mysqlへの接続処理
	db, err := sql.Open(dbType, dsn)
	// 接続できなかった場合
	if err != nil {
		panic("driverやAWS 認証部分で不具合がありパニック: " + err.Error())
	}
	// 接続確認
	err = db.Ping()
	if err != nil {
		panic("接続がうまくいきません。: " + err.Error())
	}
	fmt.Println("DB接続成功")
	return db, err
}
