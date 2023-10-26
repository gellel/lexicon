package gomap

import (
	"reflect"

	"github.com/lindsaygelle/slice"
)

// Map represents a generic map that maps keys of type K to values of type V.
type Map[K comparable, V any] map[K]V

// Add inserts a new key-value pair into the map or updates the existing value associated with the provided key.
// If the key already exists, the corresponding value is updated. If the key is new, a new key-value pair is added to the map.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//	newMap.Add("banana", 10) // Updates the value for the key "banana" to 10
//	fmt.Println(newMap) // &map[apple:5 banana:10 cherry:8]
func (gomap *Map[K, V]) Add(key K, value V) *Map[K, V] {
	(*gomap)[key] = value
	return gomap
}

// AddLength inserts a new key-value pair into the map or updates the existing value associated with the provided key.
// If the key already exists, the corresponding value is updated. If the key is new, a new key-value pair is added to the map.
// It then returns the current length of the map after the addition or update operation.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	length := newMap.AddLength("apple", 5)  // Adds "apple" with value 5, returns the length of the map (1 in this case)
//	length = newMap.AddLength("apple", 10)  // Updates the value for "apple" to 10, returns the length of the map (1)
//	length = newMap.AddLength("banana", 3)  // Adds "banana" with value 3, returns the length of the map (2)
func (gomap *Map[K, V]) AddLength(key K, value V) int {
	return gomap.Add(key, value).Length()
}

// AddMany inserts multiple key-value pairs into the map. It accepts a variadic number of maps, where each map contains
// key-value pairs to be added to the map. If a key already exists in the gomap, the corresponding value is updated
// with the new value from the input maps.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.AddMany(map[string]int{"orange": 7, "grape": 4}, map[string]int{"kiwi": 6, "pear": 9})
//	fmt.Println(newMap) // &map[orange:7 grape:4 kiwi:6 pear:9]
func (gomap *Map[K, V]) AddMany(values ...map[K]V) *Map[K, V] {
	for _, item := range values {
		for key, value := range item {
			gomap.Add(key, value)
		}
	}
	return gomap
}

// AddManyFunc inserts key-value pairs into the map based on a provided condition function.
// It accepts a slice of maps, where each map contains key-value pairs. For each key-value pair,
// the specified function is called. If the function returns true, the pair is added to the map.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.AddManyFunc([]map[K]V{{"apple": 5, "orange": -3, "banana": 10}}, func(i int, key string, value int) bool {
//		return value > 0 // Add key-value pairs with values greater than 0
//	})
//	fmt.Println(newMap) // &map[apple:5 banana:10]
func (gomap *Map[K, V]) AddManyFunc(values []map[K]V, fn func(i int, key K, value V) bool) *Map[K, V] {
	for i, item := range values {
		for key, value := range item {
			if fn(i, key, value) {
				gomap.Add(key, value)
			}
		}
	}
	return gomap
}

// AddManyOK inserts multiple key-value pairs into the map and returns a slice of booleans indicating whether each insertion was successful.
// It accepts a variadic number of maps, where each map contains key-value pairs to be added to the map.
// For each key-value pair, it checks if the key already exists in the map. If the key is not present, the pair is added,
// and the corresponding boolean in the returned slice is true. If the key already exists, the pair is not added, and the boolean is false.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	results := newMap.AddManyOK(map[string]int{"apple": 5, "orange": 3}, map[string]int{"orange": 10, "banana": 7})
//	// Returns a slice containing [true, false, true] indicating successful insertions for "apple" and "banana"
func (gomap *Map[K, V]) AddManyOK(values ...map[K]V) *slice.Slice[bool] {
	successfulInsertions := make(slice.Slice[bool], 0)
	for _, item := range values {
		for key, value := range item {
			ok := gomap.Not(key)
			if ok {
				gomap.Add(key, value)
			}
			successfulInsertions.Append(ok)
		}
	}
	return &successfulInsertions
}

// AddOK inserts a new key-value pair into the map only if the key does not already exist in the map.
// If the key already exists, the insertion fails, and false is returned. If the key is new, a new key-value pair is added to the gomap,
// and true is returned to indicate a successful insertion.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//
//	// Attempt to add key-value pairs.
//	added := newMap.AddOK("apple", 5)    // added is true, "apple" is added with value 5.
//	reAdded := newMap.AddOK("apple", 10)  // reAdded is false, "apple" already exists with value 5, no change is made.
//	addedNew := newMap.AddOK("banana", 3) // addedNew is true, "banana" is added with value 3.
func (gomap *Map[K, V]) AddOK(key K, value V) bool {
	ok := !gomap.Has(key)
	if ok {
		gomap.Add(key, value)
	}
	return ok
}

