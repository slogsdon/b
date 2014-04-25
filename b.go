// B is a static-ish blog application. Can be run
// as a standalone application/server or be used
// to locally manage and deploy posts to a remote
// server.
package b

import (
	"github.com/go-martini/martini"
	// "github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/slogsdon/b/handlers"
	"github.com/slogsdon/b/util"
)

const (
	// Current application version.
	VERSION = "0.0.1"
)

// Entry point for running the application.
// It defines all routes and middleware used
// and starts the underlying server.
func Start() {
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
		r.Get("/posts/new", a.Posts.New)
		r.Get("/posts/:id/edit", a.Posts.Edit)

	}, render.Renderer(render.Options{
		Layout: "admin/layout",
	}))

	m.Group("/api", func(r martini.Router) {
		a := handlers.Api{}

		r.Get("", a.Index)
		r.Get("/posts", a.Posts.Index)
		r.Post("/posts", a.Posts.Create)
		r.Get("/posts/:id", a.Posts.Show)
		r.Get("/render/markdown", a.Render.Markdown)
	})

	// Serve from static if possible
	m.Use(martini.Static(util.Config().App.SiteDir))
	m.Run()
}
