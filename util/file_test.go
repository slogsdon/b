package util

import (
	"reflect"
	"testing"
)

func TestReadDir(t *testing.T) {
	files := ReadDir("../fixtures/posts")

	expect(t, len(files) > 0, true)
	expect(t, files[0].Info.Name(), "2014-04-16-test-post-1.md")
}

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func refute(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
