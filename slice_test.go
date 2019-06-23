package lexicon_test

import (
	"reflect"
	"testing"

	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

var (
	slc *lexicon.Slice
)

func TestSlice(t *testing.T) {
	slc = lexicon.NewSlice(map[string]*slice.Slice{"a": slice.New(1)})

	ok := slc != nil && reflect.ValueOf(slc).Kind() == reflect.Ptr

	if ok != true {
		t.Fatalf("reflect.ValueOf(lexicon.Slice) != reflect.Ptr")
	}
	s, ok := slc.Get("a")

	if s == nil || ok != true {
		t.Fatalf("slice.Get(key string) did not return slice.Slice pointer or true")
	}
}
