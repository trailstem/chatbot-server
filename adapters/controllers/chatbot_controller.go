package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/trailstem/chatbot-server/common"
	"github.com/trailstem/chatbot-server/domain"
	"github.com/trailstem/chatbot-server/usecases"
)

// コントローラーのインターフェース
type historyListController struct {
	// ユースケースのインターフェース
	historyListUsecase usecases.HistroyListUsecase
}

// コンストラクタ
func NewHistoryListController(histroyListUsecase *usecases.HistroyListUsecase) *historyListController {
	// コントローラーのインスタンスを生成
	return &historyListController{
		historyListUsecase: *histroyListUsecase,
	}
}

// ユーザからの入力情報をもとにチャットデータを作成する
func (u *historyListController) CreateChatData(c *gin.Context) {
	var userReq domain.HistoryList
	err := c.BindJSON(&userReq)
	if err != nil {
		common.RespondWithError(c, 400, "ユーザの入力をバインドできませんでした", &userReq, err)
		return
	}

	//バリデーションチェック
	err = validation.Validate(userReq.UserInput,
		validation.Required,
	)
	if err != nil {
		common.RespondWithError(c, 400, "チャットしたい内容を入力してください", &userReq, err)
		return
	}

	// ユーザの入力によってチャットボットの応答を変える
	switch userReq.UserInput {

	case "こんにちは":
		userReq.BotResponse = "こんにちは。"

	case "今何時？":
		nowTime := common.SetNowTime()
		userReq.BotResponse = fmt.Sprintf("%sです", nowTime)

	case "今日の東京の天気は？":
		// 天気情報を取得
		currentWeather, err := common.GetWeather()
		if err != nil {
			common.RespondWithError(c, 400, "天気情報の取得に失敗しました", &userReq, err)
			return
		}
		userReq.BotResponse = fmt.Sprintf("%sです", currentWeather)

	default:
		// ChatGPTを使用して、応答テキスト生成
		res, err := common.ChatGPT(userReq.UserInput)
		if err != nil {
			common.RespondWithError(c, 400, "ChatGPTの使用に失敗しました", &userReq, err)
			return
		}
		userReq.BotResponse = res
	}

	// チャットデータ作成
	err = u.historyListUsecase.CreateChatData(&userReq)
	if err != nil {
		common.RespondWithError(c, 400, "チャットデータ作成に失敗しました", &userReq, err)
		return
	}
	c.JSON(200, gin.H{"response": userReq})
}

// チャットボットとの過去10件のやりとりを取得する
func (u *historyListController) FindAChatDataList(c *gin.Context) {
	// 実際に過去10件のやりとりを取得する
	historyList, err := u.historyListUsecase.FindChatDataList()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"history_list": historyList})
}
