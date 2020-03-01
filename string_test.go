package lex_test

import (
	"fmt"
	"sync"
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

func TestStringerConcurrency(t *testing.T) {
	var (
		ok   bool
		size = 20
		wg   sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < size/2; i++ {
			s.AddLength(i, fmt.Sprintf("%d", i))
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := size / 2; i < size; i++ {
			s.AddLength(i, fmt.Sprintf("%d", i))
		}
	}()
	wg.Wait()
	ok = s.Len() == size
	if !ok {
		t.Fatalf("(lex.Stringer.Len() (int) != int")
	}
}
