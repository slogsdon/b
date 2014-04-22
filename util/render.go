package util

import (
	"github.com/russross/blackfriday"
)

// Markdown presents MarkdownCommon with a few minor changes
func Markdown(str interface{}) []byte {
	// this did use blackfriday.MarkdownCommon, but it was stripping out <script>
	var c []byte

	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_USE_XHTML
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")

	// set up the parser
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	extensions |= blackfriday.EXTENSION_FOOTNOTES

	switch str.(type) {
	default:
		return c
	case string:
		c = []byte(str.(string))
	case []byte:
		c = []byte(str.([]byte))
	}

	return blackfriday.Markdown(c, renderer, extensions)
}
