package infrastractures

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB DSN取得処理
func getDSN() string {

	dbURL := os.Getenv("DATABASE_URL")

	// HerokuのDB接続情報を取得できた場合
	if dbURL != "" {
		// u, err := url.Parse(dbURL)
		// if err != nil {
		// 	log.Fatal("Error parsing JAWSDB_URL:", err)
		// }
		// //フォーマット出力した文字列を返す
		// user := u.User.Username()
		// password, _ := u.User.Password()
		// host := u.Hostname()
		// port := u.Port()
		// database := strings.TrimPrefix(u.Path, "/")
		// return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
		return dbURL

	} else {

		// 現在の実行ファイルの絶対パスを取得
		exe, err := os.Executable()
		if err != nil {
			panic(err)
		}
		// 現在の実行ファイルのディレクトリを取得
		/*
			local     ：/Users/.../.../chatbot/chatbot-server/cmd
			コンテナ内：/app/cmd
		*/
		exeDir := filepath.Dir(exe)
		// 1つ上のディレクトリを取得
		parentDir := filepath.Dir(exeDir)
		// .env ファイルの絶対パスを生成
		envPath := filepath.Join(parentDir, ".env")
		// .env ファイルを読み込む
		err = godotenv.Load(envPath)
		if err != nil {
			log.Fatal("環境変数又は .envファイルが読み込めませんでした")
		}
		dbURL = os.Getenv("DATABASE_URL")
		return dbURL
	}
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
