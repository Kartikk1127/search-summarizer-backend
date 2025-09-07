package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Kartikk1127/search-summarizer/internal/summarizer"
	"github.com/Kartikk1127/search-summarizer/pkg/model"
	"github.com/Kartikk1127/search-summarizer/pkg/utils"
	"github.com/gin-gonic/gin"
)

// for setting up routes,cors and middleware

var llms = []string{"meta-llama/llama-4-scout-17b-16e-instruct"}

func SetupRouter(c *gin.Context) {
	var request model.SummarizeRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, model.SummarizeResponse{Error: "invalid request"})
		return
	}

	text, err := summarizer.ExtractTextAndUrl(request.URL)
	if err != nil {
		c.JSON(http.StatusOK, model.SummarizeResponse{Error: "Content Unavailable"})
		return
	}
	fmt.Println(text)

	modelToUse := llms[0]

	prompt := fmt.Sprintf(`Summarize the following content in exactly 5 bullet points. Format each point as:
		- [Point text]
		
		Content to summarize:
		%s`, text)
	llmRequest := utils.BuildLLMRequest(
		modelToUse,
		fmt.Sprintf(prompt, text),
		400,
		0.3,
		false)

	payloadBytes, err := json.Marshal(llmRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.SummarizeResponse{Error: "Failed to serialize payload"})
		return
	}

	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(payloadBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.SummarizeResponse{Error: "Failed to create request"})
		return
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GROQ_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.SummarizeResponse{Error: "failed to call LLM"})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		c.JSON(http.StatusInternalServerError, model.SummarizeResponse{Error: fmt.Sprintf("LLM API error: %s", string(b))})
		return
	}

	var llmResp model.LLMResponse
	if err := json.NewDecoder(resp.Body).Decode(&llmResp); err != nil {
		c.JSON(http.StatusInternalServerError, model.SummarizeResponse{Error: err.Error()})
		return
	}

	generatedText := llmResp.GetGeneratedText()
	if generatedText == "" {
		c.JSON(http.StatusInternalServerError, model.SummarizeResponse{Error: "empty LLM response"})
		return
	}

	//summary := []string{}
	summarizedResponse := utils.ParseBulletPoints(generatedText)
	if len(summarizedResponse) == 0 {
		c.JSON(http.StatusInternalServerError, model.SummarizeResponse{Error: "failed to parse bullet points"})
		return
	}

	c.JSON(http.StatusOK, model.SummarizeResponse{
		Summary: summarizedResponse,
	})
}
