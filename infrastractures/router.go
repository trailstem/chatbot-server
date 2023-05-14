package infrastractures

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/trailstem/chatbot-server/adapters/controllers"
	"github.com/trailstem/chatbot-server/adapters/gateways"
	"github.com/trailstem/chatbot-server/middleware"
	"github.com/trailstem/chatbot-server/usecases"

	"github.com/gin-contrib/cors"
)

// SetupRouter ルーティング設定
func SetupRouter() *gin.Engine {

	// デフォルトのミドルウェアを使用
	r := gin.New()
	// CORS対応
	r.Use(cors.New(middleware.SetCORS()))

	// DB接続（MySQL）
	conn, err := ConnectDB()
	if err != nil {
		// DB接続に失敗した場合はエラーを出力して終了
		fmt.Println(err)
		panic(err)
	}

	//repositoryインスタンス生成
	speakBotRepo := gateways.NewHistoryListRepository(conn)
	// usecaseのインスタンス生成
	speakBotUsecase := usecases.NewHistoryListUsecase(speakBotRepo)
	// controllerのインスタンス生成
	speakBotController := controllers.NewHistoryListController(speakBotUsecase)

	// ルーティング
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	// botとのやりとり
	r.POST("/chat", speakBotController.CreateChatData)
	// 過去10件のやりとりを取得
	r.GET("/history/list", speakBotController.FindAChatDataList)
	return r
}
