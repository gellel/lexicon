package lex_test

import (
	"sync"
	"testing"

	"github.com/gellel/lex"
)

var (
	inter lex.Inter
)

func TestInter(t *testing.T) {
	var (
		ok bool
	)
	inter = lex.NewInter()
	ok = inter != nil
	if !ok {
		t.Fatalf("(lex.NewInter() lex.Inter) == nil")
	}
	var (
		k = 1
		v int
		x int
	)
	ok = inter.AddOK(k, v)
	if !ok {
		t.Fatalf("(lex.Inter.AddOK(interface{}, int) bool) != true")
	}
	x, ok = inter.Get(k)
	if !ok {
		t.Fatalf("(lex.Inter.Get(interface{}) (int, bool) != (_, true)")
	}
	ok = v == x
	if !ok {
		t.Fatalf("(lex.Inter.Get(interface{}) (int, bool) != (int, _)")
	}
	ok = inter.DelOK(k)
	if !ok {
		t.Fatalf("(lex.Inter.DelOK(interface{}) (bool) != true")
	}
}

func TestInterConcurrency(t *testing.T) {
	var (
		ok   bool
		size = 20
		wg   sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < size/2; i++ {
			inter.AddLength(i, i)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := size / 2; i < size; i++ {
			inter.AddLength(i, i)
		}
	}()
	wg.Wait()
	ok = inter.Len() == size
	if !ok {
		t.Fatalf("(lex.Inter.Len() (int) != int")
	}
}
