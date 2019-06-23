package lexicon

import (
	"fmt"

	"github.com/gellel/slice"
)

var (
	_ s = (*Slice)(nil)
)

// NewSlice instantiates a empty or populated Slice Lexicon pointer. Takes an argument of 0-N maps of slice.Slice.
func NewSlice(m ...map[string]*slice.Slice) *Slice {
	return (&Slice{
		lexicon: New()}).Mesh(m...)
}

type s interface {
	Add(key string, slice *slice.Slice) *Slice
	Del(key string) bool
	Each(f func(key string, slice *slice.Slice)) *Slice
	Empty() bool
	Fetch(key string) *slice.Slice
	Get(key string) (*slice.Slice, bool)
	Has(key string) bool
	Keys() *slice.String
	Len() int
	Map(f func(key string, slice *slice.Slice) *slice.Slice) *Slice
	Merge(slice ...*Slice) *Slice
	Mesh(m ...map[string]*slice.Slice) *Slice
	String() string
	Values() *slice.Slices
}

// Slice is a map-like object whose methods are used to perform traversal and mutation operations by key-value pair for slice.Slice pointers.
type Slice struct {
	lexicon *Lexicon
}

// Add method adds one slice.Slice pointer to the Slice Lexicon using the key reference and returns the modified Slice Lexicon.
func (pointer *Slice) Add(key string, slice *slice.Slice) *Slice {
	pointer.lexicon.Add(key, slice)
	return pointer
}

// Del method removes a entry from the Slice Lexicon if it exists. Returns a boolean to confirm if it succeeded.
func (pointer *Slice) Del(key string) bool {
	return pointer.lexicon.Del(key)
}

// Each method executes a provided function once for each Slice Lexicon element.
func (pointer *Slice) Each(f func(key string, slice *slice.Slice)) *Slice {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*slice.Slice))
	})
	return pointer
}

// Empty returns a boolean indicating whether the Slice Lexicon contains zero values.
func (pointer *Slice) Empty() bool {
	return pointer.lexicon.Empty()
}

// Fetch retrieves the slice.Slice pointer held by the argument key. Returns nil if key does not exist.
func (pointer *Slice) Fetch(key string) *slice.Slice {
	value, _ := pointer.Get(key)
	return value
}

// Get returns the slice.Slice pointer held at the argument key and a boolean indicating if it was successfully retrieved.
func (pointer *Slice) Get(key string) (*slice.Slice, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(*slice.Slice), ok
	}
	return nil, ok
}

// Has method checks that a given key exists in the Slice Lexicon.
func (pointer *Slice) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

// Keys method returns a slice.String of the Slice Lexicon's own property names, in the same order as we get with a normal loop.
func (pointer *Slice) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

// Len method returns the number of keys in the Slice Lexicon.
func (pointer *Slice) Len() int {
	return pointer.lexicon.Len()
}

// Map method executes a provided function once for each Slice Lexicon element and sets the returned value to the current key.
func (pointer *Slice) Map(f func(key string, slice *slice.Slice) *slice.Slice) *Slice {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*slice.Slice))
	})
	return pointer
}

// Merge merges N number of Slice Lexicon's to the Slice Lexicon and returns the modified Slice Lexicon.
func (pointer *Slice) Merge(s ...*Slice) *Slice {
	for _, s := range s {
		s.Each(func(key string, s *slice.Slice) {
			pointer.Add(key, s)
		})
	}
	return pointer
}

// Mesh merges N number of slice.Slice pointer maps to the Slice Lexicon.
func (pointer *Slice) Mesh(slice ...map[string]*slice.Slice) *Slice {
	for _, slice := range slice {
		for k, v := range slice {
			pointer.Add(k, v)
		}
	}
	return pointer
}

func (pointer *Slice) String() string {
	return fmt.Sprintf("%v", pointer.lexicon)
}

// Values method returns a slice.Slice pointer of the Slice Lexicon's own enumerable property values, in the same order as that provided by a for...in loop.
func (pointer *Slice) Values() *slice.Slices {
	slices := slice.NewSlices()
	pointer.Each(func(_ string, s *slice.Slice) {
		slices.Append(s)
	})
	return slices
}
