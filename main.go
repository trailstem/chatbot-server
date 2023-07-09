package main

import (
	"github.com/gin-gonic/gin"
	"github.com/trailstem/chatbot-server/infrastractures"
)

func main() {
	// 本番環境設定
	gin.SetMode(gin.ReleaseMode)
	// ルーティング設定
	r := infrastractures.SetupRouter()
	// 環境変数PORTが設定されていない場合は8080を使用

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
