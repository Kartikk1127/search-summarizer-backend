package utils

import (
	"regexp"
	"strings"
)

func ParseBulletPoints(text string) []string {
	lines := strings.Split(text, "\n")
	var bulletPoints []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "• ") {
			bulletPoints = append(bulletPoints, strings.TrimSpace(strings.TrimPrefix(line, "• ")))
		} else if strings.HasPrefix(line, "- ") {
			bulletPoints = append(bulletPoints, strings.TrimSpace(strings.TrimPrefix(line, "- ")))
		} else if matched, _ := regexp.MatchString(`^\d+\.\s+`, line); matched {
			// Handle numbered lists like "1. Point text"
			re := regexp.MustCompile(`^\d+\.\s+`)
			bulletPoints = append(bulletPoints, strings.TrimSpace(re.ReplaceAllString(line, "")))
		} else if line != "" {
			// Fallback for plain text lines
			bulletPoints = append(bulletPoints, line)
		}
	}

	return bulletPoints
}
