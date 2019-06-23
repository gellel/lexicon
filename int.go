package lexicon

import (
	"fmt"

	"github.com/gellel/slice"
)

var (
	_ i = (*Int)(nil)
)

func NewInt(m ...map[string]int) *Int {
	return (&Int{lexicon: New()}).Mesh(m...)
}

type i interface {
	Add(key string, i int) *Int
	Del(key string) bool
	Each(f func(key string, i int)) *Int
	Empty() bool
	Fetch(key string) int
	Get(key string) (int, bool)
	Has(key string) bool
	Keys() *slice.String
	Len() int
	Map(f func(key string, i int) int) *Int
	Merge(m ...*Int) *Int
	Mesh(m ...map[string]int) *Int
	String() string
	Values() *slice.Int
}

type Int struct {
	lexicon *Lexicon
}

func (pointer *Int) Add(key string, i int) *Int {
	pointer.lexicon.Add(key, i)
	return pointer
}

func (pointer *Int) Del(key string) bool {
	return pointer.lexicon.Del(key)
}

func (pointer *Int) Each(f func(key string, i int)) *Int {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(int))
	})
	return pointer
}

func (pointer *Int) Empty() bool {
	return pointer.lexicon.Empty()
}

func (pointer *Int) Fetch(key string) int {
	i, _ := pointer.Get(key)
	return i
}

func (pointer *Int) Get(key string) (int, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(int), ok
	}
	return 0, ok
}

func (pointer *Int) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

func (pointer *Int) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

func (pointer *Int) Len() int {
	return pointer.lexicon.Len()
}

func (pointer *Int) Map(f func(key string, i int) int) *Int {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(int))
	})
	return pointer
}

func (pointer *Int) Merge(i ...*Int) *Int {
	for _, i := range i {
		pointer.lexicon.Merge(i.lexicon)
	}
	return pointer
}

func (pointer *Int) Mesh(m ...map[string]int) *Int {
	for _, m := range m {
		for k, v := range m {
			pointer.Add(k, v)
		}
	}
	return pointer
}

func (pointer *Int) String() string {
	return fmt.Sprintf("%v", pointer.lexicon)
}

func (pointer *Int) Values() *slice.Int {
	s := slice.NewInt()
	pointer.Each(func(_ string, value int) {
		s.Append(value)
	})
	return s
}