// Contains checks if the given value is present in the map and returns the first key-value pair that matches the value.
// It takes a value as input and returns the key and a boolean indicating whether the value is found in the map.
// If the value is found, it returns the corresponding key and true. If the value is not found, it returns the zero value for the key type and false.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	key, found := newMap.Contains(5)  // Checks if value 5 is in the gomap, returns the key ("apple" for example) and true if found, or ("", false) if not found
func (gomap *Map[K, V]) Contains(value V) (K, bool) {
	var k K
	var ok bool
	gomap.EachBreak(func(key K, v V) bool {
		ok = reflect.DeepEqual(v, value)
		if ok {
			k = key
		}
		return !ok
	})
	return k, ok
}

// Delete removes a key-value pair from the map based on the provided key. If the key exists in the gomap,
// it is deleted, and the modified map is returned. If the key is not found, the map remains unchanged.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//
//	// Delete the key-value pair with the key "apple".
//	newMap.Delete("apple")
//	fmt.Println(newMap) // &map[banana:3]
func (gomap *Map[K, V]) Delete(key K) *Map[K, V] {
	delete(*gomap, key)
	return gomap
}

// DeleteLength removes a key-value pair from the map based on the provided key. If the key exists in the gomap,
// it is deleted, and the current length of the map after the deletion is returned. If the key is not found,
// the map remains unchanged, and the current length is returned.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//
//	// Delete the key-value pair with the key "apple" and get the updated length of the map.
//	length := newMap.DeleteLength("apple")
//	// After deletion, the length of the map is 1.
//	// The current length returned: 1
func (gomap *Map[K, V]) DeleteLength(key K) int {
	return gomap.Delete(key).Length()
}

// DeleteMany removes multiple key-value pairs from the map based on the provided keys. If a key exists in the gomap,
// it is deleted.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//
//	// Delete key-value pairs with the keys "apple" and "banana".
//	newMap.DeleteMany("apple", "banana")
//	fmt.Println(newMap) // &map[]
func (gomap *Map[K, V]) DeleteMany(keys ...K) *Map[K, V] {
	for _, key := range keys {
		gomap.Delete(key)
	}
	return gomap
}

// DeleteFunc removes key-value pairs from the map based on the provided function. The function is applied to each key-value pair,
// and if it returns true, the corresponding key-value pair is deleted from the map.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//
//	// Delete key-value pairs where the value is less than 4.
//	newMap.DeleteFunc(func(key string, value int) bool {
//		return value < 4
//	})
//	fmt.Println(newMap) // &map[apple:5]
func (gomap *Map[K, V]) DeleteManyFunc(fn func(key K, value V) bool) *Map[K, V] {
	for key, value := range *gomap {
		if fn(key, value) {
			gomap.Delete(key)
		}
	}
	return gomap
}

// DeleteManyOK removes multiple key-value pairs from the map based on the provided keys. If a key exists in the gomap,
// it is deleted, and true is appended to the result slice to indicate a successful deletion. If the key is not found, false is appended.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//
//	// Attempt to delete key-value pairs with the keys "apple" and "orange".
//	results := newMap.DeleteManyOK("apple", "orange")
//	// Results after deletion: []bool{true, false}
//	// The first deletion succeeded ("apple": 5 was deleted), and the second deletion failed as "orange" was not found.
func (gomap *Map[K, V]) DeleteManyOK(keys ...K) *slice.Slice[bool] {
	deletetions := make(slice.Slice[bool], 0)
	for _, key := range keys {
		deletetions.Append(gomap.DeleteOK(key))
	}
	return &deletetions
}

// DeleteManyValues removes key-value pairs from the map based on the provided values. If a value exists in the gomap,
// the corresponding key-value pair is deleted.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//
//	// Delete key-value pairs with the values 5 and 10.
//	newMap.DeleteManyValues(5, 10)
//	// Map after deletion: {"banana": 3}
func (gomap *Map[K, V]) DeleteManyValues(values ...V) *Map[K, V] {
	for key, value := range *gomap {
		for _, v := range values {
			if reflect.DeepEqual(v, value) {
				gomap.Delete(key)
			}
		}
	}
	return gomap
}

