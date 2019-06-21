package lexicon_test

import (
	"reflect"
	"testing"

	"github.com/gellel/lexicon"
)

var (
	l *lexicon.Lexicon
)

func Test(t *testing.T) {
	l = lexicon.New()

	ok := l != nil && reflect.ValueOf(l).Kind() == reflect.Ptr

	if ok != true {
		t.Fatalf("reflect.ValueOf(lexicon.Lexicon) != reflect.Ptr")
	}

	x := map[string]interface{}{
		"a": 1,
		"b": 2}
	y := map[string]interface{}{
		"c": 3,
		"d": 4}

	l = lexicon.NewLexicon(x, y)

	for _, k := range []string{"a", "b", "c", "d"} {
		if _, ok := (*l)[k]; ok != true {
			t.Fatalf("lexicon.NewLexicon(m ...map[string]interface{}) did not assign range of maps")
		}
	}
}

func TestAdd(t *testing.T) {

	l.Add("e", 5)

	if _, ok := (*l)["e"]; ok != true {
		t.Fatalf("lexicon.Add(key string, value interface{}) did not add new key of \"e\"")
	}
}
