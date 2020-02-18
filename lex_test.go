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
		k = rand.Intn(10)
		v = rand.Intn(10)
	)
	l.Add(k, v)
	x, ok := (l)[k]
	if !ok {
		t.Fatalf("(&lex.Lex.Add(interface{}, interface{}) (_, bool)) != true")
	}
	ok = (x == v)
	if !ok {
		t.Fatalf("(&lex.Lex.Add(interface{}, interface{}) (interface{}, _)) != interface{}")
	}
}