// DeleteOK removes a key-value pair from the map based on the provided key. If the key exists in the gomap,
// it is deleted, and true is returned to indicate a successful deletion. If the key is not found, the map remains unchanged,
// and false is returned to indicate that the deletion failed.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//
//	// Attempt to delete the key-value pair with the key "apple".
//	success := newMap.DeleteOK("apple")
//	// After deletion, the key "apple" is not found in the map.
//	// Deletion success: true
//
//	// Attempt to delete a non-existing key.
//	success = newMap.DeleteOK("orange")
//	// The key "orange" does not exist in the map.
//	// Deletion failed: false
func (gomap *Map[K, V]) DeleteOK(key K) bool {
	return !gomap.Delete(key).Has(key)
}

// Each iterates over the key-value pairs in the map and applies a function to each pair.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//
//	// Iterate over the map and print all key-value pairs.
//	newMap.Each(func(key string, value int) {
//	    fmt.Println(key, value)
//	})
//	// Output: "apple 5", "banana 3", "cherry 8"
func (gomap *Map[K, V]) Each(fn func(key K, value V)) *Map[K, V] {
	return gomap.EachBreak(func(key K, value V) bool {
		fn(key, value)
		return true
	})
}

// EachBreak applies the provided function to each key-value pair in the map. The function is applied to key-value pairs
// in the map until the provided function returns false. If the function returns false for any key-value pair,
// the iteration breaks early.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//
//	// Function to print key-value pairs until finding "banana".
//	stopPrinting := newMap.EachBreak(func(key string, value int) bool {
//	    fmt.Println(key, value)
//	    return key != "banana" // Continue printing until "banana" is encountered.
//	})
//	// Output: "apple 5", "banana 3"
func (gomap *Map[K, V]) EachBreak(fn func(key K, value V) bool) *Map[K, V] {
	for key, value := range *gomap {
		if !fn(key, value) {
			break
		}
	}
	return gomap
}

// EachKey iterates over the keys in the map and applies a function to each key.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//
//	// Iterate over the map and print each key.
//	newMap.EachKey(func(key string) {
//	    fmt.Println(key)
//	})
//	// Output: "apple", "banana", "cherry"
func (gomap *Map[K, V]) EachKey(fn func(key K)) *Map[K, V] {
	return gomap.Each(func(key K, _ V) {
		fn(key)
	})
}

// EachKeyBreak iterates over the keys in the map and applies a function to each key. It allows breaking the iteration early if the provided function returns false.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//
//	// Iterate over the map keys, print them, and break when "banana" is encountered.
//	newMap.EachKeyBreak(func(key string) bool {
//	    fmt.Println(key)
//	    return key != "banana"
//	})
//	// Output: "apple", "banana"
func (gomap *Map[K, V]) EachKeyBreak(fn func(key K) bool) *Map[K, V] {
	return gomap.EachBreak(func(key K, _ V) bool {
		return fn(key)
	})
}

// EachValue iterates over the values in the map and applies a function to each value.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//
//	// Iterate over the map values and print them.
//	newMap.EachValue(func(value int) {
//	    fmt.Println(value)
//	})
//	// Output: 5, 3, 8
func (gomap *Map[K, V]) EachValue(fn func(value V)) *Map[K, V] {
	return gomap.Each(func(_ K, value V) {
		fn(value)
	})
}

// EachValueBreak iterates over the values in the map and applies a function to each value until the function returns false.
// If the provided function returns false, the iteration breaks early.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//
//	// Iterate over the map values and process them until the value is 3.
//	newMap.EachValueBreak(func(value int) bool {
//	    fmt.Println(value)
//	    return value != 3
//	})
//	// Output: 5, 3
func (gomap *Map[K, V]) EachValueBreak(fn func(value V) bool) *Map[K, V] {
	return gomap.EachBreak(func(_ K, value V) bool {
		return fn(value)
	})
}

