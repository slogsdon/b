package models

import (
	"github.com/slogsdon/b/util"
	"gopkg.in/yaml.v1"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// SavePost writes a new file or a file's new contents to storage.
func SavePost(root string, form map[string][]string) error {
	var (
		filename string
		raw      string
	)
	if _, ok := form["filename"]; ok {
		filename = form["filename"][0]
	}
	if _, ok := form["raw"]; ok {
		raw = form["raw"][0]
	}
	hm, _ := ParsePostContent([]byte(raw), "md")
	categories := strings.Join(hm.Categories, string(os.PathSeparator)) + string(os.PathSeparator)

	err := util.MakeDir(root + string(os.PathSeparator) + categories)
	if err != nil {
		return err
	}

	fullpath := root + string(os.PathSeparator) + categories + filename
	return util.WriteFile(fullpath, raw)
}

// GetAllPosts returns all posts from the storage system by name.
func GetAllPosts(root string) []Post {
	var posts []Post

	for _, f := range util.ReadDir(root) {
		posts = append(posts, preparePost(f))
	}

	return posts
}

// GetPost returns a single post from the storage system by name.
func GetPost(name, root string) Post {
	var post Post

	for _, f := range util.ReadDir(root) {
		if f.Filename == name {
			post = preparePost(f)
			break
		}
	}

	return post
}

// ParsePostContent parses the HeadMatter and HTML from a raw post.
func ParsePostContent(contents []byte, t string) (HeadMatter, template.HTML) {
	m, c := parseHeadMatter(contents)

	switch t {
	case "md", "mdown", "markdown":
		c = util.Markdown(c)
	}

	return m, template.HTML(string(c))
}

// ParsePostSlugAndType parses a post's slug and type from
// its filename. The file extension is used for the post type.
// The slug is grabbed from the basename sans a prefixed date
// used for organization.
// It returns the post's slug and type.
func ParsePostSlugAndType(filename string) (string, string) {
	filenameNoDate := strings.Join(strings.Split(filename, "-")[3:], "-")
	split := strings.Split(filenameNoDate, ".")
	slug := strings.ToLower(strings.Join(split[:len(split)-1], "."))
	t := strings.ToLower(split[len(split)-1])
	return slug, t
}

func preparePost(f util.FileReading) Post {
	// Read file contents
	contents, _ := ioutil.ReadFile(f.Filename)

	// Grab slug and type from filename
	slug, t := ParsePostSlugAndType(f.Info.Name())

	// Parse our content/head matter from our file
	// Return our prepared Post
	head, formattedContents := ParsePostContent(contents, t)
	time, _ := time.Parse("2006-01-02 15:04:05", head.Date)
	return Post{
		Title:       head.Title,
		Slug:        slug,
		Content:     formattedContents,
		HeadMatter:  head,
		Filename:    f.Info.Name(),
		Directory:   strings.Replace(f.Filename, string(os.PathSeparator)+f.Info.Name(), "", 1),
		Type:        t,
		Raw:         string(contents),
		UpdatedAt:   f.Info.ModTime(),
		PublishedAt: time,
	}
}

// Represents the possible data contained within the
// head matter section of a post, fenced with leading
// and following --- lines.
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
