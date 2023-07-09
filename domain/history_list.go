package domain

// チャットデータの構造体
type HistoryList struct {
	UserInput         string `json:"user_input"`
	BotResponse       string `json:"bot_response"`
	ResponseTimestamp string `json:"response_timestamp"`
}

// openweathermapのレスポンス構造体
type WeatherData struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}
