package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/slogsdon/b/util"
)

func init() {
	util.ConfigPath = "../fixtures/config/app.config"
}

func TestAdminIndex(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("GET", "/admin", Admin{}.Index)

	r, err := http.NewRequest("GET", "/admin", nil)
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestAdminPostsIndex(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("GET", "/admin/posts", Admin{}.Posts.Index)

	r, err := http.NewRequest("GET", "/admin/posts", nil)
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestAdminPostsNew(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("GET", "/admin/posts/new", Admin{}.Posts.New)

	r, err := http.NewRequest("GET", "/admin/posts/new", nil)
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestAdminPostsEdit_fileExists(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("GET", "/admin/posts/edit/*id", Admin{}.Posts.Edit)

	r, err := http.NewRequest("GET", "/admin/posts/edit/2014-04-16-test-post-1.md", nil)
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestAdminPostsEdit_fileNoExists(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("GET", "/admin/posts/edit/:id", Admin{}.Posts.Edit)

	r, err := http.NewRequest("GET", "/admin/posts/2014-04-16-non-existing-file.md/edit", nil)
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 404)
}
