package lex

import "sync"

var _ Stringer = (&stringer{})

// NewStringer returns a new Stringer interface.
func NewStringer() Stringer {
	return &stringer{sync.Mutex{}, &Lex{}}
}

// Stringer is the interface that manages key value pairs for strings.
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
	Get(interface{}) (string, bool)
	GetLength(interface{}) (string, int, bool)
	Has(interface{}) bool
	HasSome(...interface{}) bool
	Keys() []interface{}
	Len() int
	Map(func(interface{}, string) string) Stringer
	MapBreak(func(interface{}, string) (string, bool)) Stringer
	MapOK(func(interface{}, string) (string, bool)) Stringer
	Not(interface{}) bool
	NotSome(...interface{}) bool
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
	return stringer.Mutate(func() {
		stringer.l.Each(func(k, v interface{}) {
			fn(k, v.(string))
		})
	})
}

func (stringer *stringer) EachBreak(fn func(k interface{}, v string) bool) Stringer {
	return stringer.Mutate(func() {
		stringer.l.EachBreak(func(k, v interface{}) bool {
			return fn(k, v.(string))
		})
	})
}

func (stringer *stringer) EachKey(fn func(k interface{})) Stringer {
	return stringer.Mutate(func() { stringer.l.EachKey(fn) })
}

func (stringer *stringer) EachValue(fn func(v string)) Stringer {
	return stringer.Mutate(func() {
		stringer.l.EachValue(func(v interface{}) {
			fn(v.(string))
		})
	})
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
	var l = stringer.Len()
	return s, l
}

func (stringer *stringer) Get(k interface{}) (string, bool) {
	var s string
	var v interface{}
	var ok bool
	stringer.Mutate(func() {
		v, ok = stringer.l.Get(k)
		if v != nil {
			s = v.(string)
		}
	})
	return s, ok
}

func (stringer *stringer) GetLength(k interface{}) (string, int, bool) {
	var s, ok = stringer.Get(k)
	var l = stringer.Len()
	return s, l, ok
}

func (stringer *stringer) Has(k interface{}) bool {
	var ok bool
	stringer.Mutate(func() {
		ok = stringer.l.Has(k)
	})
	return ok
}

func (stringer *stringer) HasSome(k ...interface{}) bool {
	var ok bool
	stringer.Mutate(func() {
		ok = stringer.l.HasSome(k...)
	})
	return ok
}

func (stringer *stringer) Keys() []interface{} {
	var s []interface{}
	stringer.Mutate(func() {
		s = stringer.l.Keys()
	})
	return s
}

func (stringer *stringer) Len() int {
	var l int
	stringer.Mutate(func() {
		l = stringer.l.Len()
	})
	return l
}

func (stringer *stringer) Map(fn func(k interface{}, s string) string) Stringer {
	return stringer.Mutate(func() {
		stringer.l.Map(func(k, v interface{}) interface{} {
			return fn(k, v.(string))
		})
	})
}

func (stringer *stringer) MapBreak(fn func(interface{}, string) (string, bool)) Stringer {
	return stringer.Mutate(func() {
		stringer.l.MapBreak(func(k, v interface{}) (interface{}, bool) {
			return fn(k, v.(string))
		})
	})
}

func (stringer *stringer) MapOK(fn func(interface{}, string) (string, bool)) Stringer {
	return stringer.Mutate(func() {
		stringer.l.MapOK(func(k, v interface{}) (interface{}, bool) {
			return fn(k, v.(string))
		})
	})
}

func (stringer *stringer) Mutate(fn func()) Stringer {
	stringer.mu.Lock()
	fn()
	stringer.mu.Unlock()
	return stringer
}

func (stringer *stringer) Not(k interface{}) bool {
	var ok bool
	stringer.Mutate(func() {
		ok = stringer.l.Not(k)
	})
	return ok
}

func (stringer *stringer) NotSome(k ...interface{}) bool {
	var ok bool
	stringer.Mutate(func() {
		ok = stringer.l.NotSome(k...)
	})
	return ok
}

func (stringer *stringer) Values() []string {
	var s = []string{}
	stringer.Mutate(func() {
		stringer.EachValue(func(v string) {
			s = append(s, v)
		})
	})
	return s
}
