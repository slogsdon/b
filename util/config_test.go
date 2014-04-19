package util

import "testing"

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	expect(t, config.App.PostsDir, "./_posts")
	expect(t, config.App.SiteDir, "./_site")
}

func TestConfig(t *testing.T) {
	ConfigPath = "../fixtures/app.config"
	config := Config()

	expect(t, config.App.PostsDir, "./posts")
	expect(t, config.App.SiteDir, "./site")
}
