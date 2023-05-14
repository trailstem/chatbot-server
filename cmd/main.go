package main

import (
	"os"

	"github.com/trailstem/chatbot-server/infrastractures"
)

func main() {

	// ルーティング設定
	r := infrastractures.SetupRouter()
	// 環境変数PORTが設定されていない場合は8080を使用
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
