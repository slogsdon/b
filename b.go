// B is a static-ish blog application. Can be run
// as a standalone application/server or be used
// to locally manage and deploy posts to a remote
// server.
package b

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/slogsdon/b/handlers"
	"github.com/slogsdon/b/util"
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
	util.Mux = httprouter.New()

	admin := handlers.Admin{}
	util.Mux.Handler("GET", "/admin", wrap(admin.Index))
	util.Mux.Handler("GET", "/admin/posts", wrap(admin.Posts.Index))
	util.Mux.Handler("GET", "/admin/posts/new", wrap(admin.Posts.New))
	util.Mux.Handler("GET", "/admin/posts/edit/*id", wrap(admin.Posts.Edit))

	api := handlers.Api{}
	util.Mux.Handler("GET", "/api", wrap(api.Index))
	util.Mux.Handler("GET", "/api/posts", wrap(api.Posts.Index))
	util.Mux.Handler("POST", "/api/posts", wrap(api.Posts.Create))
	util.Mux.Handler("GET", "/api/posts/*id", wrap(api.Posts.Show))
	util.Mux.Handler("GET", "/api/render/markdown", wrap(api.Render.Markdown))

	util.Mux.Handler("GET", "/", static)

	http.Handle("/", util.Mux)
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
