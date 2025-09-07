package model

type ChatMessage struct {
	Role    string      `json:"role" binding:"required"`
	Content interface{} `json:"content" binding:"required"`
}

type LLMRequest struct {
	Model       string        `json:"model" binding:"required"`
	Messages    []ChatMessage `json:"messages" binding:"required"`
	MaxTokens   int           `json:"max_tokens" binding:"required"`
	Temperature float64       `json:"temperature" binding:"required"`
	Stream      bool          `json:"stream" binding:"required"`
}
