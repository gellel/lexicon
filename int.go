package lexicon

import "github.com/gellel/slice"

func NewInt(m ...map[string]int) *Int {
	return (&Int{lexicon: New()})
}

type i interface {
	Add(key string, i int) *Int
	Del(key string) *Int
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