// EmptyInto transfers all key-value pairs from the current map into another gomap, emptying the current map.
// It takes another map as input and adds all key-value pairs from the current map to the other map.
//
//	// Create a new Map instance.
//	newMap1 := make(gomap.Map[string, int])
//	newMap1.Add("apple", 5)
//	// Create a new Map instance.
//	newMap2 := make(gomap.Map[string, int])
//
//	newMap1.EmptyInto(newMap2)  // Transfers "apple": 5 from newMap1 to newMap2, leaving newMap1 empty
func (gomap *Map[K, V]) EmptyInto(other *Map[K, V]) *Map[K, V] {
	gomap.Each(func(key K, value V) {
		other.Add(key, gomap.Pop(key))
	})
	return gomap
}

// Equal checks if the current map is equal to another map by comparing the key-value pairs directly using reflect.DeepEqual.
// It takes another map as input and returns true if the two hashtables are equal, false otherwise.
//
//	// Create a new Map instance.
//	newMap1 := make(gomap.Map[string, int])
//	newMap1.Add("apple", 5)
//	newMap1.Add("orange", 10)
//
//	// Create a new Map instance.
//	newMap2 := make(gomap.Map[string, int])
//	newMap2.Add("apple", 5)
//	newMap2.Add("orange", 10)
//
//	equal := newMap1.Equal(newMap2)  // Returns true because newMap1 and newMap2 have the same key-value pairs
func (gomap *Map[K, V]) Equal(other *Map[K, V]) bool {
	return gomap.EqualFunc(other, func(a, b V) bool {
		return reflect.DeepEqual(a, b)
	})
}

// EqualFunc checks if the current map is equal to another map based on a provided comparison function.
// It takes another map and a comparison function as input and returns true if the two hashtables are equal according to the function.
// The comparison function takes two values as input and returns true if they are considered equal, false otherwise.
//
//	// Create a new Map instance.
//	newMap1 := make(gomap.Map[string, int])
//	newMap1.Add("apple", 5)
//	newMap1.Add("orange", 10)
//
//	// Create a new Map instance.
//	newMap2 := make(gomap.Map[string, int])
//	newMap2.Add("apple", 5)
//	newMap2.Add("orange", 11)
//
//	equal := newMap1.EqualFunc(newMap2, func(a, b int) bool {
//		return math.Abs(float64(a - b)) <= 1
//	})  // Returns true because the values for "orange" (10 and 11) have a difference of 1, within the allowed range
func (gomap *Map[K, V]) EqualFunc(other *Map[K, V], fn func(a V, b V) bool) bool {
	if !gomap.EqualLength(other) {
		return false
	}
	for key, value := range *gomap {
		v, ok := other.Get(key)
		if !ok || !fn(value, v) {
			return false
		}
	}
	return true
}

// EqualLength checks if the current map has the same length as another map.
// It takes another map as input and returns true if the two hashtables have the same length, false otherwise.
//
//	// Create a new Map instance.
//	newMap1 := make(gomap.Map[string, int])
//	newMap1.Add("apple", 5)
//	newMap1.Add("orange", 10)
//
//	// Create a new Map instance.
//	newMap2 := make(gomap.Map[string, int])
//	newMap2.Add("apple", 5)
//
//	equalLength := newMap1.EqualLength(newMap2)  // Returns false because newMap1 has a length of 2, while newMap2 has a length of 1
func (gomap *Map[K, V]) EqualLength(other *Map[K, V]) bool {
	return gomap.Length() == other.Length()
}

// Fetch retrieves the value associated with the given key from the map.
// It returns the value if the key is found in the gomap, and the zero value for the value type if the key is not present.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	value := newMap.Fetch("apple")  // Returns 5, the value associated with the key "apple"
//	value = newMap.Fetch("orange")  // Returns 0 because "orange" is not in the gomap
func (gomap *Map[K, V]) Fetch(key K) V {
	value, _ := gomap.Get(key)
	return value
}

// Get retrieves the value associated with the provided key from the map.
// If the key exists, it returns the associated value and true. Otherwise, it returns the zero value for the value type and false.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	value, exists := newMap.Get("apple") // 5, true
//	value, exists = newMap.Get("orange")  // 0, false
func (gomap *Map[K, V]) Get(key K) (V, bool) {
	value, ok := (*gomap)[key]
	return value, ok
}

// Filter applies the given function to each key-value pair in the map and returns a new gomap
// containing only the key-value pairs for which the function returns true. The original map is not modified.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//
//	// Create a new map containing key-value pairs where the value is greater than 4.
//	filteredMap := newMap.Filter(func(key string, value int) bool {
//		return value > 4
//	})
func (gomap *Map[K, V]) Filter(fn func(key K, value V) bool) *Map[K, V] {
	other := make(Map[K, V], 0)
	gomap.Each(func(key K, value V) {
		if fn(key, value) {
			other.Add(key, value)
		}
	})
	return &other
}

