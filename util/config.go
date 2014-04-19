package util

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

var (
	ConfigPath = "./_private/app.config"
	conf       config
)

func Config() config {
	blank := config{}
	if conf == blank {
		readConfig()
	}
	return conf
}

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
}

type config struct {
	App app
}

type app struct {
	PostsDir string `toml:"posts_dir"`
	SiteDir  string `toml:"site_dir"`
}
