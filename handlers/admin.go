package handlers

import (
	"net/http"
	"os"

	"github.com/slogsdon/b/models"
	"github.com/slogsdon/b/render"
	"github.com/slogsdon/b/util"
)

type Admin struct {
	Posts AdminPosts
}

// Index returns the admin dashboard.
func (a Admin) Index(rw http.ResponseWriter, r *http.Request) {
	render.HTML(rw, "admin/index")
}

type AdminPosts struct{}

// Index lists all posts.
func (a AdminPosts) Index(rw http.ResponseWriter, r *http.Request) {
	render.HTML(rw, "admin/posts/index")
}

// New allows for adding new posts.
func (a AdminPosts) New(rw http.ResponseWriter, r *http.Request) {
	render.HTML(rw, "admin/posts/new")
}

// Edit loads a post for editing.
func (a AdminPosts) Edit(rw http.ResponseWriter, r *http.Request) {
	var post models.Post
	_, params, _ := util.Mux.Lookup(r.Method, r.URL.Path)

	root := util.Config().App.PostsDir
	path := models.ParsePostId(params.ByName("id"))
	file := root + path
	found := false

	for _, p := range models.GetAllPosts(root) {
		if p.Directory+string(os.PathSeparator)+p.Filename == file {
			post = p
			found = true
		}
	}

	if found {
		render.HTML(rw, "admin/posts/edit", post)
	} else {
		rw.WriteHeader(404)
	}
}
