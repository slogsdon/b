package models

import (
	"github.com/slogsdon/b/util"
	"html/template"
	"time"
)

type Post struct {
	Id          int64           `json:"id"`
	Title       string          `json:"title"`
	Slug        string          `json:"slug"`
	ParentId    int64           `json:"parent_id"`
	Content     template.HTML   `json:"content"`
	Raw         string          `json:"raw"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	PublishedAt time.Time       `json:"published_at"`
	HeadMatter  util.HeadMatter `json:"head_matter"`
	Filename    string          `json:"-"`
	Directory   string          `json:"-"`
	Type        string          `json:"-"`
}
