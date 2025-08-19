package utils

import (
	"bytes"
	"html/template"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

var mdParser goldmark.Markdown
var sanitizer *bluemonday.Policy

func init() {
	mdParser = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
		),
	)

	sanitizer = bluemonday.UGCPolicy()
}

func RenderMarkdown(md string) template.HTML {
	var buf bytes.Buffer
	if err := mdParser.Convert([]byte(md), &buf); err != nil {
		return template.HTML(template.HTMLEscapeString(md))
	}
	safeHTML := sanitizer.SanitizeBytes(buf.Bytes())

	return template.HTML(safeHTML)
}
