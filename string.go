package lex

import "sync"

// var _ Stringer = (&stringer{})(nil)

// Stringer is map-like interface that holds a collection of strings.
//
// Stringer accepts any interface as a key but expects a string as its value.
// Stringer is safe for concurrent use by multiple goroutines without additional locking or coordination.
type Stringer interface {
	Add(interface{}, string) Stringer
	AddLength(interface{}, string) int
	AddOK(interface{}, string) bool
	Del(interface{}) Stringer
	DelAll() Stringer
	DelLength(interface{}) int
	DelSome(...interface{}) Stringer
	DelSomeLength(...interface{}) int
	DelOK(interface{}) bool
	Each(func(interface{}, string)) Stringer
	EachBreak(func(interface{}, string) bool) Stringer
	EachKey(func(interface{})) Stringer
	EachValue(func(interface{})) Stringer
	Fetch(interface{}) interface{}
	FetchSome(...interface{}) []interface{}
	Get(interface{}) (interface{}, bool)
	GetLength(interface{}) (interface{}, int, bool)
	Has(interface{}) bool
	Keys() []interface{}
	Len() int
	Map(func(interface{}, string) interface{}) Stringer
	MapBreak(func(interface{}, string) (interface{}, bool)) Stringer
	MapOK(func(interface{}, string) (interface{}, bool)) Stringer
	Not(interface{}) bool
	NotSome(...interface{}) bool
	Values() []interface{}
}

type stringer struct {
	mu sync.Mutex
	l  *Lex
}

func (stringer *stringer) Add(k interface{}, v string) Stringer  { stringer.l.Add(k, v); return stringer }
func (stringer *stringer) AddLength(k interface{}, v string) int { return stringer.l.AddLength(k, v) }
func (stringer *stringer) AddOK(k interface{}, v string) bool    { return stringer.l.AddOK(k, v) }
func (stringer *stringer) Del(k interface{}) Stringer            { stringer.l.Del(k); return stringer }
func (stringer *stringer) DelAll() Stringer                      { stringer.l.DelAll(); return stringer }
func (stringer *stringer) DelSome(k ...interface{}) Stringer     { stringer.DelSome(k...); return stringer }
func (stringer *stringer) DelSomeLength(k ...interface{}) int    { return stringer.l.DelSomeLength() }
