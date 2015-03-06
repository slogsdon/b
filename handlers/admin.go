package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/slogsdon/b/models"
	"github.com/slogsdon/b/util"
)

type Admin struct {
	Posts AdminPosts
}

// Index returns the admin dashboard.
func (a Admin) Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "admin index")
}

type AdminPosts struct{}

// Index lists all posts.
func (a AdminPosts) Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "admin posts index")
}

// New allows for adding new posts.
func (a AdminPosts) New(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "admin new post")
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

	var data []byte
	if found {
		data, _ = json.Marshal(post)
		rw.Write(data)
	} else {
		rw.WriteHeader(404)
		rw.Write(data)
	}
}
