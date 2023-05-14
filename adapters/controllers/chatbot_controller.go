package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
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
		respondWithError(c, 400, "ユーザの入力をバインドできませんでした", &userInput, err)
		return
	}

	//バリデーションチェック
	err = validation.Validate(userInput.UserInput,
		validation.Required,
	)
	if err != nil {
		respondWithError(c, 400, "質問したい内容を入力してください", &userInput, err)
		return
	}

	switch userInput.UserInput {
	case "こんにちは":
		userInput.BotResponse = "こんにちは。"
	case "今何時？":
		nowTime := common.SetNowTime()
		userInput.BotResponse = fmt.Sprintf("%sです", nowTime)
	case "今日の東京の天気は？":
		currentWeather, err := common.GetWeather()
		if err != nil {
			respondWithError(c, 400, "天気情報の取得に失敗しました", &userInput, err)
			return
		}
		userInput.BotResponse = fmt.Sprintf("%sです", currentWeather)
	default:
		res, err := common.ChatGPT(userInput.UserInput)
		if err != nil {
			respondWithError(c, 400, "ChatGPTの使用に失敗しました", &userInput, err)
			return
		}
		userInput.BotResponse = res
	}

	err = u.historyListUsecase.CreateChatData(&userInput)
	if err != nil {
		respondWithError(c, 400, "チャットデータ作成に失敗しました", &userInput, err)
		return
	}

	c.JSON(200, gin.H{"response": userInput})
}

// respondWithError sends an error response with the specified status code and message.
func respondWithError(c *gin.Context, status int, message string, userInput *domain.HistoryList,
	err error) {
	if err != nil {
		c.JSON(status, gin.H{"error": message + ": " + err.Error()})
	} else {
		c.JSON(status, gin.H{"response": userInput})
	}
}

// チャットボットとの過去10件のやりとりを取得する
func (u *historyListController) FindAChatDataList(c *gin.Context) {
	historyList, err := u.historyListUsecase.FindChatDataList()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	//histryListの

	c.JSON(200, gin.H{"history_list": historyList})

}
