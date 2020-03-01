package lex_test

import (
	"fmt"
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
		t.Fatalf("len(&lex.Lex.DelAll()) != 0")
	}
	ok = nextLength == 0
	if !ok {
		t.Fatalf("len(&lex.Lex.DelAll()) != 0")
	}
}

func TestDelSome(t *testing.T) {
	var (
		size = (rand.Intn(10-5+1) + 5)
	)
	var (
		k      = make([]interface{}, size)
		length = len(l)
		ok     bool
	)
	for i := range k {
		k[i] = i
		l.Add(i, i)
	}
	l.DelSome(k...)
	ok = len(l) == length
	if !ok {
		t.Fatalf("len(&lex.Lex.DelSome()) != n")
	}
}

func TestDelOK(t *testing.T) {
	var (
		k  = rand.Intn(10)
		ok bool
	)
	l.Add(k, k)
	ok = l.DelOK(k)
	if !ok {
		t.Fatalf("(&lex.Lex.DelOK(interface{}) (bool)) != true")
	}
}

func TestEach(t *testing.T) {
	var (
		size = (rand.Intn(10-5+1) + 5)
	)
	var (
		k = make([]interface{}, size)
	)
	for i := 0; i < size; i++ {
		k = append(k, i)
	}
	l.Each(func(x, v interface{}) {
		var ok = k[x.(int)] == v
		if !ok {
			t.Fatalf("(&lex.Lex.Each(func(interface{}, interface{}))) != interface{}, interface{}")
		}
	})
}

func TestEachSome(t *testing.T) {
	var (
		size = (rand.Intn(10-5+1) + 5)
	)
	var (
		k = make([]interface{}, size)
	)
	for i := 0; i < size; i++ {
		k[i] = i
		l.Add(i, i)
	}
	var (
		v = l.FetchSome(k...)
	)
	var ok = len(v) == len(k)
	if !ok {
		t.Fatalf("len(&lex.Lex.EachSome(interface{}...) []interface{}) != n")
	}
}

func TestGet(t *testing.T) {
	var (
		k = (rand.Intn(10-5+1) + 5)
	)
	l.Add(k, k)
	var v, ok = l.Get(k)
	if !ok {
		t.Fatalf("(&lex.Lex.Get(interface{}) (_, bool)) != true")
	}
	ok = v == k
	if !ok {
		t.Fatalf("(&lex.Lex.Add(interface{}, interface{}) (interface{},)) != interface{}")
	}
}

func TestGetLength(t *testing.T) {
	var (
		size = len(l)
	)
	var (
		k = (rand.Intn(size*2-size+1) + size)
	)
	l.Add(k, k)
	var v, n, ok = l.GetLength(k)
	if !ok {
		t.Fatalf("(&lex.Lex.GetLength(interface{}) (_, _, bool)) != true")
	}
	ok = (n == (size + 1))
	if !ok {
		t.Fatalf("(&lex.Lex.Get(interface{}) (_, int, _)) != n + 1")
	}
	ok = v == k
	if !ok {
		t.Fatalf("(&lex.Lex.Get(interface{}) (interface{}, _, _)) != interface{}")
	}
}

func TestHas(t *testing.T) {
	var (
		size = len(l)
	)
	var (
		k = (rand.Intn(size*2-size+1) + size)
	)
	l.Add(k, k)
	var ok = l.Has(k)
	if !ok {
		t.Fatalf("(&lex.Lex.Has(interface{}) (bool)) != true")
	}
}

func TestKeys(t *testing.T) {
	var (
		a = []interface{}{}
		k interface{}
		v = []interface{}{}
	)
	for k = range l {
		v = append(v, k)
	}
	a = l.Keys()
	var ok = len(a) == len(v)
	if !ok {
		t.Fatalf("len(&lex.Lex.Keys() []interface{}) != n")
	}
}

func TestLen(t *testing.T) {
	var ok = len(l) == l.Len()
	if !ok {
		t.Fatalf("len(&lex.Lex.Len() int) != n")
	}
}

func TestMap(t *testing.T) {
	l.Map(func(_, v interface{}) interface{} {
		return fmt.Sprintf("%v", v)
	})
	for _, v := range l {
		switch v.(type) {
		case int:
			t.Fatalf("len(&lex.Lex.Map(func(interface{}, interface{}) interface{})) != string")
		}
	}
}

func TestNot(t *testing.T) {
	var (
		size = len(l)
	)
	var (
		k = size + 1
	)
	var ok = l.Not(k)
	if !ok {
		t.Fatalf("(&lex.Lex.Not(interface{}) (bool)) != true")
	}
}

func TestNotSome(t *testing.T) {
	var (
		size = len(l)
	)
	var (
		k = (rand.Intn(size*2-size+1) + size)
	)
	var v = []interface{}{}
	l = lex.Lex{}
	for i := k; i > size; i-- {
		v = append(v, i)
	}
	var ok = l.NotSome(v...)
	if !ok {
		t.Fatalf("(&lex.Lex.NotSome(...interface{}) (bool)) != true")
	}
}

func TestValues(t *testing.T) {
	var (
		k = (rand.Intn(10-5+1) + 5)
	)
	var (
		a = make([]interface{}, k)
		b []interface{}
	)
	for i := 0; i != k; i++ {
		a[i] = i
		l.Add(i, i)
	}
	b = l.Values()
	var ok = len(a) == len(b)
	if !ok {
		t.Fatalf("len(&lex.Lex.Values() []interface{}) != n")
	}
}
