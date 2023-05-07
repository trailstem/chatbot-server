package common

import (
	"fmt"
	"time"
)

// 現在時刻（JST）取得処理
func GetNowTime() time.Time {
	//locationを「東京」に設定
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("エラー:", err)
		return time.Time{}
	}
	// JSTで現在時刻を取得
	jstNowTime := time.Now().In(location)
	return jstNowTime
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
