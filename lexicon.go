// Package lexicon provides interface for managing collections of abstract data in an Map-like structure.
package lexicon

import (
	"github.com/gellel/slice"
)

var (
	_ lexicon = (*Lexicon)(nil)
)

// New instantiates a new, empty Lexicon pointer.
func New() *Lexicon {
	return &Lexicon{}
}

// NewLexicon instantiates a empty or populated Lexicon pointer. Takes an argument of 0-N maps.
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
	Empty() bool
	Fetch(key string) interface{}
	Get(key string) (interface{}, bool)
	Has(key string) bool
	Keys() *slice.String
	Len() int
	Map(f func(key string, value interface{}) interface{}) *Lexicon
	Merge(lexicon *Lexicon) *Lexicon
	Values() *slice.Slice
}

// Lexicon is a map-like object whose methods are used to perform traversal and mutation operations by key-value pair.
type Lexicon map[string]interface{}

// Add method adds one element to the Lexicon using the key reference and returns the modified Lexicon.
func (pointer *Lexicon) Add(key string, value interface{}) *Lexicon {
	(*pointer)[key] = value
	return pointer
}

// Del method removes a entry from the Lexicon if it exists. Returns a boolean to confirm if it succeeded.
func (pointer *Lexicon) Del(key string) bool {
	ok := pointer.Has(key)
	if ok {
		delete(*pointer, key)
		ok = (pointer.Has(key) == false)
	}
	return ok
}

// Each method executes a provided function once for each Lexicon element.
func (pointer *Lexicon) Each(f func(key string, value interface{})) *Lexicon {
	for key, value := range *pointer {
		f(key, value)
	}
	return pointer
}

// Empty returns a boolean indicating whether the Lexicon contains zero values.
func (pointer *Lexicon) Empty() bool {
	return pointer.Len() == 0
}

// Fetch retrieves the interface held by the argument key. Returns nil if key does not exist.
func (pointer *Lexicon) Fetch(key string) interface{} {
	value, _ := (*pointer)[key]
	return value
}

// Get returns the value held at the argument key and a boolean indicating if it was successfully retrieved.
func (pointer *Lexicon) Get(key string) (interface{}, bool) {
	value, ok := (*pointer)[key]
	return value, ok
}

// Has method checks that a given key exists in the Lexicon.
func (pointer *Lexicon) Has(key string) bool {
	_, ok := pointer.Get(key)
	return ok
}

// Keys method returns a slice.String of the Lexicon's own property names, in the same order as we get with a normal loop.
func (pointer *Lexicon) Keys() *slice.String {
	s := slice.NewString()
	for key := range *pointer {
		s.Append(key)
	}
	return s
}

// Len method returns the number of keys in the Lexicon.
func (pointer *Lexicon) Len() int {
	return len(*pointer)
}

// Map method executes a provided function once for each Lexicon element and sets the returned value to the current key.
func (pointer *Lexicon) Map(f func(key string, value interface{}) interface{}) *Lexicon {
	for key, value := range *pointer {
		pointer.Add(key, f(key, value))
	}
	return pointer
}

// Merge merges two Lexicons.
func (pointer *Lexicon) Merge(lexicon *Lexicon) *Lexicon {
	lexicon.Each(func(key string, value interface{}) {
		pointer.Add(key, value)
	})
	return pointer
}

// Values method returns a slice.Slice pointer of the Lexicon's own enumerable property values, in the same order as that provided by a for...in loop.
func (pointer *Lexicon) Values() *slice.Slice {
	s := &slice.Slice{}
	for _, value := range *pointer {
		s.Append(value)
	}
	return s
}
