package common

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trailstem/chatbot-server/domain"
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
	t := time.Now().In(location)            // 現在時刻をJSTで取得
	formattedTime := t.Format(time.RFC3339) // ISO 8601形式にフォーマット

	// 正規表現でタイムゾーン情報を削除
	re := regexp.MustCompile(`[-+]\d{2}:\d{2}$`)
	formattedTime = re.ReplaceAllString(formattedTime, "")
	return formattedTime
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

// 指定のステータスコードとメッセージでエラー応答を送信
func RespondWithError(c *gin.Context, status int, message string, userReq *domain.HistoryList,
	err error) {
	// エラー内容を出力
	if err != nil {
		c.JSON(status, gin.H{"error": message + ": " + err.Error()})
	} else {
		// エラー内容がない場合はチャットデータを設定
		c.JSON(status, gin.H{"response": userReq})
	}
}
