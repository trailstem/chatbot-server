package common

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	// "github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"github.com/trailstem/chatbot-server/domain"
)

// OpenWeatherAPIを使用して天気情報を取得する
func GetWeather() (string, error) {
	// .envファイルを読み込む
	wk := os.Getenv("WEATHER_KEY")

	//都市「東京」、言語「日本語」、温度単位「`C」に設定
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&lang=ja&units=metric", "Tokyo", wk)

	//urlをもとにリクエスト処理実行
	res, err := http.Get(url)
	// 天気情報が取得できない場合エラー
	if err != nil {
		return "", err
	}
	// レスポンスのボディを閉じる
	defer res.Body.Close()

	// レスポンスの処理
	bytes, _ := io.ReadAll(res.Body)

	// bodyをJSONに変換
	weatherData := domain.WeatherData{}
	if err := json.Unmarshal([]byte(bytes), &weatherData); err != nil {
		return "", err
	}

	// 天気情報を取得
	description := weatherData.Weather[0].Description
	return description, nil
}

// OpenAI ChatGPTを使用してチャットボットとのやりとりを行う
func ChatGPT(userInput string) (string, error) {

	// envからOPEN_API取得
	oiKey := os.Getenv("OPEN_API")
	client := openai.NewClient(oiKey)
	// チャットボットとのやりとりを行う
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userInput,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	resText := resp.Choices[0].Message.Content
	//画面に表示しやすいように「.」「。」で改行して返却
	return ReplacePeriodWithNewline(resText), nil
}
