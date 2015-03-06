package handlers

import (
	"fmt"
	"net/http"
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
	fmt.Fprintf(rw, "admin edit post")
}
