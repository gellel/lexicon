package lex

import "sync"

var _ Byter = (&byter{})

// NewByter returns a new Byter interface.
func NewByter() Byter {
	return &byter{sync.Mutex{}, &Lex{}}
}

// Byter is the interface that manages key value pairs for bytes.
//
// Byter accepts any interface as a key but expects a byte as its value.
// Byter is safe for concurrent use by multiple goroutines without additional locking or coordination.
type Byter interface {
	Add(interface{}, byte) Byter
	AddLength(interface{}, byte) int
	AddOK(interface{}, byte) bool
	Del(interface{}) Byter
	DelAll() Byter
	DelLength(interface{}) int
	DelSome(...interface{}) Byter
	DelSomeLength(...interface{}) int
	DelOK(interface{}) bool
	Each(func(interface{}, byte)) Byter
	EachBreak(func(interface{}, byte) bool) Byter
	EachKey(func(interface{})) Byter
	EachValue(func(byte)) Byter
	Fetch(interface{}) byte
	FetchSome(...interface{}) []byte
	FetchSomeLength(...interface{}) ([]byte, int)
	Get(interface{}) (byte, bool)
	GetLength(interface{}) (byte, int, bool)
	Has(interface{}) bool
	HasSome(...interface{}) bool
	Keys() []interface{}
	Len() int
	Map(func(interface{}, byte) byte) Byter
	MapBreak(func(interface{}, byte) (byte, bool)) Byter
	MapOK(func(interface{}, byte) (byte, bool)) Byter
	Not(interface{}) bool
	NotSome(...interface{}) bool
	Values() []byte
}

type byter struct {
	mu sync.Mutex
	l  *Lex
}

func (byter *byter) Add(k interface{}, v byte) Byter {
	return byter.Mutate(func() { byter.l.Add(k, v) })
}

func (byter *byter) AddLength(k interface{}, v byte) int {
	var l int
	byter.Mutate(func() { l = byter.l.AddLength(k, v) })
	return l
}

func (byter *byter) AddOK(k interface{}, v byte) bool {
	var ok bool
	byter.Mutate(func() { ok = byter.l.AddOK(k, v) })
	return ok
}

func (byter *byter) Del(k interface{}) Byter {
	return byter.Mutate(func() { byter.l.Del(k) })
}

func (byter *byter) DelAll() Byter {
	return byter.Mutate(func() { byter.l.DelAll() })
}

func (byter *byter) DelLength(k interface{}) int {
	var l int
	byter.Mutate(func() { byter.l.DelLength(k) })
	return l
}

func (byter *byter) DelSome(k ...interface{}) Byter {
	return byter.Mutate(func() { byter.l.DelSome(k...) })
}

func (byter *byter) DelSomeLength(k ...interface{}) int {
	var l int
	byter.Mutate(func() { l = byter.l.DelSomeLength() })
	return l
}

func (byter *byter) DelOK(k interface{}) bool {
	var ok bool
	byter.Mutate(func() { ok = byter.l.DelOK(k) })
	return ok
}

func (byter *byter) Each(fn func(interface{}, byte)) Byter {
	return byter.Mutate(func() {
		byter.l.Each(func(k, v interface{}) {
			fn(k, v.(byte))
		})
	})
}

func (byter *byter) EachBreak(fn func(k interface{}, v byte) bool) Byter {
	return byter.Mutate(func() {
		byter.l.EachBreak(func(k, v interface{}) bool {
			return fn(k, v.(byte))
		})
	})
}

func (byter *byter) EachKey(fn func(k interface{})) Byter {
	return byter.Mutate(func() { byter.l.EachKey(fn) })
}

func (byter *byter) EachValue(fn func(v byte)) Byter {
	return byter.Mutate(func() {
		byter.l.EachValue(func(v interface{}) {
			fn(v.(byte))
		})
	})
}

func (byter *byter) Fetch(k interface{}) byte {
	var s byte
	byter.Mutate(func() {
		var v = byter.l.Fetch(k)
		if v != nil {
			s = v.(byte)
		}
	})
	return s
}

func (byter *byter) FetchSome(k ...interface{}) []byte {
	var s []byte
	byter.Mutate(func() {
		var x interface{}
		for _, x = range k {
			var v = byter.l.Fetch(x)
			if v == nil {
				continue
			}
			s = append(s, v.(byte))
		}
	})
	return s
}

func (byter *byter) FetchSomeLength(k ...interface{}) ([]byte, int) {
	var s = byter.FetchSome(k...)
	var l = byter.Len()
	return s, l
}

func (byter *byter) Get(k interface{}) (byte, bool) {
	var s byte
	var v interface{}
	var ok bool
	byter.Mutate(func() {
		v, ok = byter.l.Get(k)
		if v != nil {
			s = v.(byte)
		}
	})
	return s, ok
}

func (byter *byter) GetLength(k interface{}) (byte, int, bool) {
	var s, ok = byter.Get(k)
	var l = byter.Len()
	return s, l, ok
}

func (byter *byter) Has(k interface{}) bool {
	var ok bool
	byter.Mutate(func() {
		ok = byter.l.Has(k)
	})
	return ok
}

func (byter *byter) HasSome(k ...interface{}) bool {
	var ok bool
	byter.Mutate(func() {
		ok = byter.l.HasSome(k...)
	})
	return ok
}

func (byter *byter) Keys() []interface{} {
	var s []interface{}
	byter.Mutate(func() {
		s = byter.l.Keys()
	})
	return s
}

func (byter *byter) Len() int {
	var l int
	byter.Mutate(func() {
		l = byter.l.Len()
	})
	return l
}

func (byter *byter) Map(fn func(k interface{}, s byte) byte) Byter {
	return byter.Mutate(func() {
		byter.l.Map(func(k, v interface{}) interface{} {
			return fn(k, v.(byte))
		})
	})
}

func (byter *byter) MapBreak(fn func(interface{}, byte) (byte, bool)) Byter {
	return byter.Mutate(func() {
		byter.l.MapBreak(func(k, v interface{}) (interface{}, bool) {
			return fn(k, v.(byte))
		})
	})
}

func (byter *byter) MapOK(fn func(interface{}, byte) (byte, bool)) Byter {
	return byter.Mutate(func() {
		byter.l.MapOK(func(k, v interface{}) (interface{}, bool) {
			return fn(k, v.(byte))
		})
	})
}

func (byter *byter) Mutate(fn func()) Byter {
	byter.mu.Lock()
	fn()
	byter.mu.Unlock()
	return byter
}

func (byter *byter) Not(k interface{}) bool {
	var ok bool
	byter.Mutate(func() {
		ok = byter.l.Not(k)
	})
	return ok
}

func (byter *byter) NotSome(k ...interface{}) bool {
	var ok bool
	byter.Mutate(func() {
		ok = byter.l.NotSome(k...)
	})
	return ok
}

func (byter *byter) Values() []byte {
	var s = []byte{}
	byter.Mutate(func() {
		byter.EachValue(func(v byte) {
			s = append(s, v)
		})
	})
	return s
}
