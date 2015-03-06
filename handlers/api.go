package handlers

import (
	"fmt"
	"net/http"
)

type Api struct {
	Posts  ApiPosts
	Render ApiRender
}

// Index returns the api's index.
func (a Api) Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "api index")
}

type ApiPosts struct{}

// Index returns all available posts.
func (a ApiPosts) Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "api posts index")
}

// Create allows for the creation of new posts. It returns a 204
// response on creation or a 500 response on error.
func (a ApiPosts) Create(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "api post create")
}

// Show returns a single post.
func (a ApiPosts) Show(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "api post show")
}

type ApiRender struct{}

// Markdown renders a POST ruest into HTML.
func (a ApiRender) Markdown(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "api render markdown")
}

type apiRenderResponse struct {
	Data string `json:"data"`
}
