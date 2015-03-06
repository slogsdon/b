// B is a static-ish blog application. Can be run
// as a standalone application/server or be used
// to locally manage and deploy posts to a remote
// server.
package b

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/slogsdon/b/handlers"
)

const (
	// Current application version.
	VERSION = "0.0.1"
)

var (
	defaultOptions = Options{
		Port: "3000",
	}
	static = http.FileServer(http.Dir("./_site"))
)

// Entry point for running the application.
// It defines all routes and middleware used
// and starts the underlying server.
func Start(args ...interface{}) {
	opts := getOptions(args)
	m := httprouter.New()

	admin := handlers.Admin{}
	m.Handler("GET", "/admin/posts/new", wrap(admin.Posts.New))
	m.Handler("GET", "/admin/posts/edit/:id", wrap(admin.Posts.Edit))
	m.Handler("GET", "/admin/posts", wrap(admin.Posts.Index))
	m.Handler("GET", "/admin", wrap(admin.Index))

	api := handlers.Api{}
	m.Handler("GET", "/api/render/markdown", wrap(api.Render.Markdown))
	m.Handler("GET", "/api/posts/:id", wrap(api.Posts.Show))
	m.Handler("GET", "/api/posts", wrap(api.Posts.Index))
	m.Handler("POST", "/api/posts", wrap(api.Posts.Create))
	m.Handler("GET", "/api", wrap(api.Index))

	m.Handler("GET", "/", static)

	http.Handle("/", m)
	http.ListenAndServe(opts.Hostname+":"+opts.Port, nil)
}

func getOptions(args []interface{}) Options {
	options := args[0]
	switch options := options.(type) {
	case Options:
		return options
	default:
		return defaultOptions
	}
}

func wrap(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(h)
}

type Options struct {
	Hostname string
	Port     string
}
