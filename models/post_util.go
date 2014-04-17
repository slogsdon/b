package models

import (
	"github.com/slogsdon/b/util"
	"io/ioutil"
	"strings"
	"time"
)

func GetAllPosts() []Post {
	var posts []Post
	root := "./_posts"

	for _, f := range util.ReadDir(root) {
		contents, _ := ioutil.ReadFile(f.Filename)
		filenameNoDate := strings.Join(strings.Split(f.Info.Name(), "-")[3:], "-")
		split := strings.Split(filenameNoDate, ".")
		slug := strings.ToLower(strings.Join(split[:len(split)-1], "."))
		t := strings.ToLower(split[len(split)-1])
		head, formattedContents := util.ParseContent(contents, t)
		time, _ := time.Parse("2006-01-02 15:04:05", head.Date)
		posts = append(posts, Post{
			Title:       head.Title,
			Slug:        slug,
			Content:     formattedContents,
			HeadMatter:  head,
			UpdatedAt:   f.Info.ModTime(),
			PublishedAt: time,
			Filename:    f.Info.Name(),
			Directory:   strings.Replace(f.Filename, "/"+f.Info.Name(), "", 1),
			Type:        t,
			Raw:         string(contents),
		})
	}

	return posts
}
