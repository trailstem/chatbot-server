package middleware

import (
	"github.com/gin-contrib/cors"
)

// // CORS設定処理
func SetCORS() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://chatbot-client-react.herokuapp.com", "http://localhost:3000"} // 許可するオリジンを指定
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}            // メソッド指定
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}                            // ヘッダー指定
	return config
}
