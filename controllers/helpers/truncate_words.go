package helpers

import "strings"

func TruncateWords(content string, wordLimit int) string {
	words := strings.Fields(content)
	if len(words) <= wordLimit {
		return content
	}
	return strings.Join(words[:wordLimit], " ") + "..."
}
