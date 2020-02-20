package lex

var _ lexer = (&Lex{})

type lexer interface {
	Add(interface{}, interface{}) *Lex
	AddLength(interface{}, interface{}) int
	AddOK(interface{}, interface{}) bool
	Del(interface{}) *Lex
	DelAll() *Lex
	DelLength(interface{}) int
	DelSome(...interface{}) *Lex
	DelSomeLength(...interface{}) int
	DelOK(interface{}) bool
	Each(func(interface{}, interface{})) *Lex
	EachBreak(func(interface{}, interface{}) bool) *Lex
	EachKey(func(interface{})) *Lex
	EachValue(func(interface{})) *Lex
	Fetch(interface{}) interface{}
	FetchSome(...interface{}) []interface{}
	Get(interface{}) (interface{}, bool)
	GetLength(interface{}) (interface{}, int, bool)
	Has(interface{}) bool
	Keys() []interface{}
	Len() int
	Map(func(interface{}, interface{}) interface{}) *Lex
	MapBreak(func(interface{}, interface{}) (interface{}, bool)) *Lex
	MapOK(func(interface{}, interface{}) (interface{}, bool)) *Lex
	Not(interface{}) bool
	NotSome(...interface{}) bool
	Values() []interface{}
}

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

// AddLength adds a new key and value to the map and returns the length of the modified map.
func (lex *Lex) AddLength(k interface{}, v interface{}) int { return lex.Add(k, v).Len() }

// AddOK adds a key and value to the map and returns a boolean confirming whether the key, value pair was successfully added.
func (lex *Lex) AddOK(k interface{}, v interface{}) bool {
	var (
		ok bool
	)
	if lex.Not(k) {
		lex.Add(k, v)
		ok = true
	}
	return ok
}

// Del deletes the key and value from the map and returns the modified map.
func (lex *Lex) Del(k interface{}) *Lex {
	delete((*lex), k)
	return lex
}

// DelAll deletes all keys and values from the map and returns the modified map.
func (lex *Lex) DelAll() *Lex {
	(*lex) = (Lex{})
	return lex
}

// DelLength deletes the key and value from the map and returns the length of modified map.
func (lex *Lex) DelLength(k interface{}) int { return (lex.Del(k).Len()) }

// DelSome deletes some keys and values from the map and returns the modified map. Arguments are treated as keys to the map.
func (lex *Lex) DelSome(k ...interface{}) *Lex {
	for _, k := range k {
		lex.Del(k)
	}
	return lex
}

// DelSomeLength deletes some keys and values from the map and returns the length of modified map.
func (lex *Lex) DelSomeLength(k ...interface{}) int { return (lex.DelSome(k...).Len()) }

// DelOK deletes the key and value from the map and returns a boolean on the outcome of the transaction.
func (lex *Lex) DelOK(k interface{}) bool { return (lex.Del(k).Has(k) == false) }

// Each executes a provided function once for each map key, value pair.
func (lex *Lex) Each(fn func(k interface{}, v interface{})) *Lex {
	var (
		k, v interface{}
	)
	for k, v = range *lex {
		fn(k, v)
	}
	return lex
}

// EachBreak executes a provided function once for each key, value with an optional break when the function returns false.
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

// EachKey executes a provided function once for each key in the map.
func (lex *Lex) EachKey(fn func(k interface{})) *Lex {
	lex.Each(func(k, _ interface{}) {
		fn(k)
	})
	return lex
}

// EachValue executes a provided function once for each value in the map.
func (lex *Lex) EachValue(fn func(v interface{})) *Lex {
	lex.Each(func(_, v interface{}) {
		fn(v)
	})
	return lex
}

// Fetch retrieves the value held by the argument key without checking if the key exists.
func (lex *Lex) Fetch(k interface{}) interface{} {
	var v, _ = lex.Get(k)
	return v
}

// FetchSome retrieves a slice of values from the map if they are found in the argument of keys.
func (lex *Lex) FetchSome(k ...interface{}) []interface{} {
	var (
		s = []interface{}{}
	)
	for _, k := range k {
		var (
			v = lex.Fetch(k)
		)
		if v != nil {
			s = append(s, v)
		}
	}
	return s
}

// Get gets the value from the map using the provided key. Returns a boolean on the outcome of the transaction.
func (lex *Lex) Get(k interface{}) (interface{}, bool) {
	var v, ok = (*lex)[k]
	return v, ok
}

// GetLength gets the value from the map using the provided key and returns the length of the map.
// Returns an additional boolean on the outcome of the transaction.
func (lex *Lex) GetLength(k interface{}) (interface{}, int, bool) {
	var v, ok = lex.Get(k)
	var l = lex.Len()
	return v, l, ok
}

// Has checks that the map has a key of the corresponding value in the map.
func (lex *Lex) Has(k interface{}) bool {
	var _, ok = lex.Get(k)
	return ok
}

// Keys returns a slice of the maps keys in the order found.
func (lex *Lex) Keys() []interface{} {
	var (
		s = []interface{}{}
	)
	lex.Each(func(k, _ interface{}) {
		s = append(s, k)
	})
	return s
}

// Len returns the number of values in the map.
func (lex *Lex) Len() int { return (len(*lex)) }

// Map executes a provided function once for each key, value in the map and sets the returned value to the current key.
func (lex *Lex) Map(fn func(k interface{}, v interface{}) interface{}) *Lex {
	lex.Each(func(k interface{}, v interface{}) {
		lex.Add(k, fn(k, v))
	})
	return lex
}

// MapBreak executes a provided function once for each key, value in the map and sets
// returned value to the current key with an optional break when the function returns false.
func (lex *Lex) MapBreak(fn func(k interface{}, v interface{}) (interface{}, bool)) *Lex {
	var ok bool
	lex.EachBreak(func(k, v interface{}) bool {
		v, ok = fn(k, v)
		if ok {
			lex.Add(k, v)
		}
		return ok
	})
	return lex
}

// MapOK executes a provided function once for each key, value in the map and sets
// the returned value to the current key if a boolean of true is returned.
func (lex *Lex) MapOK(fn func(k interface{}, v interface{}) (interface{}, bool)) *Lex {
	var ok bool
	lex.Each(func(k interface{}, v interface{}) {
		v, ok = fn(k, v)
		if ok {
			lex.Add(k, v)
		}
	})
	return lex
}

// Not checks that the map does not contain the argument key in the map.
func (lex *Lex) Not(k interface{}) bool { return (lex.Has(k) == false) }

// NotSome checks that the map does not contain the series of key in the map.
func (lex *Lex) NotSome(k ...interface{}) bool {
	var ok = true
	var x interface{}
	for _, x = range k {
		ok = lex.Not(x)
		if !ok {
			break
		}
	}
	return ok
}

// Values returns a slice of the map values in order found.
func (lex *Lex) Values() []interface{} {
	var (
		s = []interface{}{}
	)
	lex.Each(func(_, v interface{}) {
		s = append(s, v)
	})
	return s
}