// GetMany retrieves the values associated with the provided keys from the map. It accepts a variadic number of keys,
// and returns a slice containing the values corresponding to the keys found in the map. If a key is not found in the gomap,
// the corresponding position in the returned slice will be the zero value for the value type.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//
//	// Get values for specific keys.
//	values := newMap.GetMany("apple", "banana", "orange") // The resulting values slice: {5, 3}
func (gomap *Map[K, V]) GetMany(keys ...K) *slice.Slice[V] {
	values := &slice.Slice[V]{}
	for _, key := range keys {
		if value, ok := gomap.Get(key); ok {
			values.Append(value)
		}
	}
	return values
}

// Has checks if the provided key exists in the map.
// It returns true if the key exists, and false otherwise.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	exists := newMap.Has("apple") // true
//	exists = newMap.Has("orange")  // false
func (gomap *Map[K, V]) Has(key K) bool {
	_, ok := (*gomap)[key]
	return ok
}

// HasMany checks the existence of multiple keys in the map and returns a slice of boolean values
// indicating whether each corresponding key exists in the map.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//
//	// Check the existence of multiple keys.
//	keysToCheck := []string{"apple", "orange", "banana"}
//	results := newMap.HasMany(keysToCheck...)
//
//	// The resulting boolean slice: {true, false, true}
func (gomap *Map[K, V]) HasMany(keys ...K) *slice.Slice[bool] {
	values := make(slice.Slice[bool], len(keys))
	for i, key := range keys {
		if gomap.Has(key) {
			values.Replace(i, true)
		}
	}
	return &values
}

// Intersection creates a new map containing key-value pairs that exist in both the current map and another map.
// It compares values using reflect.DeepEqual to determine equality between the pairs.
// It takes another map as input and returns a new map containing the intersecting key-value pairs.
//
//	// Create a new Map instance.
//	newMap1 := make(gomap.Map[string, int])
//	newMap1.Add("apple", 5)
//
//	// Create a new Map instance.
//	newMap2 := make(gomap.Map[string, int])
//	newMap2.Add("apple", 5)
//	newMap2.Add("orange", 10)
//
//	newMap := newMap1.Intersection(newMap2)  // Creates a new map with the pair "apple": 5
func (gomap *Map[K, V]) Intersection(other *Map[K, V]) *Map[K, V] {
	return gomap.IntersectionFunc(other, func(key K, a, b V) bool {
		return reflect.DeepEqual(a, b)
	})
}

// IntersectionFunc creates a new map containing key-value pairs that exist in both the current map and another map.
// It takes another map and a condition function as input and adds key-value pairs from the current map to the new gomap
// if the condition function evaluates to true for the corresponding key-value pair in the other map.
//
//	// Create a new Map instance.
//	newMap1 := make(gomap.Map[string, int])
//	newMap1.Add("apple", 5)
//	newMap1.Add("orange", 8)
//
//	// Create a new Map instance.
//	newMap2 := make(gomap.Map[string, int])
//	newMap2.Add("orange", 8)
//	newMap2.Add("banana", 6)
//
//	newMap := newMap1.IntersectionFunc(newMap2, func(key string, a, b int) bool {
//		return a == b
//	})  // Creates a new map with the pair "orange": 8
func (gomap *Map[K, V]) IntersectionFunc(other *Map[K, V], fn func(key K, a V, b V) bool) *Map[K, V] {
	newMap := make(Map[K, V], 0)
	gomap.Each(func(key K, value V) {
		if v, ok := other.Get(key); ok && fn(key, value, v) {
			newMap.Add(key, value)
		}
	})
	return &newMap
}

// IsEmpty checks if the map is empty, i.e., it contains no key-value pairs.
// It returns true if the map is empty and false otherwise.
//
//	// Create a new Map instance.
//	newMap := make(Map[string, int])
//	empty := newMap.IsEmpty()  // Returns true since the map is empty
func (gomap *Map[K, V]) IsEmpty() bool {
	return gomap.Length() == 0
}

func (gomap *Map[K, V]) IsPopulated() bool {
	return !gomap.IsEmpty()
}

