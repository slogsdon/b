package models

import (
	"html/template"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/slogsdon/b/util"
	"gopkg.in/yaml.v1"
)

// ParsePostId parses a file path from an url parameter.
func ParsePostId(id string) string {
	split := strings.Split(id, "_")
	path := split[0]

	if len(split) == 1 {
		return path
	}

	if path == "" {
		return strings.Join(split[1:], "_")
	}

	if strings.Contains(path, "-") {
		path = strings.Join(strings.Split(path, "-"), "/")
	}

	return path + "/" + strings.Join(split[1:], "_")
}

// SavePost writes a new file or a file's new contents to storage.
func SavePost(root string, p interface{}) error {
	var (
		categories string
		filename   string
		raw        string
		hm         string
	)

	switch p := p.(type) {
	case Post:
		//json
		filename = p.Filename
		categories = strings.Join(p.HeadMatter.Categories, string(os.PathSeparator))
		raw = p.Raw
		hm = FormatPostHeadMatter(p.HeadMatter)
	default:
		var (
			title            string
			date             string
			c                []string
			meta_description string
			feature_image    string
		)
		// x-www-form-urlencoded
		form := url.Values(p.(url.Values))
		if _, ok := form["filename"]; ok {
			filename = form["filename"][0]
		}
		if _, ok := form["raw"]; ok {
			raw = form["raw"][0]
		}
		if _, ok := form["title"]; ok {
			title = form["title"][0]
		}
		if _, ok := form["date"]; ok {
			date = form["date"][0]
		}
		if _, ok := form["categories"]; ok {
			c = form["categories"]
		}
		if _, ok := form["meta_description"]; ok {
			meta_description = form["meta_description"][0]
		}
		if _, ok := form["feature_image"]; ok {
			feature_image = form["feature_image"][0]
		}
		hm = FormatPostHeadMatter(HeadMatter{
			Title:           title,
			Date:            date,
			Categories:      c,
			MetaDescription: meta_description,
			FeatureImage:    feature_image,
		})
		categories = strings.Join(c, string(os.PathSeparator))
	}

	categories += string(os.PathSeparator)
	raw = hm + raw

	dest := root + string(os.PathSeparator) + categories
	err := util.MakeDir(dest)

	if err != nil {
		return err
	}

	fullpath := dest + filename

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
func ParsePostContent(contents []byte, t string) template.HTML {
	var c []byte

	switch t {
	case "md", "mdown", "markdown":
		c = util.Markdown(contents)
	}

	return template.HTML(string(c))
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
	head, contentsNoHead := ParsePostHeadMatter(contents)
	formattedContents := ParsePostContent(contentsNoHead, t)
	time, _ := time.Parse("2006-01-02 15:04:05", head.Date)
	return Post{
		Title:      head.Title,
		Slug:       slug,
		Content:    formattedContents,
		HeadMatter: head,
		Filename:   f.Info.Name(),
		Directory:  strings.Replace(f.Filename, string(os.PathSeparator)+f.Info.Name(), "", 1),
		Type:       t,
		Raw:        string(contentsNoHead),
		// CreatedAt:   f.Info.Sys().Ctim,
		UpdatedAt:   f.Info.ModTime(),
		PublishedAt: time,
	}
}

// Represents the possible data contained within the
// head matter section of a post, fenced with leading
// and following --- lines.
type HeadMatter struct {
	Title           string   `json:"title",yaml:"title,omitempty"`
	Date            string   `json:"date",yaml:"date,omitempty"`
	Categories      []string `json:"categories",yaml:"categories,omitempty"`
	MetaDescription string   `json:"meta_description",yaml:"meta_description,omitempty"`
	FeatureImage    string   `json:"feature_image",yaml:"feature_image,omitempty"`
}

func FormatPostHeadMatter(hm HeadMatter) string {
	var r string
	b, _ := yaml.Marshal(hm)
	if len(b) > 0 {
		r = "---\n" + string(b) + "---\n\n"
	}
	return r
}

func ParsePostHeadMatter(contents []byte) (HeadMatter, []byte) {
	m := HeadMatter{}
	c := string(contents)

	if strings.Count(c, "---") >= 2 {
		split := strings.Split(c, "---")
		_ = yaml.Unmarshal([]byte(split[1]), &m)
		c = strings.Trim(strings.Join(split[2:], "---"), "\r\n")
	}

	return m, []byte(c)
}
