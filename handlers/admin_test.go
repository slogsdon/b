package handlers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/slogsdon/b/util"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdminIndex(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory: "../fixtures/templates",
		Layout:    "admin/layout",
	}))
	m.Get("/admin", Admin{}.Index)

	r, err := http.NewRequest("GET", "/admin", nil)
	m.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestAdminPostsIndex(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.Classic()
	util.ConfigPath = "../fixtures/config/app.config"
	m.Use(render.Renderer(render.Options{
		Directory: "../fixtures/templates",
		Layout:    "admin/layout",
	}))
	m.Get("/admin/posts", Admin{}.Posts.Index)

	r, err := http.NewRequest("GET", "/admin/posts", nil)
	m.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestAdminPostsNew(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.Classic()
	util.ConfigPath = "../fixtures/config/app.config"
	m.Use(render.Renderer(render.Options{
		Directory: "../fixtures/templates",
		Layout:    "admin/layout",
	}))
	m.Get("/admin/posts/new", Admin{}.Posts.New)

	r, err := http.NewRequest("GET", "/admin/posts/new", nil)
	m.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestAdminPostsEdit_fileExists(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.Classic()
	util.ConfigPath = "../fixtures/config/app.config"
	m.Use(render.Renderer(render.Options{
		Directory: "../fixtures/templates",
		Layout:    "admin/layout",
	}))
	m.Get("/admin/posts/:id/edit", Admin{}.Posts.Edit)

	r, err := http.NewRequest("GET", "/admin/posts/2014-04-16-test-post-1.md/edit", nil)
	m.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestAdminPostsEdit_fileNoExists(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.Classic()
	util.ConfigPath = "../fixtures/config/app.config"
	m.Use(render.Renderer(render.Options{
		Directory: "../fixtures/templates",
		Layout:    "admin/layout",
	}))
	m.Get("/admin/posts/:id/edit", Admin{}.Posts.Edit)

	r, err := http.NewRequest("GET", "/admin/posts/2014-04-16-non-existing-file.md/edit", nil)
	m.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 404)
}
