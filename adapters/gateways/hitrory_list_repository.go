package gateways

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/trailstem/chatbot-server/domain"
)

type HisoryListRepository struct {
	db *sql.DB
}

// コンストラクタ
func NewHistoryListRepository(db *sql.DB) *HisoryListRepository {
	return &HisoryListRepository{db: db}
}

// ChatBotインターフェースを実装
func (r *HisoryListRepository) CreateChatData(userInput *domain.HistoryList) error {
	//プリペアドステートメントを使用
	stmt, err := r.db.Prepare("INSERT INTO history_list (user_input, bot_response, response_timestamp) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userInput.UserInput, userInput.BotResponse, userInput.ResponseTimestamp)
	if err != nil {
		return err
	}
	return nil
}

func (r *HisoryListRepository) FindChatDataList() (*[]domain.HistoryList, error) {
	//history_listテーブルからresponse_timestampの降順でソートして10件取得
	rows, err := r.db.Query("SELECT user_input, bot_response, response_timestamp FROM history_list ORDER BY response_timestamp DESC LIMIT 10")
	if err != nil {
		return nil, err
	}

	//取得したデータを格納する構造体HistoryListにバインド
	var historyList []domain.HistoryList
	for rows.Next() {
		var history domain.HistoryList

		err := rows.Scan(&history.UserInput, &history.BotResponse, &history.ResponseTimestamp)
		if err != nil {
			return nil, err
		}

		historyList = append(historyList, history)
	}
	return &historyList, nil
}
