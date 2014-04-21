package handlers

import (
	"github.com/martini-contrib/render"
	"github.com/slogsdon/b/models"
	"github.com/slogsdon/b/util"
	"net/http"
	"strings"
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
	var (
		filename   string
		raw        string
		categories []string
	)
	root := util.Config().App.PostsDir

	if err := req.ParseForm(); err != nil {
		panic(err)
	}

	filename = req.Form["filename"][0]
	raw = req.Form["raw"][0]
	hm, _ := models.ParsePostContent([]byte(raw), "md")
	categories = hm.Categories

	if err := util.MakeDir(root + "/" + strings.Join(categories, "/")); err != nil {
		panic(err)
	}

	if err := util.WriteFile(filename, raw); err != nil {
		panic(err)
	}

	r.Data(204, []byte("Created"))
}
