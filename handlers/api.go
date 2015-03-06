package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/slogsdon/b/models"
	"github.com/slogsdon/b/render"
	"github.com/slogsdon/b/util"
)

type Api struct {
	Posts  ApiPosts
	Render ApiRender
}

// Index returns the api's index.
func (a Api) Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "hello")
}

type ApiPosts struct{}

// Index returns all available posts.
func (a ApiPosts) Index(rw http.ResponseWriter, r *http.Request) {
	root := util.Config().App.PostsDir
	posts := models.GetAllPosts(root)

	render.JSON(rw, posts)
}

// Create allows for the creation of new posts. It returns a 204
// response on creation or a 500 response on error.
func (a ApiPosts) Create(rw http.ResponseWriter, r *http.Request) {
	var (
		post models.Post
		err  error
		t    = "urlencoded"
		p    interface{}
	)
	root := util.Config().App.PostsDir
	contentType := r.Header.Get("content-type")

	if strings.Contains(contentType, "application/json") {
		t = "json"
	}

	switch t {
	case "json":
		dec := json.NewDecoder(r.Body)
		err = dec.Decode(&post)
		p = post
	default:
		err = r.ParseForm()
		p = r.Form
	}

	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	if err = models.SavePost(root, p); err == nil {
		rw.WriteHeader(204)
		rw.Write([]byte("Created"))
	} else {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
	}
}

// Show returns a single post.
func (a ApiPosts) Show(rw http.ResponseWriter, r *http.Request) {
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
		render.JSON(rw, post)
	} else {
		rw.WriteHeader(404)
	}
}

type ApiRender struct{}

// Markdown renders a POST ruest into HTML.
func (a ApiRender) Markdown(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		return
	}

	if _, ok := r.Form["raw"]; !ok {
		rw.WriteHeader(500)
		rw.Write([]byte("No Data"))
		return
	}

	raw := r.Form["raw"][0]
	data := util.Markdown([]byte(raw))

	render.JSON(rw, apiRenderResponse{Data: string(data)})
}

type apiRenderResponse struct {
	Data string `json:"data"`
}
