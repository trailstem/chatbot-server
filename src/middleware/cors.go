package middleware

import (
	"github.com/gin-contrib/cors"
)

// // CORS設定処理
func SetCORS() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}                                                        // フロントエンド オリジンを指定
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"} // メソッド指定
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}                 // ヘッダー指定
	return config
}
