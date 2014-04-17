package handlers

import (
	"github.com/go-martini/martini"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdminIndex(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.Classic()
	m.Get("/admin", Api{}.Index)

	r, err := http.NewRequest("GET", "/admin", nil)
	m.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestAdminPostsIndex(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.Classic()
	m.Get("/admin/posts", Api{}.Index)

	r, err := http.NewRequest("GET", "/admin/posts", nil)
	m.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestAdminPostsEdit(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.Classic()
	m.Get("/admin/posts/edit/2014-04-16-test-post-1.md", Api{}.Index)

	r, err := http.NewRequest("GET", "/admin/posts/edit/2014-04-16-test-post-1.md", nil)
	m.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}
