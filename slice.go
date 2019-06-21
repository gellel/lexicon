package lexicon

import "github.com/gellel/slice"

var (
	_ s = (*Slice)(nil)
)

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
	Mesh(m ...map[string]*slice.Slice) *Slice
}

func NewSlice() *Slice {
	return &Slice{
		lexicon: New()}
}

type Slice struct {
	lexicon *Lexicon
}

func (pointer *Slice) Add(key string, slice *slice.Slice) *Slice {
	pointer.lexicon.Add(key, slice)
	return pointer
}

func (pointer *Slice) Del(key string) bool {
	return pointer.lexicon.Del(key)
}

func (pointer *Slice) Each(f func(key string, slice *slice.Slice)) *Slice {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*slice.Slice))
	})
	return pointer
}

func (pointer *Slice) Empty() bool {
	return pointer.lexicon.Empty()
}

func (pointer *Slice) Fetch(key string) *slice.Slice {
	value, _ := pointer.Get(key)
	return value
}

func (pointer *Slice) Get(key string) (*slice.Slice, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(*slice.Slice), ok
	}
	return nil, ok
}

func (pointer *Slice) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

func (pointer *Slice) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

func (pointer *Slice) Len() int {
	return pointer.lexicon.Len()
}

func (pointer *Slice) Map(f func(key string, slice *slice.Slice) *slice.Slice) *Slice {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*slice.Slice))
	})
	return pointer
}

func (pointer *Slice) Mesh(slice ...map[string]*slice.Slice) *Slice {
	for _, slice := range slice {
		for k, v := range slice {
			pointer.Add(k, v)
		}
	}
	return pointer
}
