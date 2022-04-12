// Markdown parsing and rendering functions
package main

import (
	"bytes"

	markdown "github.com/teekennedy/goldmark-markdown"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
	goldmark.WithRenderer(markdown.NewRenderer()),
)

func ParseEntries(source []byte) bytes.Buffer {
	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}
	return buf
}
