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

func BenchmarkMakeDir(b *testing.B) {
	name := "../fixtures/test-make-dir"
	for i := 0; i < b.N; i++ {
		_ = MakeDir(name)
	}
	os.Remove(name)
}

func TestReadDir(t *testing.T) {
	file_chan := ReadDir("../fixtures/posts")
	files := []FileReading{}
	for f := range file_chan {
		files = append(files, f)
	}

	expect(t, len(files) > 0, true)
	expect(t, files[0].Info.Name(), "2014-04-16-test-post-1.md")
}

func BenchmarkReadDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ReadDir("../fixtures/posts")
	}
}

func TestWriteFile_string(t *testing.T) {
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

func TestWriteFile_byteSlice(t *testing.T) {
	name := "../fixtures/test-make-file.txt"
	contents := []byte("test file contents")

	// Can write to file
	err := WriteFile(name, contents)
	expect(t, err, nil)

	// Can read contents
	readContents, err := ioutil.ReadFile(name)
	expect(t, string(readContents), string(contents))
	expect(t, err, nil)

	os.Remove(name)
}

func TestWriteFile_int(t *testing.T) {
	name := "../fixtures/test-make-file.txt"
	contents := 1

	// Can write to file
	err := WriteFile(name, contents)
	refute(t, err, nil)

	os.Remove(name)
}

func BenchmarkWriteFile(b *testing.B) {
	name := "../fixtures/test-make-file.txt"
	contents := "test file contents"
	for i := 0; i < b.N; i++ {
		_ = WriteFile(name, contents+string(i))
	}
	os.Remove(name)
}
