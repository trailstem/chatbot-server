package domain

import "time"

type HistoryList struct {
	ID                int       `json:"id"`
	UserInput         string    `json:"user_input"`
	BotResponse       string    `json:"bot_response"`
	ResponseTimestamp time.Time `json:"response_timestamp"`
}

type WeatherData struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}
