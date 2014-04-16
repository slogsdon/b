package models

import "time"

type Post struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	ParentId    int64     `json:"parent_id"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`
}
