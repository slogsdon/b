package handlers

import (
	"github.com/martini-contrib/render"
)

type Admin struct{}

func (a Admin) Index(r render.Render) {
	r.HTML(200, "admin/index", "")
}

func (a Admin) PostsIndex(r render.Render) {
	r.HTML(200, "admin/posts/index", "")
}
