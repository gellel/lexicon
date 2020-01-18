package lexicon

// Lex is an implementation of a *map[interface{}]interface{}.
//
// Lex has methods to perform traversal and mutation operations.
//
// To extend a Lex construct a struct and a supporting interface that implements the Lex methods.
type Lex map[interface{}]interface{}

// Add adds a new key and element to the map and returns the modified map.
func (lex *Lex) Add(k interface{}, v interface{}) *Lex {
	(*lex)[k] = v
	return lex
}

// Del deletes the key and element from the map and returns the modified map.
func (lex *Lex) Del(k interface{}) *Lex {
	delete(*lex, k)
	return lex
}

// Each executes a provided function once for each map element.
func (lex *Lex) Each(fn func(interface{}, interface{})) *Lex {
	var (
		k, v interface{}
	)
	for k, v = range *lex {
		fn(k, v)
	}
	return lex
}

// EachBreak executes a provided function once for each
// element with an optional break when the function returns false.
func (lex *Lex) EachBreak(fn func(interface{}, interface{}) bool) *Lex {
	var (
		ok   = true
		k, v interface{}
	)
	for k, v = range *lex {
		ok = fn(k, v)
		if !ok {
			break
		}
	}
	return lex
}

// Get gets the element from the map at the key address.
// Returns a bool if the element was found using the key.
func (lex *Lex) Get(k interface{}) (interface{}, bool) {
	var v, ok = (*lex)[k]
	return v, ok
}

// Has checks that the map has a key of the corresponding element in the map.
func (lex *Lex) Has(k interface{}) bool {
	var _, ok = lex.Get(k)
	return ok
}
