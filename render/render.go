package render

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	Config = DefaultOptions()
	cache  *template.Template
	// Included helper functions for use when rendering html
	helperFuncs = template.FuncMap{
		"yield": func() (string, error) {
			return "", fmt.Errorf("yield called with no layout defined")
		},
		"current": func() (string, error) {
			return "", nil
		},
	}
)

type Options struct {
	TemplateDir        string
	TemplateExtensions []string
}

func init() {
	cache = compile(Config)
}

func DefaultOptions() *Options {
	return &Options{
		TemplateDir:        "./templates",
		TemplateExtensions: []string{"tmpl"},
	}
}

func HTML(rw http.ResponseWriter, name string, binding ...interface{}) (err error) {
	var b interface{}
	if len(binding) == 1 {
		b = binding[0]
	}

	addYield(rw, name, b)
	name = "admin/layout"

	buf, err := execute(rw, name, b)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("content-type", "text/html; charset=utf-8")
	io.Copy(rw, buf)
	return
}

func JSON(w io.Writer, data interface{}) (err error) {
	json, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.Write(json)
	return
}

func addYield(rw http.ResponseWriter, name string, binding interface{}) {
	funcs := template.FuncMap{
		"yield": func() (template.HTML, error) {
			buf, err := execute(rw, name, binding)
			// return safe html here since we are rendering our own template
			return template.HTML(buf.String()), err
		},
		"current": func() (string, error) {
			return name, nil
		},
	}
	cache.Funcs(funcs)
}

func compile(config *Options) (t *template.Template) {
	dir := config.TemplateDir
	exts := config.TemplateExtensions

	t = template.New(dir)
	t.Delims("{{", "}}")
	template.Must(t.Parse("b"))

	for r := range findAll(dir, exts) {
		buf, e := ioutil.ReadFile(dir + string(os.PathSeparator) + r)
		if e != nil {
			panic(e)
		}

		ext := getExt(r)
		name := r[0 : len(r)-len(ext)-1]
		tmpl := t.New(filepath.ToSlash(name))
		template.Must(tmpl.Funcs(helperFuncs).Parse(string(buf)))
	}

	return
}

func execute(rw http.ResponseWriter, name string, binding interface{}) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer([]byte{})
	return buf, cache.ExecuteTemplate(rw, name, binding)
}

func findAll(dir string, extensions []string) (c chan string) {
	c = make(chan string)
	go func() {
		filepath.Walk(dir, func(path string, file os.FileInfo, err error) (e error) {
			r, e := filepath.Rel(dir, path)
			if e != nil {
				return e
			}
			ext := getExt(r)

			if !file.IsDir() && file.Name()[0:1] != "." && isInSlice(ext, extensions) {
				c <- r
			}
			return
		})
		defer close(c)
	}()
	return
}

func getExt(st string) string {
	if strings.Index(st, ".") == -1 {
		return ""
	}
	return strings.Join(strings.Split(st, ".")[1:], ".")
}

func isInSlice(st string, sl []string) bool {
	for _, s := range sl {
		if s == st {
			return true
		}
	}
	return false
}
