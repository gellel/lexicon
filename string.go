package lexicon

import (
	"fmt"

	"github.com/gellel/slice"
)

var (
	_ str = (*String)(nil)
)

// NewString instantiates a empty or populated Lexicon pointer. Takes an argument of 0-N maps of strings.
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

// String is a map-like object whose methods are used to perform traversal and mutation operations by key-value pair for strings.
type String struct {
	lexicon *Lexicon
}

// Add method adds one string to the String Lexicon using the key reference and returns the modified String Lexicon.
func (pointer *String) Add(key, value string) *String {
	pointer.lexicon.Add(key, value)
	return pointer
}

// Del method removes a entry from the String Lexicon if it exists. Returns a boolean to confirm if it succeeded.
func (pointer *String) Del(key string) bool {
	return pointer.lexicon.Del(key)
}

// Each method executes a provided function once for each String Lexicon element.
func (pointer *String) Each(f func(key, value string)) *String {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(string))
	})
	return pointer
}

// Empty returns a boolean indicating whether the Lexicon contains zero values.
func (pointer *String) Empty() bool {
	return pointer.lexicon.Empty()
}

// Fetch retrieves the string held by the argument key. Returns nil string if key does not exist.
func (pointer *String) Fetch(key string) string {
	value, _ := pointer.Get(key)
	return value
}

// Get returns the string held at the argument key and a boolean indicating if it was successfully retrieved.
func (pointer *String) Get(key string) (string, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(string), ok
	}
	return fmt.Sprintf("%v", value), ok
}

// Has method checks that a given key exists in the String Lexicon.
func (pointer *String) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

// Keys method returns a slice.String of the String Lexicon's own property names, in the same order as we get with a normal loop.
func (pointer *String) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

// Len method returns the number of keys in the String Lexicon.
func (pointer *String) Len() int {
	return pointer.lexicon.Len()
}

// Map method executes a provided function once for each String Lexicon element and sets the returned value to the current key.
func (pointer *String) Map(f func(key, value string) string) *String {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(string))
	})
	return pointer
}

// Merge merges N number of String Lexicons.
func (pointer *String) Merge(s ...*String) *String {
	for _, s := range s {
		pointer.lexicon.Merge(s.lexicon)
	}
	return pointer
}

// Mesh merges a collection maps to the String Lexicon.
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

// Values method returns a slice.String pointer of the String Lexicon's own enumerable property values, in the same order as that provided by a for...in loop.
func (pointer *String) Values() *slice.String {
	s := slice.NewString()
	pointer.Each(func(_, value string) {
		s.Append(value)
	})
	return s
}
