package models

import (
	"github.com/slogsdon/b/util"
	"io/ioutil"
	"testing"
)

func TestSavePost_properResponse(t *testing.T) {
	root := "../fixtures/posts"
	form := map[string][]string{
		"filename": {
			"2014-04-16-test-post-3.md",
		},
		"raw": {
			"---\ntitle: Test Post 1\ndate: 2014-04-16 22:00:00\ncategories: [test]\n---\n\nThis is a test post.\n\n## Test Posts\n\nPosting.",
		},
	}

	err := SavePost(root, form)

	expect(t, err, nil)
}

func TestSavePost_badTargetDir(t *testing.T) {
	root := "/etc/not-getting/here"
	form := map[string][]string{
		"filename": {
			"2014-04-16-not-going-to-happen.md",
		},
		"raw": {
			"testing.",
		},
	}

	err := SavePost(root, form)

	refute(t, err, nil)
}

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
