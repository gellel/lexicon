package lex

import "sync"

var _ Inter = (&inter{})

// NewInter returns a new Inter interface.
func NewInter() Inter {
	return &inter{sync.Mutex{}, &Lex{}}
}

// Inter is the interface that manages key value pairs
//
// Inter accepts any interface as a key but expects a int as its value.
// Inter is safe for concurrent use by multiple goroutines without additional locking or coordination.
type Inter interface {
	Add(interface{}, int) Inter
	AddLength(interface{}, int) int
	AddOK(interface{}, int) bool
	Del(interface{}) Inter
	DelAll() Inter
	DelLength(interface{}) int
	DelSome(...interface{}) Inter
	DelSomeLength(...interface{}) int
	DelOK(interface{}) bool
	Each(func(interface{}, int)) Inter
	EachBreak(func(interface{}, int) bool) Inter
	EachKey(func(interface{})) Inter
	EachValue(func(int)) Inter
	Fetch(interface{}) int
	FetchSome(...interface{}) []int
	FetchSomeLength(...interface{}) ([]int, int)
	Get(interface{}) (int, bool)
	GetLength(interface{}) (int, int, bool)
	Has(interface{}) bool
	HasSome(...interface{}) bool
	Keys() []interface{}
	Len() int
	Map(func(interface{}, int) int) Inter
	MapBreak(func(interface{}, int) (int, bool)) Inter
	MapOK(func(interface{}, int) (int, bool)) Inter
	Not(interface{}) bool
	NotSome(...interface{}) bool
	Values() []int
}

type inter struct {
	mu sync.Mutex
	l  *Lex
}

func (inter *inter) Add(k interface{}, v int) Inter {
	return inter.Mutate(func() { inter.l.Add(k, v) })
}

func (inter *inter) AddLength(k interface{}, v int) int {
	var l int
	inter.Mutate(func() { l = inter.l.AddLength(k, v) })
	return l
}

func (inter *inter) AddOK(k interface{}, v int) bool {
	var ok bool
	inter.Mutate(func() { ok = inter.l.AddOK(k, v) })
	return ok
}

func (inter *inter) Del(k interface{}) Inter {
	return inter.Mutate(func() { inter.l.Del(k) })
}

func (inter *inter) DelAll() Inter {
	return inter.Mutate(func() { inter.l.DelAll() })
}

func (inter *inter) DelLength(k interface{}) int {
	var l int
	inter.Mutate(func() { inter.l.DelLength(k) })
	return l
}

func (inter *inter) DelSome(k ...interface{}) Inter {
	return inter.Mutate(func() { inter.l.DelSome(k...) })
}

func (inter *inter) DelSomeLength(k ...interface{}) int {
	var l int
	inter.Mutate(func() { l = inter.l.DelSomeLength() })
	return l
}

func (inter *inter) DelOK(k interface{}) bool {
	var ok bool
	inter.Mutate(func() { ok = inter.l.DelOK(k) })
	return ok
}

func (inter *inter) Each(fn func(interface{}, int)) Inter {
	return inter.Mutate(func() {
		inter.l.Each(func(k, v interface{}) {
			fn(k, v.(int))
		})
	})
}

func (inter *inter) EachBreak(fn func(k interface{}, v int) bool) Inter {
	return inter.Mutate(func() {
		inter.l.EachBreak(func(k, v interface{}) bool {
			return fn(k, v.(int))
		})
	})
}

func (inter *inter) EachKey(fn func(k interface{})) Inter {
	return inter.Mutate(func() { inter.l.EachKey(fn) })
}

func (inter *inter) EachValue(fn func(v int)) Inter {
	return inter.Mutate(func() {
		inter.l.EachValue(func(v interface{}) {
			fn(v.(int))
		})
	})
}

func (inter *inter) Fetch(k interface{}) int {
	var s int
	inter.Mutate(func() {
		var v = inter.l.Fetch(k)
		if v != nil {
			s = v.(int)
		}
	})
	return s
}

func (inter *inter) FetchSome(k ...interface{}) []int {
	var s []int
	inter.Mutate(func() {
		var x interface{}
		for _, x = range k {
			var v = inter.l.Fetch(x)
			if v == nil {
				continue
			}
			s = append(s, v.(int))
		}
	})
	return s
}

func (inter *inter) FetchSomeLength(k ...interface{}) ([]int, int) {
	var s = inter.FetchSome(k...)
	var l = inter.Len()
	return s, l
}

func (inter *inter) Get(k interface{}) (int, bool) {
	var s int
	var v interface{}
	var ok bool
	inter.Mutate(func() {
		v, ok = inter.l.Get(k)
		if v != nil {
			s = v.(int)
		}
	})
	return s, ok
}

func (inter *inter) GetLength(k interface{}) (int, int, bool) {
	var s, ok = inter.Get(k)
	var l = inter.Len()
	return s, l, ok
}

func (inter *inter) Has(k interface{}) bool {
	var ok bool
	inter.Mutate(func() {
		ok = inter.l.Has(k)
	})
	return ok
}

func (inter *inter) HasSome(k ...interface{}) bool {
	var ok bool
	inter.Mutate(func() {
		ok = inter.l.HasSome(k...)
	})
	return ok
}

func (inter *inter) Keys() []interface{} {
	var s []interface{}
	inter.Mutate(func() {
		s = inter.l.Keys()
	})
	return s
}

func (inter *inter) Len() int {
	var l int
	inter.Mutate(func() {
		l = inter.l.Len()
	})
	return l
}

func (inter *inter) Map(fn func(k interface{}, s int) int) Inter {
	return inter.Mutate(func() {
		inter.l.Map(func(k, v interface{}) interface{} {
			return fn(k, v.(int))
		})
	})
}

func (inter *inter) MapBreak(fn func(interface{}, int) (int, bool)) Inter {
	return inter.Mutate(func() {
		inter.l.MapBreak(func(k, v interface{}) (interface{}, bool) {
			return fn(k, v.(int))
		})
	})
}

func (inter *inter) MapOK(fn func(interface{}, int) (int, bool)) Inter {
	return inter.Mutate(func() {
		inter.l.MapOK(func(k, v interface{}) (interface{}, bool) {
			return fn(k, v.(int))
		})
	})
}

func (inter *inter) Mutate(fn func()) Inter {
	inter.mu.Lock()
	fn()
	inter.mu.Unlock()
	return inter
}

func (inter *inter) Not(k interface{}) bool {
	var ok bool
	inter.Mutate(func() {
		ok = inter.l.Not(k)
	})
	return ok
}

func (inter *inter) NotSome(k ...interface{}) bool {
	var ok bool
	inter.Mutate(func() {
		ok = inter.l.NotSome(k...)
	})
	return ok
}

func (inter *inter) Values() []int {
	var s = []int{}
	inter.Mutate(func() {
		inter.EachValue(func(v int) {
			s = append(s, v)
		})
	})
	return s
}
