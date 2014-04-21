package models

import (
	"github.com/slogsdon/b/util"
	"io/ioutil"
	"testing"
)

func TestParsePostSlugAndType(t *testing.T) {
	files := util.ReadDir("../fixtures/posts")
	file := files[0]

	slug, ty := ParsePostSlugAndType(file.Info.Name())

	expect(t, slug, "test-post-1")
	expect(t, ty, "md")
}

func TestPostParseContent(t *testing.T) {
	files := util.ReadDir("../fixtures/posts")
	file := files[0]

	contents, err := ioutil.ReadFile(file.Filename)
	if err != nil {
		t.Errorf("Error reading file '%v'", file.Filename)
	}

	_, ty := ParsePostSlugAndType(file.Info.Name())
	hm, _ := ParsePostContent(contents, ty)

	expect(t, hm.Title, "Test Post 1")
	expect(t, hm.Date, "2014-04-16 22:00:00")
	expect(t, len(hm.Categories) == 1, true)
	expect(t, hm.Categories[0], "test")
}

func TestGetAllPosts(t *testing.T) {
	posts := GetAllPosts("../fixtures/posts")

	expect(t, len(posts) > 0, true)

	post := posts[0]

	expect(t, post.Title, "Test Post 1")
	expect(t, post.Slug, "test-post-1")
}

func TestGetPost(t *testing.T) {
	post := GetPost("../fixtures/posts/2014-04-16-test-post-1.md", "../fixtures/posts")

	expect(t, post.Title, "Test Post 1")
	expect(t, post.Slug, "test-post-1")
}
