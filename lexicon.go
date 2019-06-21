package lexicon

import (
	"github.com/gellel/slice"
)

var (
	_ lexicon = (*Lexicon)(nil)
)

func New() *Lexicon {
	return &Lexicon{}
}

func NewLexicon(m ...map[string]interface{}) *Lexicon {
	lexicon := Lexicon{}
	for _, m := range m {
		for k, v := range m {
			lexicon[k] = v
		}
	}
	return &lexicon
}

type lexicon interface {
	Add(key string, value interface{}) *Lexicon
	Del(key string) bool
	Each(f func(key string, value interface{})) *Lexicon
	Get(key string) (interface{}, bool)
	Has(key string) bool
	Keys() *slice.String
	Len() int
	Map(f func(key string, value interface{}) interface{}) *Lexicon
	Merge(lexicon *Lexicon) *Lexicon
	Values() *slice.Slice
}

type Lexicon map[string]interface{}

func (pointer *Lexicon) Add(key string, value interface{}) *Lexicon {
	(*pointer)[key] = value
	return pointer
}

func (pointer *Lexicon) Del(key string) bool {
	delete(*pointer, key)
	ok := (pointer.Has(key) == false)
	return ok
}

func (pointer *Lexicon) Each(f func(key string, value interface{})) *Lexicon {
	for key, value := range *pointer {
		f(key, value)
	}
	return pointer
}

func (pointer *Lexicon) Get(key string) (interface{}, bool) {
	value, ok := (*pointer)[key]
	return value, ok
}

func (pointer *Lexicon) Has(key string) bool {
	_, ok := pointer.Get(key)
	return ok
}

func (pointer *Lexicon) Keys() *slice.String {
	s := slice.NewString()
	for key := range *pointer {
		s.Append(key)
	}
	return s
}

func (pointer *Lexicon) Len() int {
	return len(*pointer)
}

func (pointer *Lexicon) Map(f func(key string, value interface{}) interface{}) *Lexicon {
	for key, value := range *pointer {
		pointer.Add(key, f(key, value))
	}
	return pointer
}

func (pointer *Lexicon) Merge(lexicon *Lexicon) *Lexicon {
	lexicon.Each(func(key string, value interface{}) {
		pointer.Add(key, value)
	})
	return pointer
}

func (pointer *Lexicon) Values() *slice.Slice {
	s := &slice.Slice{}
	for _, value := range *pointer {
		s.Append(value)
	}
	return s
}