// Keys returns a slice containing all the keys present in the map.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//
//	// Get all keys from the map.
//	keys := newMap.Keys() // Result: {"apple", "banana", "cherry"}
func (gomap *Map[K, V]) Keys() *slice.Slice[K] {
	keys := make(slice.Slice[K], 0)
	gomap.EachKey(func(key K) {
		keys.Append(key)
	})
	return &keys
}

// KeysFunc applies the provided function to each key in the map and returns a slice containing the keys for which the function returns true.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//
//	// Get keys from the map where the key length is greater than 5.
//	keys := newMap.KeysFunc(func(key string) bool {
//	    return len(key) > 5
//	})
//	// Result: {"banana"}
func (gomap *Map[K, V]) KeysFunc(fn func(key K) bool) *slice.Slice[K] {
	keys := make(slice.Slice[K], 0)
	gomap.EachKey(func(key K) {
		if fn(key) {
			keys.Append(key)
		}
	})
	return &keys
}

// Length returns the number of key-value pairs in the map.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//
//	length := newMap.Length() // Result: 3
func (gomap *Map[K, V]) Length() int {
	return len(*gomap)
}

// Map applies the provided function to each key-value pair in the map and returns a new map containing the mapped key-value pairs.
// The original map remains unchanged.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//
//	// Create a new map with doubled values.
//	newMap := newMap.Map(func(key string, value int) int {
//		return value * 2
//	})
//	// New gomap: {"apple": 10, "banana": 6}
func (gomap *Map[K, V]) Map(fn func(key K, value V) V) *Map[K, V] {
	return gomap.MapBreak(func(key K, value V) (V, bool) {
		return fn(key, value), true
	})
}

// MapBreak applies the provided function to each key-value pair in the map. It creates a new map containing the mapped key-value pairs
// until the function returns false for any pair, at which point the mapping breaks early. The original map remains unchanged.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Add("cherry", 8)
//
//	// Create a new map with doubled values until a value is greater than or equal to 10.
//	newMap := newMap.MapBreak(func(key string, value int) (int, bool) {
//		newValue := value * 2
//		return newValue, newValue < 10
//	})
//	// New gomap: {"apple": 10, "banana": 6}
func (gomap *Map[K, V]) MapBreak(fn func(key K, value V) (V, bool)) *Map[K, V] {
	newMap := make(Map[K, V])
	for key, value := range *gomap {
		value, ok := fn(key, value)
		if !ok {
			break
		}
		newMap.Add(key, value)
	}
	return &newMap
}

// Merge merges all key-value pairs from another map into the current map.
// It takes another map as input and adds all its key-value pairs to the current map.
//
//	// Create a new Map instance.
//	newMap1 := make(gomap.Map[string, int])
//	newMap1.Add("apple", 5)
//
//	// Create a new Map instance.
//	newMap2 := make(gomap.Map[string, int])
//	newMap2.Add("orange", 10)
//
//	newMap1.Merge(newMap2)  // Adds "orange": 10 to newMap1
func (gomap *Map[K, V]) Merge(other *Map[K, V]) *Map[K, V] {
	return gomap.MergeFunc(other, func(key K, value V) bool { return true })
}

// MergeFunc merges the key-value pairs from another map into the current map based on a provided condition function.
// It takes another map and a condition function as input and adds key-value pairs from the other map to the current gomap
// if the condition function evaluates to true for the given key-value pair from the other map.
//
//	// Create a new Map instance.
//	newMap1 := make(gomap.Map[string, int])
//	newMap1.Add("apple", 5)
//	newMap1.Add("orange", 10)
//
//	// Create a new Map instance.
//	newMap2 := make(gomap.Map[string, int])
//	newMap2.Add("orange", 8)
//	newMap2.Add("banana", 6)
//
//	// Condition function to merge pairs where the value in newMap2 is greater than 7
//	newMap1.MergeFunc(newMap2, func(key string, value int) bool {
//		return value > 7
//	})  // Adds "orange": 8 to newMap1, "banana": 6 does not meet the condition and is not added
func (gomap *Map[K, V]) MergeFunc(other *Map[K, V], fn func(key K, value V) bool) *Map[K, V] {
	other.Each(func(key K, value V) {
		if fn(key, value) {
			gomap.Add(key, value)
		}
	})
	return gomap
}

