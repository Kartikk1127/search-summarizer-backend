package server

import (
	"fmt"
	"net/http"

	"github.com/Kartikk1127/search-summarizer/internal/summarizer"
	"github.com/gin-gonic/gin"
)

// for setting up routes,cors and middleware

type SummarizeRequest struct {
	URL string `json:"url" binding:"required"`
}

type SummarizeResponse struct {
	Summary []string `json:"summary,omitempty"`
	Error   string   `json:"error,omitempty"`
}

func SetupRouter(c *gin.Context) {
	var request SummarizeRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, SummarizeResponse{Error: "invalid request"})
		return
	}

	text, err := summarizer.ExtractTextAndUrl(request.URL)
	if err != nil {
		c.JSON(http.StatusOK, SummarizeResponse{Error: "Content Unavailable"})
		return
	}
	fmt.Println(text)

	// call LLM with text
	bullets := []string{
		"Dummy summary 1",
		"Dummy summary 2",
	}
	c.JSON(http.StatusOK, SummarizeResponse{Summary: bullets})
}
