package util

import (
	"io/ioutil"
	"os"
)

func ReadDir(dir string) []FileReading {
	var m []FileReading
	return read(dir, m)
}

type FileReading struct {
	Filename string
	Info     os.FileInfo
}

func read(dir string, m []FileReading) []FileReading {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		file := dir + string(os.PathSeparator) + f.Name()
		if f.IsDir() {
			m = read(file, m)
		} else {
			// Ignore hidden files
			if f.Name()[0:1] != "." {
				m = append(m, FileReading{Filename: file, Info: f})
			}
		}
	}
	return m
}
