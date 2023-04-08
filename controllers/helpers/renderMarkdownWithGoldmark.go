package helpers

import (
	"bytes"
	"io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func RenderMarkdownWithGoldmark(md string) string {
	var buf bytes.Buffer
	markdown := goldmark.New(goldmark.WithExtensions(extension.GFM))
	if err := markdown.Convert([]byte(md), io.Writer(&buf)); err != nil {
		return md // If there is an error, return the original markdown string
	}
	return buf.String()
}
