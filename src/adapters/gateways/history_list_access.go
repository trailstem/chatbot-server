package gateways

import "github.com/trailstem/chatbot-server/domain"

type SpeakBotRepo interface {
	//チャットデータを作成
	CreateChatData(userInput *domain.HistoryList) error
	//条件10件
	FindChatDataList() (*[]domain.HistoryList, error)
}
