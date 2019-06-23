package lexicon

import (
	"fmt"

	"github.com/gellel/slice"
)

var (
	_ i = (*Int)(nil)
)

// NewInt instantiates a empty or populated Int Lexicon pointer. Takes an argument of 0-N maps of integers.
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

// Int is a map-like object whose methods are used to perform traversal and mutation operations by key-value pair for integers.
type Int struct {
	lexicon *Lexicon
}

// Add method adds one string to the Int Lexicon using the key reference and returns the modified Int Lexicon.
func (pointer *Int) Add(key string, i int) *Int {
	pointer.lexicon.Add(key, i)
	return pointer
}

// Del method removes a entry from the Int Lexicon if it exists. Returns a boolean to confirm if it succeeded.
func (pointer *Int) Del(key string) bool {
	return pointer.lexicon.Del(key)
}

// Each method executes a provided function once for each Int Lexicon element.
func (pointer *Int) Each(f func(key string, i int)) *Int {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(int))
	})
	return pointer
}

// Empty returns a boolean indicating whether the Int Lexicon contains zero values.
func (pointer *Int) Empty() bool {
	return pointer.lexicon.Empty()
}

// Fetch retrieves the string held by the argument key. Returns nil string if key does not exist.
func (pointer *Int) Fetch(key string) int {
	i, _ := pointer.Get(key)
	return i
}

// Get returns the string held at the argument key and a boolean indicating if it was successfully retrieved.
func (pointer *Int) Get(key string) (int, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(int), ok
	}
	return 0, ok
}

// Has method checks that a given key exists in the Int Lexicon.
func (pointer *Int) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

// Keys method returns a slice.String of the Int Lexicon's own property names, in the same order as we get with a normal loop.
func (pointer *Int) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

// Len method returns the number of keys in the Int Lexicon.
func (pointer *Int) Len() int {
	return pointer.lexicon.Len()
}

// Map method executes a provided function once for each Int Lexicon element and sets the returned value to the current key.
func (pointer *Int) Map(f func(key string, i int) int) *Int {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(int))
	})
	return pointer
}

// Merge merges N number of Int Lexicons.
func (pointer *Int) Merge(i ...*Int) *Int {
	for _, i := range i {
		pointer.lexicon.Merge(i.lexicon)
	}
	return pointer
}

// Mesh merges a collection maps to the Int Lexicon.
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

// Values method returns a slice.Int pointer of the Int Lexicon's own enumerable property values, in the same order as that provided by a for...in loop.
func (pointer *Int) Values() *slice.Int {
	s := slice.NewInt()
	pointer.Each(func(_ string, value int) {
		s.Append(value)
	})
	return s
}
