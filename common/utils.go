package common

import (
	"fmt"
	"strings"
	"time"
)

// 現在時刻（JST）取得処理
func GetNowTime() string {
	//locationを「東京」に設定
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("現在時刻を取得できませんでした:", err)
		return ""
	}
	// JSTで現在時刻を取得

	timestampString := time.Now().In(location).Format("2006-01-02 15:04:05")
	return timestampString
}

// SetNowTimeを実装
func SetNowTime() string {

	//locationを「東京」に設定
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("エラー:", err)
		return ""
	}
	// JSTで現在時刻を取得
	now := time.Now().In(location)

	// 現在時刻から時と分を取得
	hour, min := now.Hour(), now.Minute()
	// 結果を出力
	return fmt.Sprintf("%d時%d分", hour, min)
}

// 「。」で改行「\n」に変換する処理
func ReplacePeriodWithNewline(text string) string {

	replacer := strings.NewReplacer(
		"。", "。\n",
	)
	covText := replacer.Replace(text)
	return covText
}
