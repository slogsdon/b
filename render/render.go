package render

import (
	"encoding/json"
	"html/template"
	"io"
)

var (
	Config = DefaultConfig()
	t      *template.Template
)

type RenderConfig struct {
	TemplateDir        string
	TemplateExtensions []string
}

func init() {
	pattern := Config.TemplateDir + "/*." + Config.TemplateExtensions[0]
	t, _ = template.ParseGlob(pattern)
}

func DefaultConfig() *RenderConfig {
	return &RenderConfig{
		TemplateDir:        "./templates",
		TemplateExtensions: []string{"tmpl"},
	}
}

func HTML(w io.Writer, name string, bindings ...interface{}) (err error) {
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
