package util

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileReading struct {
	// Filename is a application relative path of a file.
	Filename string
	Info     os.FileInfo
}

// MakeDir creates the directory dir and all parent
// directories if they do not exist. It will return
// nil on success or if the directory already exists,
// otherwise, an error is returned.
func MakeDir(dir string) (err error) {
	_, err = os.Stat(dir)

	if err != nil {
		return os.MkdirAll(dir, 0755)
	}

	return
}

// ReadDir reads dir contents recursively.
// It returns a slice of FileReading on completion.
func ReadDir(dir string) (c chan FileReading) {
	c = make(chan FileReading)
	go func() {
		filepath.Walk(dir, func(path string, file os.FileInfo, _ error) (err error) {
			if !file.IsDir() && file.Name()[0:1] != "." {
				c <- FileReading{Filename: path, Info: file}
			}
			return
		})
		defer close(c)
	}()
	return
}

// Write file writes contents to name. If necessary, it
// handles type casting from string to []byte.
// It returns an error if there are any issues writing
// the file contents.
func WriteFile(name string, contents interface{}) error {
	var c []byte

	switch contents := contents.(type) {
	default:
		return errors.New("unexpected type when writing a file")
	case string:
		c = []byte(contents)
	case []byte:
		c = contents
	}

	return ioutil.WriteFile(name, c, 0755)
}
