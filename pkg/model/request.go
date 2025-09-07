package model

type SummarizeRequest struct {
	URL string `json:"url" binding:"required"`
}
