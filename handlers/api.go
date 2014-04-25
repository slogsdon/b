package handlers

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/slogsdon/b/models"
	"github.com/slogsdon/b/util"
	"net/http"
	"os"
	"strings"
)

type Api struct {
	Posts  ApiPosts
	Render ApiRender
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
	var (
		post models.Post
		err  error
		t    = "urlencoded"
		p    interface{}
	)
	root := util.Config().App.PostsDir
	contentType := req.Header.Get("content-type")

	if strings.Contains(contentType, "application/json") {
		t = "json"
	}

	switch t {
	case "json":
		dec := json.NewDecoder(req.Body)
		err = dec.Decode(&post)
		p = post
	default:
		err = req.ParseForm()
		p = req.Form
	}

	if err != nil {
		r.Data(500, []byte(err.Error()))
		return
	}

	if err = models.SavePost(root, p); err == nil {
		r.Data(204, []byte("Created"))
	} else {
		r.Data(500, []byte(err.Error()))
	}
}

// Show returns a single post.
func (ap ApiPosts) Show(params martini.Params, r render.Render) {
	var post models.Post

	root := util.Config().App.PostsDir
	path := models.ParsePostId(params["id"])
	file := root + string(os.PathSeparator) + path
	found := false

	for _, p := range models.GetAllPosts(root) {
		if p.Directory+string(os.PathSeparator)+p.Filename == file {
			post = p
			found = true
		}
	}

	if found {
		r.JSON(200, post)
	} else {
		r.Error(404)
	}
}

type ApiRender struct{}

// Markdown renders a POST request into HTML.
func (ap ApiRender) Markdown(r render.Render, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		r.Data(500, []byte(err.Error()))
		return
	}

	if _, ok := req.Form["raw"]; !ok {
		r.Data(500, []byte("No Data"))
		return
	}

	raw := req.Form["raw"][0]
	data := util.Markdown([]byte(raw))

	r.JSON(200, apiRenderResponse{Data: string(data)})
}

type apiRenderResponse struct {
	Data string `json:"data"`
}
