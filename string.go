package lex

import "sync"

// var _ Stringer = (&stringer{})(nil)

// Stringer is the interface that manages key value pairs
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
	EachValue(func(string)) Stringer
	Fetch(interface{}) string
	FetchSome(...interface{}) []string
	FetchSomeLength(...interface{}) ([]string, int)
	Get(interface{}) (interface{}, bool)
	GetLength(interface{}) (string, int, bool)
	Has(interface{}) bool
	Keys() []interface{}
	Len() int
	Map(func(interface{}, string) string) Stringer
	MapBreak(func(interface{}, string) (string, bool)) Stringer
	MapOK(func(interface{}, string) (string, bool)) Stringer
	Not(interface{}) bool
	NotSome(...interface{}) bool
	NotSomeLength(...interface{})
	Values() []string
}

type stringer struct {
	mu sync.Mutex
	l  *Lex
}

func (stringer *stringer) Add(k interface{}, v string) Stringer {
	return stringer.Mutate(func() { stringer.l.Add(k, v) })
}

func (stringer *stringer) AddLength(k interface{}, v string) int {
	var l int
	stringer.Mutate(func() { l = stringer.l.AddLength(k, v) })
	return l
}

func (stringer *stringer) AddOK(k interface{}, v string) bool {
	var ok bool
	stringer.Mutate(func() { ok = stringer.l.AddOK(k, v) })
	return ok
}

func (stringer *stringer) Del(k interface{}) Stringer {
	return stringer.Mutate(func() { stringer.l.Del(k) })
}

func (stringer *stringer) DelAll() Stringer {
	return stringer.Mutate(func() { stringer.l.DelAll() })
}

func (stringer *stringer) DelLength(k interface{}) int {
	var l int
	stringer.Mutate(func() { stringer.l.DelLength(k) })
	return l
}

func (stringer *stringer) DelSome(k ...interface{}) Stringer {
	return stringer.Mutate(func() { stringer.l.DelSome(k...) })
}

func (stringer *stringer) DelSomeLength(k ...interface{}) int {
	var l int
	stringer.Mutate(func() { l = stringer.l.DelSomeLength() })
	return l
}

func (stringer *stringer) DelOK(k interface{}) bool {
	var ok bool
	stringer.Mutate(func() { ok = stringer.l.DelOK(k) })
	return ok
}

func (stringer *stringer) Each(fn func(interface{}, string)) Stringer {
	stringer.Mutate(func() {
		stringer.l.Each(func(k, v interface{}) {
			fn(k, v.(string))
		})
	})
	return stringer
}

func (stringer *stringer) EachBreak(fn func(k interface{}, v string) bool) Stringer {
	stringer.Mutate(func() {
		stringer.l.EachBreak(func(k, v interface{}) bool {
			return fn(k, v.(string))
		})
	})
	return stringer
}

func (stringer *stringer) EachKey(fn func(k interface{})) Stringer {
	stringer.Mutate(func() {
		stringer.l.EachKey(fn)
	})
	return stringer
}

func (stringer *stringer) EachValue(fn func(v string)) Stringer {
	stringer.Mutate(func() {
		stringer.l.EachValue(func(v interface{}) {
			fn(v.(string))
		})
	})
	return stringer
}

func (stringer *stringer) Fetch(k interface{}) string {
	var s string
	stringer.Mutate(func() {
		var v = stringer.l.Fetch(k)
		if v != nil {
			s = v.(string)
		}
	})
	return s
}

func (stringer *stringer) FetchSome(k ...interface{}) []string {
	var s []string
	stringer.Mutate(func() {
		var x interface{}
		for _, x = range k {
			var v = stringer.l.Fetch(x)
			if v == nil {
				continue
			}
			s = append(s, v.(string))
		}
	})
	return s
}

func (stringer *stringer) FetchSomeLength(k ...interface{}) ([]string, int) {
	var s = stringer.FetchSome(k...)
	var l = len(s)
	return s, l
}

func (stringer *stringer) Get(k interface{}) (string, bool) {
	var s string
	var v, ok = stringer.l.Get(k)
	if v != nil {
		s = v.(string)
	}
	return s, ok
}

func (stringer *stringer) Lock() Stringer {
	stringer.mu.Lock()
	return stringer
}

func (stringer *stringer) Mutate(fn func()) Stringer {
	stringer.Lock()
	fn()
	stringer.Unlock()
	return stringer
}

func (stringer *stringer) Unlock() Stringer {
	stringer.mu.Unlock()
	return stringer
}
