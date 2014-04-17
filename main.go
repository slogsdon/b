package main

import (
	"github.com/go-martini/martini"
	// "github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/slogsdon/b/handlers"
)

func main() {
	// Set up our Martini instance
	m := martini.Classic()

	// Middleware
	// m.Use(gzip.All())
	m.Use(render.Renderer())

	// Routes
	m.Group("/admin", func(r martini.Router) {
		a := handlers.Admin{}

		r.Get("", a.Index)
		r.Get("/posts", a.Posts.Index)
		r.Get("/posts/edit/**", a.Posts.Edit)

	}, render.Renderer(render.Options{
		Layout: "admin/layout",
	}))

	m.Group("/api", func(r martini.Router) {
		a := handlers.Api{}

		r.Get("", a.Index)
		r.Get("/posts", a.Posts.Index)
	})

	// Serve from static if possible
	m.Use(martini.Static("_site"))
	m.Run()
}
