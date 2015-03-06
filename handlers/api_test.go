package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/slogsdon/b/util"
)

func init() {
	util.ConfigPath = "../fixtures/config/app.config"
}

func TestApiIndex(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("GET", "/api", Api{}.Index)

	r, err := http.NewRequest("GET", "/api", nil)
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestApiPostsIndex(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("GET", "/api/posts", Api{}.Posts.Index)

	r, err := http.NewRequest("GET", "/api/posts", nil)
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestApiPostsCreate_goodRequestJson(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("POST", "/api/posts", Api{}.Posts.Create)
	buf := bytes.NewBufferString(`{"filename":"2014-04-16-test-post-3.md","raw":"This is a test post.\n\n## Test Posts\n\nPosting.", "head_matter":{"title": "Test Post 1","date": "2014-04-16 22:00:00","categories": ["test"]}}`)

	r, err := http.NewRequest("POST", "/api/posts", buf)
	r.Header.Set("Content-Type", "application/json; charset=utf-8")
	util.Mux.ServeHTTP(recorder, r)

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

func TestApiPostsCreate_goodRequestUrlEncoded(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("POST", "/api/posts", Api{}.Posts.Create)
	buf := bytes.NewBufferString("filename=2014-04-16-test-post-3.md&title=Test Post 1&date=2014-04-16 22:00:00&categories=test&raw=This is a test post.\n\n## Test Posts\n\nPosting.")

	r, err := http.NewRequest("POST", "/api/posts", buf)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	util.Mux.ServeHTTP(recorder, r)

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

func TestApiPostsCreate_badFilename(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("POST", "/api/posts", Api{}.Posts.Create)
	buf := bytes.NewBufferString("filename=&2014-04-16-test-post-3.md&raw=testing.")

	r, err := http.NewRequest("POST", "/api/posts", buf)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 500)

	_, err = os.Stat("../fixtures/posts/&2014-04-16-test-post-3.md")

	refute(t, err, nil)
}

func TestApiPostsCreate_badPastValues(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("POST", "/api/posts", Api{}.Posts.Create)

	r, err := http.NewRequest("POST", "/api/posts", nil)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 500)

	_, err = os.Stat("../fixtures/posts/&2014-04-16-test-post-3.md")

	refute(t, err, nil)
}

func TestApiPostsShow_fileExists(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("GET", "/api/posts/*id", Api{}.Posts.Show)

	r, err := http.NewRequest("GET", "/api/posts/2014-04-16-test-post-1.md", nil)
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
}

func TestApiPostsShow_fileNotExists(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("GET", "/api/posts/*id", Api{}.Posts.Show)

	r, err := http.NewRequest("GET", "/api/posts/2014-04-16-does-not-exists.md", nil)
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 404)
}

func TestApiRenderMarkdown_goodData(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("POST", "/api/render/markdown", Api{}.Render.Markdown)
	buf := bytes.NewBufferString("raw=## title")

	r, err := http.NewRequest("POST", "/api/render/markdown", buf)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 200)
	expect(t, recorder.Body.String(), `{"data":"\u003ch2\u003etitle\u003c/h2\u003e\n"}`)
}

func TestApiRenderMarkdown_badRequest(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("POST", "/api/render/markdown", Api{}.Render.Markdown)

	r, err := http.NewRequest("POST", "/api/render/markdown", nil)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 500)
}

func TestApiRenderMarkdown_noData(t *testing.T) {
	recorder := httptest.NewRecorder()
	util.Mux = httprouter.New()
	util.Mux.HandlerFunc("POST", "/api/render/markdown", Api{}.Render.Markdown)
	buf := bytes.NewBufferString("")

	r, err := http.NewRequest("POST", "/api/render/markdown", buf)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	util.Mux.ServeHTTP(recorder, r)

	expect(t, err, nil)
	expect(t, recorder.Code, 500)
}
