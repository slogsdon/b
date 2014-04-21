package util

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMakeDir(t *testing.T) {
	name := "../fixtures/test-make-dir"

	// No error when dir doesn't exist
	err := MakeDir(name)
	expect(t, err, nil)

	// No error when dir does exist
	err = MakeDir(name)
	expect(t, err, nil)

	os.Remove(name)
}

func TestReadDir(t *testing.T) {
	files := ReadDir("../fixtures/posts")

	expect(t, len(files) > 0, true)
	expect(t, files[0].Info.Name(), "2014-04-16-test-post-1.md")
}

func TestWriteFile(t *testing.T) {
	name := "../fixtures/test-make-file.txt"
	contents := "test file contents"

	// Can write to file
	err := WriteFile(name, contents)
	expect(t, err, nil)

	// Can read contents
	readContents, err := ioutil.ReadFile(name)
	expect(t, string(readContents), contents)
	expect(t, err, nil)

	os.Remove(name)
}
