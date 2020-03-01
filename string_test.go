package lex_test

import (
	"testing"

	"github.com/gellel/lex"
)

var (
	s lex.Stringer
)

func TestStringer(t *testing.T) {
	var (
		ok bool
	)
	s = lex.NewStringer()
	ok = s != nil
	if !ok {
		t.Fatalf("(lex.NewStringer() lex.Stringer) == nil")
	}
	var (
		k = 1
		v string
		x string
	)
	ok = s.AddOK(k, v)
	if !ok {
		t.Fatalf("(lex.Stringer.AddOK(interface{}, string) bool) != true")
	}
	x, ok = s.Get(k)
	if !ok {
		t.Fatalf("(lex.Stringer.Get(interface{}) (string, bool) != (_, true)")
	}
	ok = v == x
	if !ok {
		t.Fatalf("(lex.Stringer.Get(interface{}) (string, bool) != (string, _)")
	}
	ok = s.DelOK(k)
	if !ok {
		t.Fatalf("(lex.Stringer.DelOK(interface{}) (bool) != true")
	}
}
