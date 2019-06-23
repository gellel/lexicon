package lexicon

import (
	"fmt"

	"github.com/gellel/slice"
)

var (
	_ str = (*String)(nil)
)

func NewString(m ...map[string]string) *String {
	return (&String{
		lexicon: New()}).Mesh(m...)
}

type str interface {
	Add(key, value string) *String
	Del(key string) bool
	Each(f func(key, value string)) *String
	Empty() bool
	Fetch(key string) string
	Get(key string) (string, bool)
	Has(key string) bool
	Keys() *slice.String
	Len() int
	Map(f func(key, value string) string) *String
	Merge(s ...*String) *String
	Mesh(m ...map[string]string) *String
	String() string
	Values() *slice.String
}

type String struct {
	lexicon *Lexicon
}

func (pointer *String) Add(key, value string) *String {
	pointer.lexicon.Add(key, value)
	return pointer
}

func (pointer *String) Del(key string) bool {
	return pointer.lexicon.Del(key)
}

func (pointer *String) Each(f func(key, value string)) *String {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(string))
	})
	return pointer
}

func (pointer *String) Empty() bool {
	return pointer.lexicon.Empty()
}

func (pointer *String) Fetch(key string) string {
	value, _ := pointer.Get(key)
	return value
}

func (pointer *String) Get(key string) (string, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(string), ok
	}
	return fmt.Sprintf("%v", value), ok
}

func (pointer *String) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

func (pointer *String) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

func (pointer *String) Len() int {
	return pointer.lexicon.Len()
}

func (pointer *String) Map(f func(key, value string) string) *String {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(string))
	})
	return pointer
}

func (pointer *String) Merge(s ...*String) *String {
	for _, s := range s {
		pointer.lexicon.Merge(s.lexicon)
	}
	return pointer
}

func (pointer *String) Mesh(m ...map[string]string) *String {
	for _, m := range m {
		for k, v := range m {
			pointer.Add(k, v)
		}
	}
	return pointer
}

func (pointer *String) String() string {
	return fmt.Sprintf("%v", pointer.lexicon)
}

func (pointer *String) Values() *slice.String {
	s := slice.NewString()
	pointer.Each(func(_, value string) {
		s.Append(value)
	})
	return s
}
