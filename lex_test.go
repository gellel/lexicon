package lex_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/gellel/lex"
)

var (
	l = (*&lex.Lex{})
)

func Test(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
}

func TestAdd(t *testing.T) {
	var (
		k  = rand.Intn(10)
		ok bool
		v  = rand.Intn(10)
		x  interface{}
	)
	l.Add(k, v)
	x, ok = (l)[k]
	if !ok {
		t.Fatalf("(&lex.Lex.Add(interface{}, interface{}) (_, bool)) != true")
	}
	ok = (x == v)
	if !ok {
		t.Fatalf("(&lex.Lex.Add(interface{}, interface{}) (interface{}, _)) != interface{}")
	}
}

func TestAddOK(t *testing.T) {
	var (
		max = 20
		min = 11
		ok  bool
	)
	var (
		k = (rand.Intn(max-min+1) + min)
		v = (rand.Intn(k*2-k+1) + k)
	)
	ok = l.AddOK(k, v)
	if !ok {
		t.Fatalf("(&lex.Lex.AddOK(interface{}, interface{}) (bool)) != true")
	}
}

func TestDel(t *testing.T) {
	var (
		k  interface{}
		ok bool
	)
	for k = range l {
		l.Del(k)
	}
	ok = len(l) == 0
	if !ok {
		t.Fatalf("len(&lex.Lex.Del(interface{})) != 0")
	}
}

func TestDelAll(t *testing.T) {
	var (
		length     int
		nextLength int
		ok         bool
	)
	for i := 0; i < (rand.Intn(10-5+1) + 5); i++ {
		l.Add(i, i)
	}
	length = len(l)
	l.DelAll()
	nextLength = len(l)
	ok = nextLength != length
	if !ok {
		t.Fatalf("len(&lex.Lex.DelSome()) != 0")
	}
	ok = nextLength == 0
	if !ok {
		t.Fatalf("len(&lex.Lex.DelSome()) != 0")
	}
}
