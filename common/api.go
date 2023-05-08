package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	// "github.com/joho/godotenv"
	"github.com/trailstem/chatbot-server/domain"
)

// OpenWeatherAPIを使用して天気情報を取得する
func GetWeather() (string, error) {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	wk := os.Getenv("WEATHER_KEY")

	//都市「東京」、言語「日本語」、温度単位「`C」に設定
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&lang=ja&units=metric", "Tokyo", wk)

	//urlをもとにリクエスト処理実行
	res, err := http.Get(url)
	// 天気情報が取得できない場合エラー
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// レスポンスの処理
	bytes, _ := io.ReadAll(res.Body)

	// bodyをJSONに変換
	weatherData := domain.WeatherData{}
	if err := json.Unmarshal([]byte(bytes), &weatherData); err != nil {
		return "", err
	}

	description := weatherData.Weather[0].Description
	fmt.Println(description)
	return description, nil
}
