package usecases

import (
	"github.com/trailstem/chatbot-server/adapters/gateways"
	"github.com/trailstem/chatbot-server/domain"
)

type HistroyListUsecase struct {
	historyListRepo gateways.SpeakBotRepo
}

func NewHistoryListUsecase(historyListRepo gateways.SpeakBotRepo) *HistroyListUsecase {
	return &HistroyListUsecase{
		historyListRepo: historyListRepo,
	}
}

// 実際のロジック処理を記載
func (u *HistroyListUsecase) CreateChatData(userInput *domain.HistoryList) error {
	return u.historyListRepo.CreateChatData(userInput)
}

// チャットボットとの過去10件のやりとりを取得する実際のロジック
func (u *HistroyListUsecase) FindChatDataList() (*[]domain.HistoryList, error) {
	return u.historyListRepo.FindChatDataList()
}
