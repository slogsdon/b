package util

import (
	"errors"
	"io/ioutil"
	"os"
)

// MakeDir creates the directory dir and all parent
// directories if they do not exist. It will return
// nil on success or if the directory already exists,
// otherwise, an error is returned.
func MakeDir(dir string) error {
	_, err := os.Stat(dir)

	if err != nil {
		return os.MkdirAll(dir, 0755)
	}

	return nil
}

// ReadDir reads dir contents recursively.
// It returns a slice of FileReading on completion.
func ReadDir(dir string) []FileReading {
	var m []FileReading
	return recursiveRead(dir, m)
}

// Write file writes contents to name. If necessary, it
// handles type casting from string to []byte.
// It returns an error if there are any issues writing
// the file contents.
func WriteFile(name string, contents interface{}) error {
	var c []byte

	switch contents.(type) {
	default:
		return errors.New("unexpected type when writing a file")
	case string:
		c = []byte(contents.(string))
	case []byte:
		c = []byte(contents.([]byte))
	}

	return ioutil.WriteFile(name, c, 0755)
}

type FileReading struct {
	// Filename is a application relative path of a file.
	Filename string
	Info     os.FileInfo
}

func recursiveRead(dir string, m []FileReading) []FileReading {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		file := dir + string(os.PathSeparator) + f.Name()
		if f.IsDir() {
			m = recursiveRead(file, m)
		} else {
			// Ignore hidden files
			if f.Name()[0:1] != "." {
				m = append(m, FileReading{Filename: file, Info: f})
			}
		}
	}
	return m
}
