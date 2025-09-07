package model

type SummarizeResponse struct {
	Summary []string `json:"summary,omitempty"`
	Error   string   `json:"error,omitempty"`
}
