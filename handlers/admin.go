package handlers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/slogsdon/b/models"
	"github.com/slogsdon/b/util"
	"os"
)

type Admin struct {
	Posts adminPosts
}

func (a Admin) Index(r render.Render) {
	r.HTML(200, "admin/index", "")
}

type adminPosts struct{}

func (ap adminPosts) Index(r render.Render) {
	r.HTML(200, "admin/posts/index", "")
}

func (ap adminPosts) Edit(params martini.Params, r render.Render) {
	var post models.Post

	root := util.Config().App.PostsDir
	file := root + string(os.PathSeparator) + params["_1"]
	found := false

	for _, p := range models.GetAllPosts(root) {
		if p.Directory+string(os.PathSeparator)+p.Filename == file {
			post = p
			found = true
		}
	}

	if found {
		r.HTML(200, "admin/posts/edit", post)
	} else {
		r.Error(404)
	}
}