// MergeMany merges key-value pairs from multiple hashtables into the current map.
// It takes a variadic number of hashtables as input and adds all their key-value pairs to the current map.
//
//	// Create a new Map instance.
//	newMap1 := make(gomap.Map[string, int])
//	newMap1.Add("apple", 5)
//	// Create a new Map instance.
//	newMap2 := make(gomap.Map[string, int])
//	newMap2.Add("orange", 10)
//	// Create a new Map instance.
//	newMap3 := make(gomap.Map[string, int])
//	newMap3.Add("banana", 7)
//
//	newMap1.MergeMany(newMap2, newMap3)  // Merges key-value pairs from newMap2 and newMap3 into newMap1
func (gomap *Map[K, V]) MergeMany(others ...*Map[K, V]) *Map[K, V] {
	for _, other := range others {
		gomap.Merge(other)
	}
	return gomap
}

// MergeManyFunc merges key-value pairs from multiple hashtables into the current map based on a provided condition function.
// It takes a slice of hashtables and a condition function as input. For each key-value pair in the other hashtables,
// the function is applied, and if it evaluates to true, the pair is added to the current map.
//
//	// Create a new Map instance.
//	newMap1 := make(gomap.Map[string, int])
//	newMap1.Add("apple", 5)
//	// Create a new Map instance.
//	newMap2 := make(gomap.Map[string, int])
//	newMap2.Add("orange", 10)
//	// Create a new Map instance.
//	newMap3 := make(gomap.Map[string, int])
//	newMap3.Add("banana", 7)
//
//	// Condition function to include pairs from the second map only if the value is greater than 7
//	newMap1.MergeManyFunc([]*Map[string, int]{newMap2, newMap3}, func(i int, key string, value int) bool {
//		return value > 7
//	})
func (gomap *Map[K, V]) MergeManyFunc(others []*Map[K, V], fn func(i int, key K, value V) bool) *Map[K, V] {
	for i, other := range others {
		gomap.MergeFunc(other, func(key K, value V) bool {
			return fn(i, key, value)
		})
	}
	return gomap
}

// Not checks if the given key is not present in the map.
// It returns true if the key is not found, and false if the key exists in the map.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	result := newMap.Not("apple")  // Returns true if "apple" is not in the gomap, false otherwise
func (gomap *Map[K, V]) Not(key K) bool {
	return !gomap.Has(key)
}

// NotMany checks if multiple keys are not present in the map.
// It takes a variadic number of keys as input and returns a slice of booleans indicating whether each key is not found in the map.
// For each key, if it is not present in the gomap, the corresponding boolean in the returned slice is true. Otherwise, it is false.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	results := newMap.NotMany("apple", "orange", "banana")
//	// Returns a slice containing [true, true, false] indicating "apple" and "orange" are not in the gomap, but "banana" is present
func (gomap *Map[K, V]) NotMany(keys ...K) *slice.Slice[bool] {
	values := make(slice.Slice[bool], len(keys))
	for i, key := range keys {
		if gomap.Not(key) {
			values.Replace(i, true)
		}
	}
	return &values
}

// Pop removes a key-value pair from the map based on the provided key and returns the removed value.
// If the key is found in the gomap, the corresponding value is returned. If the key is not present,
// the zero value for the value type is returned.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	removedValue := newMap.Pop("apple")  // Removes the key "apple" and returns its associated value 5, or 0 if "apple" is not found
func (gomap *Map[K, V]) Pop(key K) V {
	value, _ := gomap.PopOK(key)
	return value
}

// PopOK removes a key-value pair from the map based on the provided key.
// It returns the removed value and a boolean indicating whether the key was found and removed successfully.
// If the key is present in the gomap, the corresponding value is returned, and the key-value pair is deleted.
// If the key is not found, it returns the zero value for the value type and false.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	removedValue, ok := newMap.PopOK("apple")  // Removes the key "apple" and returns its associated value 5, ok is true
//	removedValue, ok = newMap.PopOK("banana")   // Key "banana" not found, removedValue is 0 and ok is false
func (gomap *Map[K, V]) PopOK(key K) (V, bool) {
	value, ok := gomap.Get(key)
	if ok {
		ok = gomap.DeleteOK(key)
	}
	return value, ok
}

