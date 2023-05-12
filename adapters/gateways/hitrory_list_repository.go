package gateways

import (
	"database/sql"
	"time"

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
	_, err = stmt.Exec(userInput.UserInput, userInput.BotResponse, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *HisoryListRepository) FindChatDataList() (*[]domain.HistoryList, error) {
	//history_listテーブルからresponse_timestampの降順でソートして10件取得
	rows, err := r.db.Query("SELECT * FROM history_list ORDER BY response_timestamp DESC LIMIT 10")
	if err != nil {
		return nil, err
	}

	//取得したデータを格納する構造体HistoryListにバインド
	var historyList []domain.HistoryList
	for rows.Next() {
		var history domain.HistoryList
		var timestampStr string

		err := rows.Scan(&history.ID, &history.UserInput, &history.BotResponse, &timestampStr)
		if err != nil {
			return nil, err
		}
		// timestamp, err := time.Parse("2006-01-02 15:04:05", timestampStr)
		// if err != nil {
		// 	return nil, err
		// }
		// history.ResponseTimestamp = timestamp
		historyList = append(historyList, history)
	}
	return &historyList, nil
}
