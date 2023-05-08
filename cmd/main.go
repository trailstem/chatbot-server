package main

import (
	"os"
	"time"

	"github.com/trailstem/chatbot-server/infrastractures"
)

type HistoryList struct {
	UserInput         string    `json:"user_input"`
	BotResponse       string    `json:"bot_response"`
	ResponseTimestamp time.Time `json:"response_timestamp"`
}

func main() {

	r := infrastractures.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
