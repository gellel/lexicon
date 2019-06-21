package lexicon_test

import (
	"reflect"
	"testing"

	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
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

func TestDel(t *testing.T) {

	if ok := l.Del("e"); ok != true {
		t.Fatalf("lexicon.Del(key string) did not return true")
	}
	if ok := l.Del("e"); ok != false {
		t.Fatalf("lexicon.Del(key string) did not return false")
	}
}

func TestEach(t *testing.T) {

	l.Each(func(key string, value interface{}) {
		if _, ok := (*l)[key]; ok != true {
			t.Fatalf("lexicon.Each(f func(key string, value interface{})) did not return an accurate key")
		}
		if v, _ := (*l)[key]; v != value {
			t.Fatalf("lexicon.Each(f func(key string, value interface{})) did not return the same interface at accessed key position")
		}
	})
}

func TestEmpty(t *testing.T) {

	if ok := (&lexicon.Lexicon{}).Empty(); ok != true {
		t.Fatalf("lexicon.Empty() did not return true for an empty lexicon")
	}
}

func TestFetch(t *testing.T) {

	if ok := l.Fetch("a") != nil; ok != true {
		t.Fatalf("lexicon.Fetch(key string) did not return an interface for a known key")
	}
	if ok := l.Fetch("NIL") == nil; ok != true {
		t.Fatalf("lexicon.Fetch(key string) did not return nil for a missing key")
	}
}

func TestGet(t *testing.T) {

	value, ok := l.Get("a")
	if ok != true {
		t.Fatalf("lexicon.Get(key string, value interface{}) did not return true for a known key")
	}
	if value.(int) != 1 {
		t.Fatalf("lexicon.Get(key string, value interface{}) did not return an interface")
	}
}

func TestHas(t *testing.T) {

	if ok := l.Has("b"); ok != true {
		t.Fatalf("lexicon.Has(key string) did not return true for a known key")
	}
}

func TestKeys(t *testing.T) {

	s := l.Keys()

	if ok := reflect.TypeOf(s).Kind() == reflect.TypeOf(slice.NewString()).Kind(); ok != true {
		t.Fatalf("lexicon.Keys() did not return a slice.String pointer")
	}

	s.Each(func(i int, key string) {

		if ok := l.Has(key); ok != true {
			t.Fatalf("lexicon.Keys() did not return a collection of valid lexicon key references")
		}
	})
}

func TestLen(t *testing.T) {

	if ok := l.Len() == len(*l); ok != true {
		t.Fatalf("lexicon.Len() returned an incorrect length")
	}
}

func TestMap(t *testing.T) {

	l.Map(func(key string, value interface{}) interface{} {

		value = 0

		return value
	})

	l.Each(func(_ string, value interface{}) {
		if ok := value.(int) == 0; ok != true {
			t.Fatalf("lexicon.Map(f (key string, value interface{}) interface{}) did not mutate the lexicon")
		}
	})
}

func TestMerge(t *testing.T) {

	l = &lexicon.Lexicon{}

	previousLength := l.Len()

	x := &lexicon.Lexicon{
		"1": 1,
		"2": 2}

	l.Merge(x)

	if ok := l.Len() != previousLength; ok != true {
		t.Fatalf("lexicon.Merge(lexicon *Lexicon) did not modify the length of the receiver lexicon")
	}
	if ok := l.Len() == x.Len(); ok != true {
		t.Fatalf("lexicon.Merge(lexicon *Lexicon) did not assign same number of values to receiver lexicon")
	}
	l.Each(func(key string, _ interface{}) {
		if ok := x.Has(key); ok != true {
			t.Fatalf("lexicon.Merge(lexicon *Lexicon) receiver lexicon does not have corresponding key to argument lexicon")
		}
	})
}

func TestMesh(t *testing.T) {

	l = &lexicon.Lexicon{}

	a := map[string]interface{}{"a": 1}
	b := map[string]interface{}{"b": 2}

	l.Mesh(a, b)

	for _, x := range []map[string]interface{}{a, b} {
		for k, v := range x {
			if ok := l.Has(k); ok != true {
				t.Fatalf("lexicon.Mesh(m ...map[string]interface{}) did not add key " + k + " to receiver Lexicon")
			}
			if ok := l.Fetch(k) == v; ok != true {
				t.Fatalf("lexicon.Mesh(m ...map[string]interface{}) did not add value " + string(v.(int)) + " to receiver Lexicon")
			}
		}
	}
}
