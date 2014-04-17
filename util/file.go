package util

import (
	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v1"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func ReadDir(dir string) []fileReading {
	var m []fileReading
	return read(dir, m)
}

func ParseContent(contents []byte, t string) (HeadMatter, template.HTML) {
	m, c := parseHeadMatter(contents)

	switch t {
	case "md", "mdown", "markdown":
		c = markdown(c)
	}

	return m, template.HTML(string(c))
}

type fileReading struct {
	Filename string
	Info     os.FileInfo
}

func read(dir string, m []fileReading) []fileReading {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		file := dir + string(os.PathSeparator) + f.Name()
		if f.IsDir() {
			m = read(file, m)
		} else {
			if f.Name()[0:1] != "." {
				m = append(m, fileReading{Filename: file, Info: f})
			}
		}
	}
	return m
}

type HeadMatter struct {
	Title      string   `json:"title"`
	Date       string   `json:"date"`
	Categories []string `json:"categories"`
}

func parseHeadMatter(contents []byte) (HeadMatter, []byte) {
	m := HeadMatter{}
	c := string(contents)

	if strings.Count(c, "---") >= 2 {
		split := strings.Split(c, "---")
		_ = yaml.Unmarshal([]byte(split[1]), &m)
		c = strings.Join(split[2:], "---")
	}

	return m, []byte(c)
}

func markdown(str []byte) []byte {
	// this did use blackfriday.MarkdownCommon, but it was stripping out <script>

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

	return blackfriday.Markdown(str, renderer, extensions)
}
