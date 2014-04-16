package handlers

import (
	"github.com/martini-contrib/render"
	"github.com/slogsdon/b/db"
	"github.com/slogsdon/b/models"
	"net/http"
)

type Api struct {
	Posts apiPosts
}

func (a Api) Index(r *http.Request, rw http.ResponseWriter) string {
	return "hello"
}

type apiPosts struct{}

func (ap apiPosts) Index(r render.Render) {
	var posts []models.Post
	db.DB.Order("published_at").Find(&posts)

	r.JSON(200, posts)
}
