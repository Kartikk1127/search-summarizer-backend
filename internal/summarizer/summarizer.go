package summarizer

import (
	"fmt"
	"time"

	"github.com/go-shiori/go-readability"
)

func ExtractTextAndUrl(url string) (string, error) {
	article, err := readability.FromURL(url, 10*time.Second)
	if err != nil {
		return "", err
	}

	if len(article.TextContent) < 50 {
		return "", fmt.Errorf("content too short from %s (len=%d)", url, len(article.TextContent))
	}
	return article.TextContent, nil
}
