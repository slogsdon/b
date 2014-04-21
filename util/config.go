package util

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"strings"
)

var (
	// ConfigPath is an exported pacakge variable for declaring
	// the location of the application configuration file. This
	// can be set using the -c application flag when starting
	// the application via the command line.
	ConfigPath = "./_private/app.config"
	conf       config
)

// Config provides an exported facility for accessing the
// application configuration throughout the application.
// It returns the current value held by the unexported conf.
// If blank, it will read from the configuration file located
// at ConfigPath or the default configuration if ConfigPath
// does not exist.
func Config() config {
	blank := config{}
	if conf == blank {
		readConfig()
	}
	return conf
}

// DefaultConfig returns a set of default configuration values
// needed to run the application.
func DefaultConfig() config {
	return config{
		app{
			PostsDir: "./_posts",
			SiteDir:  "./_site",
		},
	}
}

func readConfig() {
	blob, err := ioutil.ReadFile(ConfigPath)
	if _, err = toml.Decode(string(blob), &conf); err != nil {
		fmt.Println(err)
		conf = DefaultConfig()
	}
	conf.App.PostsDir = strings.TrimRight(conf.App.PostsDir, "/")
	conf.App.SiteDir = strings.TrimRight(conf.App.SiteDir, "/")
}

type config struct {
	App app
}

type app struct {
	PostsDir string `toml:"posts_dir"`
	SiteDir  string `toml:"site_dir"`
}
