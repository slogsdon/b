package handlers

import (
	"github.com/martini-contrib/render"
	"github.com/slogsdon/b/models"
	"github.com/slogsdon/b/util"
	"net/http"
)

type Api struct {
	Posts ApiPosts
}

func (a Api) Index(r *http.Request, rw http.ResponseWriter) string {
	return "hello"
}

type ApiPosts struct{}

func (ap ApiPosts) Index(r render.Render) {
	root := util.Config().App.PostsDir
	posts := models.GetAllPosts(root)

	r.JSON(200, posts)
}

func (ap ApiPosts) Create(r render.Render, req *http.Request) {
	root := util.Config().App.PostsDir

	if err := req.ParseForm(); err != nil {
		panic(err)
	}

	if err := models.SavePost(root, req.Form); err == nil {
		r.Data(204, []byte("Created"))
	} else {
		r.Data(500, []byte(""))
	}
}
