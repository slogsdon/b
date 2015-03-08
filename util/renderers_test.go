package util

import "testing"

func TestMarkdown_byteSlice(t *testing.T) {
	generated := Markdown([]byte("## title"))
	expected := []byte("\u003ch2\u003etitle\u003c/h2\u003e\n")
	expect(t, string(generated), string(expected))
}

func TestMarkdown_string(t *testing.T) {
	generated := Markdown("## title")
	expected := "\u003ch2\u003etitle\u003c/h2\u003e\n"
	expect(t, string(generated), expected)
}

func TestMarkdown_unexpected(t *testing.T) {
	generated := Markdown(0123)
	expect(t, string(generated), "")
}
