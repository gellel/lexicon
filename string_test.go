package lexicon_test

import (
	"reflect"
	"testing"

	"github.com/gellel/lexicon"
)

var (
	str *lexicon.String
)

func TestString(t *testing.T) {
	str = lexicon.NewString(map[string]string{"a": "a"})

	ok := str != nil && reflect.ValueOf(str).Kind() == reflect.Ptr

	if ok != true {
		t.Fatalf("reflect.ValueOf(lexicon.String) != reflect.Ptr")
	}
	value, ok := str.Get("a")

	if value != "a" || ok != true {
		t.Fatalf("string.Get(key string) did not return string or true")
	}
}
