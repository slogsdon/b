package handlers

import (
	"bytes"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestApiIndex(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.Classic()
	m.Get("/api", Api{}.Index)

	r, err := http.NewRequest("GET", "/api", nil)
	m.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestApiPostsIndex(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/api/posts", Api{}.Posts.Index)

	r, err := http.NewRequest("GET", "/api/posts", nil)
	m.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestApiPostsCreate(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Post("/api/posts", Api{}.Posts.Create)
	buf := bytes.NewBufferString("filename=2014-04-16-test-post-3.md&raw=---\ntitle: Test Post 1\ndate: 2014-04-16 22:00:00\ncategories: [test]\n---\n\nThis is a test post.\n\n## Test Posts\n\nPosting.")

	r, err := http.NewRequest("POST", "/api/posts", buf)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	m.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 204)

	file, err := os.Stat("../fixtures/posts/test/2014-04-16-test-post-3.md")

	expect(t, err, nil)
	if file != nil {
		expect(t, file.Name(), "2014-04-16-test-post-3.md")
		os.Remove("../fixtures/posts/test/2014-04-16-test-post-3.md")
	}
	os.Remove("../fixtures/posts/test")
}
