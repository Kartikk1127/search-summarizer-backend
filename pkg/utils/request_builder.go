package utils

import "github.com/Kartikk1127/search-summarizer/pkg/model"

func BuildLLMRequest(modelName, userText string, maxTokens int, temperature float64, stream bool) model.LLMRequest {
	return model.LLMRequest{
		Model: modelName,
		Messages: []model.ChatMessage{
			{
				Role:    "user",
				Content: userText,
			},
		},
		MaxTokens:   maxTokens,
		Temperature: temperature,
		Stream:      stream,
	}
}
