package models

import (
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/slogsdon/b/util"
)

func TestParsePostId_noCategories(t *testing.T) {
	path := ParsePostId("2014-04-16-test-post-1.md")

	expect(t, path, "2014-04-16-test-post-1.md")
}

func TestParsePostId_emptyCategory(t *testing.T) {
	path := ParsePostId("_2014-04-16-test-post-1.md")

	expect(t, path, "2014-04-16-test-post-1.md")
}

func TestParsePostId_oneCategory(t *testing.T) {
	path := ParsePostId("test_2014-04-16-test-post-1.md")

	expect(t, path, "test/2014-04-16-test-post-1.md")
}

func TestParsePostId_multipleCategories(t *testing.T) {
	path := ParsePostId("test-category_2014-04-16-test-post-1.md")

	expect(t, path, "test/category/2014-04-16-test-post-1.md")
}

func TestSavePost_properResponsePostStruct(t *testing.T) {
	root := "../fixtures/posts"
	post := Post{}
	post.Filename = "2014-04-16-test-post-3.md"
	post.HeadMatter = HeadMatter{
		Title:      "Test Post 1",
		Date:       "2014-04-16 22:00:00",
		Categories: []string{"test"},
	}
	post.Raw = "This is a test post.\n\n## Test Posts\n\nPosting."

	err := SavePost(root, post)

	expect(t, err, nil)
}

func TestSavePost_properResponseUrlValues(t *testing.T) {
	root := "../fixtures/posts"
	form := url.Values{}
	form.Add("filename", "2014-04-16-test-post-3.md")
	form.Add("title", "Test Post 1")
	form.Add("date", "2014-04-16 22:00:00")
	form.Add("categories", "test")
	form.Add("raw", "This is a test post.\n\n## Test Posts\n\nPosting.")

	err := SavePost(root, form)

	expect(t, err, nil)
}

func TestSavePost_badTargetDir(t *testing.T) {
	root := "/etc/not-getting/here"
	form := url.Values{}
	form.Add("filename", "2014-04-16-not-going-to-happen.md")
	form.Add("raw", "testing.")

	err := SavePost(root, form)

	refute(t, err, nil)
}

func TestParsePostSlugAndType(t *testing.T) {
	file_chan := util.ReadDir("../fixtures/posts")
	files := []util.FileReading{}
	for f := range file_chan {
		files = append(files, f)
	}
	file := files[0]

	slug, ty := ParsePostSlugAndType(file.Info.Name())

	expect(t, slug, "test-post-1")
	expect(t, ty, "md")
}

func TestParsePostHeadMatter(t *testing.T) {
	file_chan := util.ReadDir("../fixtures/posts")
	files := []util.FileReading{}
	for f := range file_chan {
		files = append(files, f)
	}
	file := files[0]

	contents, err := ioutil.ReadFile(file.Filename)
	if err != nil {
		t.Errorf("Error reading file '%v'", file.Filename)
	}

	hm, _ := ParsePostHeadMatter(contents)

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
