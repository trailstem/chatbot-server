package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/trailstem/chatbot-server/common"
	"github.com/trailstem/chatbot-server/domain"
	"github.com/trailstem/chatbot-server/usecases"
)

type historyListController struct {
	// historyListUsecase usecases.HistroyListUsecaseを設定する
	historyListUsecase usecases.HistroyListUsecase
}

// コンストラクタ
func NewHistoryListController(histroyListUsecase *usecases.HistroyListUsecase) *historyListController {
	return &historyListController{
		historyListUsecase: *histroyListUsecase,
	}
}

// ユーザからの入力情報をもとにチャットデータを作成する
func (u *historyListController) CreateChatData(c *gin.Context) {
	var userInput domain.HistoryList
	err := c.BindJSON(&userInput)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// user_inputが空の場合、エラー
	if userInput.UserInput == "" {
		c.JSON(400, gin.H{"error": "質問したい内容を入力してください！"})
		return
	}

	switch userInput.UserInput {

	case "こんにちは":
		//「こんにちは」に適したレスポンスを登録する
		userInput.BotResponse = "こんにちは。"
		err = u.historyListUsecase.CreateChatData(&userInput)
	case "今何時？":
		//「今何時？」の返答として、リクエスト時刻を登録・返却する
		nowTime := common.SetNowTime()
		userInput.BotResponse = fmt.Sprintf("%sです", nowTime)
		jst := common.GetNowTime()
		//取得した現在時刻を設定
		userInput.ResponseTimestamp = jst
		err = u.historyListUsecase.CreateChatData(&userInput)
	case "今日の東京の天気は？":
		//OpenWeatherAPIを使用して天気情報を取得する
		currentWeather, err := common.GetWeather()

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		userInput.BotResponse = fmt.Sprintf("%sです", currentWeather)
		err = u.historyListUsecase.CreateChatData(&userInput)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	default:
		c.JSON(400, gin.H{"error": "現在は「「こんにちは」「今何時？」「今日の東京の天気は？」のみ対応しています"})
		//switchを抜ける
		return
	}

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"response": userInput})
}

// チャットボットとの過去10件のやりとりを取得する
func (u *historyListController) FindAChatDataList(c *gin.Context) {
	historyList, err := u.historyListUsecase.FindChatDataList()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"history_list": historyList})

}
