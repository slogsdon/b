package util

import (
	"testing"
)

func TestReadDir(t *testing.T) {
	files := ReadDir("../fixtures/posts")

	expect(t, len(files) > 0, true)
	expect(t, files[0].Info.Name(), "2014-04-16-test-post-1.md")
}
