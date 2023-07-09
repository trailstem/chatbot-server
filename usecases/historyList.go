package usecases

import (
	"github.com/trailstem/chatbot-server/adapters/gateways"
	"github.com/trailstem/chatbot-server/common"
	"github.com/trailstem/chatbot-server/domain"
)

// ユースケースのインターフェース
type HistroyListUsecase struct {
	historyListRepo gateways.SpeakBotRepo
}

// コンストラクタ
func NewHistoryListUsecase(historyListRepo gateways.SpeakBotRepo) *HistroyListUsecase {
	return &HistroyListUsecase{
		historyListRepo: historyListRepo,
	}
}

// チャットデータ作成する実際のロジック
func (u *HistroyListUsecase) CreateChatData(userInput *domain.HistoryList) error {
	//取得した現在時刻を設定
	userInput.ResponseTimestamp = common.GetNowTime()
	//repositoryのCreateChatDataを呼び出し
	return u.historyListRepo.CreateChatData(userInput)
}

// チャットボットとの過去10件のやりとりを取得する実際のロジック
func (u *HistroyListUsecase) FindChatDataList() (*[]domain.HistoryList, error) {
	//空の*[]domain.HistoryListを作成
	historyList, err := u.historyListRepo.FindChatDataList()

	if err != nil {
		return nil, err
	}
	//取得したhistoryListをrangeで回して、replacePeriodWithNewlineでuserInputとbotResponseの改行を置換
	for i := range *historyList {
		(*historyList)[i].UserInput = common.ReplacePeriodWithNewline((*historyList)[i].UserInput)
		(*historyList)[i].BotResponse = common.ReplacePeriodWithNewline((*historyList)[i].BotResponse)
	}
	return historyList, nil

}
