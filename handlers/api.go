package handlers

import (
	"fmt"
	"github.com/martini-contrib/render"
	"github.com/slogsdon/b/models"
	"github.com/slogsdon/b/util"
	"net/http"
)

type Api struct {
	Posts ApiPosts
}

// Index returns the api's index.
func (a Api) Index(r *http.Request, rw http.ResponseWriter) string {
	return "hello"
}

type ApiPosts struct{}

// Index returns all available posts.
func (ap ApiPosts) Index(r render.Render) {
	root := util.Config().App.PostsDir
	posts := models.GetAllPosts(root)

	r.JSON(200, posts)
}

// Create allows for the creation of new posts. It returns a 204
// response on creation or a 500 response on error.
func (ap ApiPosts) Create(r render.Render, req *http.Request) {
	root := util.Config().App.PostsDir

	if err := req.ParseForm(); err != nil {
		r.Data(500, []byte(err.Error()))
	}

	fmt.Println(root)
	fmt.Println(req.Form)

	if err := models.SavePost(root, req.Form); err == nil {
		fmt.Println(err)
		r.Data(204, []byte("Created"))
	} else {
		r.Data(500, []byte(err.Error()))
	}
}
