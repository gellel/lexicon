package lexicon_test

import (
	"reflect"
	"testing"

	"github.com/gellel/lexicon"
)

var (
	i *lexicon.Int
)

func TestInt(t *testing.T) {
	i = lexicon.NewInt(map[string]int{"a": 1})

	ok := i != nil && reflect.ValueOf(i).Kind() == reflect.Ptr

	if ok != true {
		t.Fatalf("reflect.ValueOf(lexicon.Int) != reflect.Ptr")
	}
	value, ok := i.Get("a")

	if value != 1 || ok != true {
		t.Fatalf("string.Get(key string) did not return int or true")
	}
}
