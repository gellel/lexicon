package lex

import "sync"

var _ Booler = (&booler{})

// NewBooler returns a new Booler interface.
func NewBooler() Booler {
	return &booler{sync.Mutex{}, &Lex{}}
}

// Booler is the interface that manages key value pairs for bools.
//
// Booler accepts any interface as a key but expects a bool as its value.
// Booler is safe for concurrent use by multiple goroutines without additional locking or coordination.
type Booler interface {
	Add(interface{}, bool) Booler
	AddLength(interface{}, bool) int
	AddOK(interface{}, bool) bool
	Del(interface{}) Booler
	DelAll() Booler
	DelLength(interface{}) int
	DelSome(...interface{}) Booler
	DelSomeLength(...interface{}) int
	DelOK(interface{}) bool
	Each(func(interface{}, bool)) Booler
	EachBreak(func(interface{}, bool) bool) Booler
	EachKey(func(interface{})) Booler
	EachValue(func(bool)) Booler
	Fetch(interface{}) bool
	FetchSome(...interface{}) []bool
	FetchSomeLength(...interface{}) ([]bool, int)
	Get(interface{}) (bool, bool)
	GetLength(interface{}) (bool, int, bool)
	Has(interface{}) bool
	HasSome(...interface{}) bool
	Keys() []interface{}
	Len() int
	Map(func(interface{}, bool) bool) Booler
	MapBreak(func(interface{}, bool) (bool, bool)) Booler
	MapOK(func(interface{}, bool) (bool, bool)) Booler
	Not(interface{}) bool
	NotSome(...interface{}) bool
	Values() []bool
}

type booler struct {
	mu sync.Mutex
	l  *Lex
}

func (booler *booler) Add(k interface{}, v bool) Booler {
	return booler.Mutate(func() { booler.l.Add(k, v) })
}

func (booler *booler) AddLength(k interface{}, v bool) int {
	var l int
	booler.Mutate(func() { l = booler.l.AddLength(k, v) })
	return l
}

func (booler *booler) AddOK(k interface{}, v bool) bool {
	var ok bool
	booler.Mutate(func() { ok = booler.l.AddOK(k, v) })
	return ok
}

func (booler *booler) Del(k interface{}) Booler {
	return booler.Mutate(func() { booler.l.Del(k) })
}

func (booler *booler) DelAll() Booler {
	return booler.Mutate(func() { booler.l.DelAll() })
}

func (booler *booler) DelLength(k interface{}) int {
	var l int
	booler.Mutate(func() { booler.l.DelLength(k) })
	return l
}

func (booler *booler) DelSome(k ...interface{}) Booler {
	return booler.Mutate(func() { booler.l.DelSome(k...) })
}

func (booler *booler) DelSomeLength(k ...interface{}) int {
	var l int
	booler.Mutate(func() { l = booler.l.DelSomeLength() })
	return l
}

func (booler *booler) DelOK(k interface{}) bool {
	var ok bool
	booler.Mutate(func() { ok = booler.l.DelOK(k) })
	return ok
}

func (booler *booler) Each(fn func(interface{}, bool)) Booler {
	return booler.Mutate(func() {
		booler.l.Each(func(k, v interface{}) {
			fn(k, v.(bool))
		})
	})
}

func (booler *booler) EachBreak(fn func(k interface{}, v bool) bool) Booler {
	return booler.Mutate(func() {
		booler.l.EachBreak(func(k, v interface{}) bool {
			return fn(k, v.(bool))
		})
	})
}

func (booler *booler) EachKey(fn func(k interface{})) Booler {
	return booler.Mutate(func() { booler.l.EachKey(fn) })
}

func (booler *booler) EachValue(fn func(v bool)) Booler {
	return booler.Mutate(func() {
		booler.l.EachValue(func(v interface{}) {
			fn(v.(bool))
		})
	})
}

func (booler *booler) Fetch(k interface{}) bool {
	var s bool
	booler.Mutate(func() {
		var v = booler.l.Fetch(k)
		if v != nil {
			s = v.(bool)
		}
	})
	return s
}

func (booler *booler) FetchSome(k ...interface{}) []bool {
	var s []bool
	booler.Mutate(func() {
		var x interface{}
		for _, x = range k {
			var v = booler.l.Fetch(x)
			if v == nil {
				continue
			}
			s = append(s, v.(bool))
		}
	})
	return s
}

func (booler *booler) FetchSomeLength(k ...interface{}) ([]bool, int) {
	var s = booler.FetchSome(k...)
	var l = booler.Len()
	return s, l
}

func (booler *booler) Get(k interface{}) (bool, bool) {
	var s bool
	var v interface{}
	var ok bool
	booler.Mutate(func() {
		v, ok = booler.l.Get(k)
		if v != nil {
			s = v.(bool)
		}
	})
	return s, ok
}

func (booler *booler) GetLength(k interface{}) (bool, int, bool) {
	var s, ok = booler.Get(k)
	var l = booler.Len()
	return s, l, ok
}

func (booler *booler) Has(k interface{}) bool {
	var ok bool
	booler.Mutate(func() {
		ok = booler.l.Has(k)
	})
	return ok
}

func (booler *booler) HasSome(k ...interface{}) bool {
	var ok bool
	booler.Mutate(func() {
		ok = booler.l.HasSome(k...)
	})
	return ok
}

func (booler *booler) Keys() []interface{} {
	var s []interface{}
	booler.Mutate(func() {
		s = booler.l.Keys()
	})
	return s
}

func (booler *booler) Len() int {
	var l int
	booler.Mutate(func() {
		l = booler.l.Len()
	})
	return l
}

func (booler *booler) Map(fn func(k interface{}, s bool) bool) Booler {
	return booler.Mutate(func() {
		booler.l.Map(func(k, v interface{}) interface{} {
			return fn(k, v.(bool))
		})
	})
}

func (booler *booler) MapBreak(fn func(interface{}, bool) (bool, bool)) Booler {
	return booler.Mutate(func() {
		booler.l.MapBreak(func(k, v interface{}) (interface{}, bool) {
			return fn(k, v.(bool))
		})
	})
}

func (booler *booler) MapOK(fn func(interface{}, bool) (bool, bool)) Booler {
	return booler.Mutate(func() {
		booler.l.MapOK(func(k, v interface{}) (interface{}, bool) {
			return fn(k, v.(bool))
		})
	})
}

func (booler *booler) Mutate(fn func()) Booler {
	booler.mu.Lock()
	fn()
	booler.mu.Unlock()
	return booler
}

func (booler *booler) Not(k interface{}) bool {
	var ok bool
	booler.Mutate(func() {
		ok = booler.l.Not(k)
	})
	return ok
}

func (booler *booler) NotSome(k ...interface{}) bool {
	var ok bool
	booler.Mutate(func() {
		ok = booler.l.NotSome(k...)
	})
	return ok
}

func (booler *booler) Values() []bool {
	var s = []bool{}
	booler.Mutate(func() {
		booler.EachValue(func(v bool) {
			s = append(s, v)
		})
	})
	return s
}
