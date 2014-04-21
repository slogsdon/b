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

func TestConfig(t *testing.T) {
	ConfigPath = "../fixtures/app.config"
	config := Config()

	expect(t, config.App.PostsDir, "./posts")
	expect(t, config.App.SiteDir, "./site")
}

func BenchmarkConfig(b *testing.B) {
	ConfigPath = "../fixtures/app.config"
	for i := 0; i < b.N; i++ {
		_ = Config()
	}
}