// PopMany removes multiple key-value pairs from the map based on the provided keys.
// It takes a variadic number of keys as input and removes the corresponding key-value pairs from the map.
// It returns a slice containing the removed values and does not guarantee any specific order of values in the result.
// If a key is not found in the gomap, the corresponding value in the result slice will be the zero value for the value type.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	removedValues := newMap.PopMany("apple", "orange")  // Removes "apple", returns a slice containing [5, 0]
func (gomap *Map[K, V]) PopMany(keys ...K) *slice.Slice[V] {
	values := make(slice.Slice[V], 0)
	for _, key := range keys {
		value, ok := gomap.PopOK(key)
		if ok {
			values.Append(value)
		}
	}
	return &values
}

// PopManyFunc removes key-value pairs from the map based on the provided condition function and returns the removed values in a slice.
// It takes a condition function as input and removes key-value pairs for which the function evaluates to true.
// It returns a slice containing the removed values.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("orange", 10)
//
//	removedValues := newMap.PopManyFunc(func(key string, value int) bool {
//		return value > 7  // Remove values greater than 7
//	})  // Returns a slice containing [10]
func (gomap *Map[K, V]) PopManyFunc(fn func(key K, value V) bool) *slice.Slice[V] {
	values := make(slice.Slice[V], 0)
	gomap.Each(func(key K, value V) {
		if fn(key, value) {
			removedValue := gomap.Pop(key)
			values.Append(removedValue)
		}
	})
	return &values
}

// ReplaceMany iterates over the key-value pairs in the map and applies the provided function to each pair.
// The function can modify the value and return a boolean indicating whether the update should be performed.
// If the function returns true, the key-value pair is updated in the same map with the modified value.
// If the function returns false, the key-value pair is not modified.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("banana", 3)
//	newMap.Replace(func(key string, value int) (int, bool) {
//	    if key == "banana" {
//	        return value * 2, true // Modify the value for the "banana" key
//	    }
//	    return value, false // Leave other values unchanged
//	})
//	// newMap: {"apple": 5, "banana": 6}
func (gomap *Map[K, V]) ReplaceMany(fn func(key K, value V) (V, bool)) *Map[K, V] {
	for key, value := range *gomap {
		if updatedValue, ok := fn(key, value); ok {
			gomap.Add(key, updatedValue)
		}
	}
	return gomap
}

// TakeFrom transfers all key-value pairs from another map into the current gomap, emptying the other map.
// It takes another map as input and adds all key-value pairs from the other map to the current map.
//
//	// Create a new Map instance.
//	newMap1 := make(gomap.Map[string, int])
//	newMap1.Add("apple", 5)
//	// Create a new Map instance.
//	newMap2 := make(gomap.Map[string, int])
//	newMap2.Add("orange", 10)
//
//	newMap1.TakeFrom(newMap2)  // Transfers "orange": 10 from newMap2 to newMap1, leaving newMap2 empty
func (gomap *Map[K, V]) TakeFrom(other *Map[K, V]) *Map[K, V] {
	other.Each(func(key K, value V) {
		gomap.Add(key, other.Pop(key))
	})
	return gomap
}

// Values returns a slice containing all the values present in the map.
// It iterates over the map and collects all the values in the order of insertion.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("orange", 10)
//	values := newMap.Values()  // Returns a slice containing [5, 10]
func (gomap *Map[K, V]) Values() *slice.Slice[V] {
	i := 0
	values := make(slice.Slice[V], gomap.Length())
	gomap.EachValue(func(value V) {
		values.Replace(i, value)
		i++
	})
	return &values
}

// ValuesFunc returns a slice containing values from the map that satisfy the given condition function.
// The condition function takes a key-value pair as input and returns true if the pair meets the condition, false otherwise.
// It iterates over the map and includes the values in the returned slice for which the condition function evaluates to true.
//
//	// Create a new Map instance.
//	newMap := make(gomap.Map[string, int])
//	newMap.Add("apple", 5)
//	newMap.Add("orange", 10)
//	values := newMap.ValuesFunc(func(key string, value int) bool {
//		return value > 7  // Include values greater than 7 in the result
//	})  // Returns a slice containing [10]
func (gomap *Map[K, V]) ValuesFunc(fn func(key K, value V) bool) *slice.Slice[V] {
	values := make(slice.Slice[V], 0)
	gomap.Each(func(key K, value V) {
		if fn(key, value) {
			values.Append(value)
		}
	})
	return &values
}
