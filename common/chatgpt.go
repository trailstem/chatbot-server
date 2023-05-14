package common

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func ChatGPT(userInput string) (string, error) {

	// envからOPEN_API取得
	oiKey := os.Getenv("OPEN_API")
	client := openai.NewClient(oiKey)
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

	fmt.Println(resp.Choices[0].Message.Content)
	resText := resp.Choices[0].Message.Content
	//画面に表示しやすいように「.」「。」で改行追加
	replacer := strings.NewReplacer(
		"。", "。\n",
	)
	resT := replacer.Replace(resText)
	return resT, nil
}
