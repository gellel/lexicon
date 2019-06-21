package lexicon

import "github.com/gellel/slice"

var (
	_ lexicon = (*Lexicon)(nil)
)

type lexicon interface {
	Add(key string, value interface{}) *Lexicon
	Del(key string) bool
	Each(f func(key string, value interface{})) *Lexicon
	Get(key string) (interface{}, bool)
	Has(key string) bool
	Keys() *slice.String
	Len() int
	Map(f func(key string, value interface{}) interface{}) *Lexicon
	Merge(lexicon *Lexicon)
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
