package util

import "testing"

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	expect(t, config.App.PostsDir, "./_posts")
	expect(t, config.App.SiteDir, "./_site")
}

func BenchmarkDefaultConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DefaultConfig()
	}
}

func TestConfig_noFile(t *testing.T) {
	conf = config{}
	config := Config()
	def := DefaultConfig()

	expect(t, config.App.PostsDir, def.App.PostsDir)
	expect(t, config.App.SiteDir, def.App.SiteDir)
}

func TestConfig_fileNotExists(t *testing.T) {
	conf = config{}
	ConfigPath = "./nonexistent/app.config"
	config := Config()
	def := DefaultConfig()

	expect(t, config.App.PostsDir, def.App.PostsDir)
	expect(t, config.App.SiteDir, def.App.SiteDir)
}

func TestConfig_fileExists(t *testing.T) {
	conf = config{}
	ConfigPath = "../fixtures/config/app.config"
	config := Config()

	expect(t, config.App.PostsDir, "../fixtures/posts")
	expect(t, config.App.SiteDir, "../fixtures/site")
}

func TestConfig_badFile(t *testing.T) {
	conf = config{}
	ConfigPath = "../fixtures/config/bad.config"
	config := Config()
	def := DefaultConfig()

	expect(t, config.App.PostsDir, def.App.PostsDir)
	expect(t, config.App.SiteDir, def.App.SiteDir)
}

func BenchmarkConfig(b *testing.B) {
	ConfigPath = "../fixtures/app.config"
	for i := 0; i < b.N; i++ {
		_ = Config()
	}
}
