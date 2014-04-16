package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/slogsdon/b/db"
	"github.com/slogsdon/b/handlers"
)

func main() {
	// Set up our Martini instance
	m := martini.Classic()
	m.Map(&db.DB)

	// Middleware
	m.Use(gzip.All())
	m.Use(render.Renderer())

	// Routes
	m.Group("/admin", func(r martini.Router) {
		a := handlers.Admin{}

		m.Get("", a.Index)
		m.Get("/posts", a.PostsIndex)
	})

	m.Group("/api", func(r martini.Router) {
		a := handlers.Api{}

		m.Get("", a.Index)
		m.Get("/posts", a.Posts.Index)
	})

	// Serve from static if possible
	m.Use(martini.Static("_site"))
	m.Run()
}
