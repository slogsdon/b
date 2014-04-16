package models

import "time"

type Category struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	ParentId  int64     `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
