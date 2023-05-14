package domain

/*
	OpenAI 「Create chat completion」用構造体
*/

type Message struct {
	Role    string `json:"role"`    // メッセージのロール
	Content string `json:"content"` // メッセージ内容
}

// Parameter用
type OpenAIReq struct {
	Model    string    `json:"model"`    // モデル指定
	Messages []Message `json:"messages"` // request用メッセージ配列
}
type Choice struct {
	Index        int     `json:"index"`         // 選択肢インデックス
	Message      Message `json:"message"`       // レスポンスメッセージ
	FinishReason string  `json:"finish_reason"` // 終了理由
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`     // プロンプトトークン数
	CompletionTokens int `json:"completion_tokens"` // コンプリーショントークン数
	TotalTokens      int `json:"total_tokens"`      // 総トークン数
}

// Response用
type OpenAIRes struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}
