package hashtable

import (
	"reflect"

	"github.com/lindsaygelle/slice"
)

// Hashtable represents a generic hash table that maps keys of type K to values of type V. It provides efficient key-value storage and retrieval operations.
type Hashtable[K comparable, V any] map[K]V

// Add inserts a new key-value pair into the hashtable or updates the existing value associated with the provided key.
// If the key already exists, the corresponding value is updated. If the key is new, a new key-value pair is added to the hashtable.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//	newHashtable.Add("banana", 10) // Updates the value for the key "banana" to 10
//	fmt.Println(newHashtable) // &map[apple:5 banana:10 cherry:8]
func (hashtable *Hashtable[K, V]) Add(key K, value V) *Hashtable[K, V] {
	(*hashtable)[key] = value
	return hashtable
}

// AddFunc inserts key-value pairs into the hashtable based on the values provided in the input slice of maps.
// It applies the provided function to each key-value pair and adds them to the hashtable if the function returns true.
// If the key already exists in the hashtable, the corresponding value is updated with the new value from the input.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.AddFunc([]map[string]int{{"apple": 1, "orange": 2}}, func(key string, value int) bool {
//	    // Only add key-value pairs where the value is greater than 1.
//	    return value > 1
//	})
//	fmt.Println(newHashtable) // &map[orange:10]
func (hashtable *Hashtable[K, V]) AddFunc(values []map[K]V, fn func(key K, value V) bool) *Hashtable[K, V] {
	for _, item := range values {
		for key, value := range item {
			if fn(key, value) {
				hashtable.Add(key, value)
			}
		}
	}
	return hashtable
}

// AddLength inserts a new key-value pair into the hashtable or updates the existing value associated with the provided key.
// If the key already exists, the corresponding value is updated. If the key is new, a new key-value pair is added to the hashtable.
// It then returns the current length of the hashtable after the addition or update operation.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	length := newHashtable.AddLength("apple", 5)  // Adds "apple" with value 5, returns the length of the hashtable (1 in this case)
//	length = newHashtable.AddLength("apple", 10)  // Updates the value for "apple" to 10, returns the length of the hashtable (1)
//	length = newHashtable.AddLength("banana", 3)  // Adds "banana" with value 3, returns the length of the hashtable (2)
func (hashtable *Hashtable[K, V]) AddLength(key K, value V) int {
	return hashtable.Add(key, value).Length()
}

// AddMany inserts multiple key-value pairs into the hashtable. It accepts a variadic number of maps, where each map contains
// key-value pairs to be added to the hashtable. If a key already exists in the hashtable, the corresponding value is updated
// with the new value from the input maps.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.AddMany(map[string]int{"orange": 7, "grape": 4}, map[string]int{"kiwi": 6, "pear": 9})
//	fmt.Println(newHashtable) // &map[orange:7 grape:4 kiwi:6 pear:9]
func (hashtable *Hashtable[K, V]) AddMany(values ...map[K]V) *Hashtable[K, V] {
	for _, item := range values {
		for key, value := range item {
			hashtable.Add(key, value)
		}
	}
	return hashtable
}

// AddManyFunc inserts key-value pairs into the hashtable based on a provided condition function.
// It accepts a slice of maps, where each map contains key-value pairs. For each key-value pair,
// the specified function is called. If the function returns true, the pair is added to the hashtable.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.AddManyFunc([]map[K]V{{"apple": 5, "orange": -3, "banana": 10}}, func(i int, key string, value int) bool {
//		return value > 0 // Add key-value pairs with values greater than 0
//	})
//	fmt.Println(newHashtable) // &map[apple:5 banana:10]
func (hashtable *Hashtable[K, V]) AddManyFunc(values []map[K]V, fn func(i int, key K, value V) bool) *Hashtable[K, V] {
	for i, item := range values {
		for key, value := range item {
			if fn(i, key, value) {
				hashtable.Add(key, value)
			}
		}
	}
	return hashtable
}

// AddManyOK inserts multiple key-value pairs into the hashtable and returns a slice of booleans indicating whether each insertion was successful.
// It accepts a variadic number of maps, where each map contains key-value pairs to be added to the hashtable.
// For each key-value pair, it checks if the key already exists in the hashtable. If the key is not present, the pair is added,
// and the corresponding boolean in the returned slice is true. If the key already exists, the pair is not added, and the boolean is false.
//
//	ht := make(Hashtable[string, int])
//	results := ht.AddManyOK(map[string]int{"apple": 5, "orange": 3}, map[string]int{"orange": 10, "banana": 7})
//	// Returns a slice containing [true, false, true] indicating successful insertions for "apple" and "banana"
func (hashtable *Hashtable[K, V]) AddManyOK(values ...map[K]V) *slice.Slice[bool] {
	successfulInsertions := make(slice.Slice[bool], 0)
	for _, item := range values {
		for key, value := range item {
			ok := hashtable.Not(key)
			if ok {
				hashtable.Add(key, value)
			}
			successfulInsertions.Append(ok)
		}
	}
	return &successfulInsertions
}

// AddOK inserts a new key-value pair into the hashtable only if the key does not already exist in the hashtable.
// If the key already exists, the insertion fails, and false is returned. If the key is new, a new key-value pair is added to the hashtable,
// and true is returned to indicate a successful insertion.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//
//	// Attempt to add key-value pairs.
//	added := newHashtable.AddOK("apple", 5)    // added is true, "apple" is added with value 5.
//	reAdded := newHashtable.AddOK("apple", 10)  // reAdded is false, "apple" already exists with value 5, no change is made.
//	addedNew := newHashtable.AddOK("banana", 3) // addedNew is true, "banana" is added with value 3.
func (hashtable *Hashtable[K, V]) AddOK(key K, value V) bool {
	ok := !hashtable.Has(key)
	if ok {
		hashtable.Add(key, value)
	}
	return ok
}

// Delete removes a key-value pair from the hashtable based on the provided key. If the key exists in the hashtable,
// it is deleted, and the modified hashtable is returned. If the key is not found, the hashtable remains unchanged.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//
//	// Delete the key-value pair with the key "apple".
//	newHashtable.Delete("apple")
//	fmt.Println(newHashtable) // &map[banana:3]
func (hashtable *Hashtable[K, V]) Delete(key K) *Hashtable[K, V] {
	delete(*hashtable, key)
	return hashtable
}

// DeleteFunc removes key-value pairs from the hashtable based on the provided function. The function is applied to each key-value pair,
// and if it returns true, the corresponding key-value pair is deleted from the hashtable.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//
//	// Delete key-value pairs where the value is less than 4.
//	newHashtable.DeleteFunc(func(key string, value int) bool {
//		return value < 4
//	})
//	fmt.Println(newHashtable) // &map[apple:5]
func (hashtable *Hashtable[K, V]) DeleteFunc(fn func(key K, value V) bool) *Hashtable[K, V] {
	for key, value := range *hashtable {
		if fn(key, value) {
			hashtable.Delete(key)
		}
	}
	return hashtable
}

// DeleteLength removes a key-value pair from the hashtable based on the provided key. If the key exists in the hashtable,
// it is deleted, and the current length of the hashtable after the deletion is returned. If the key is not found,
// the hashtable remains unchanged, and the current length is returned.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//
//	// Delete the key-value pair with the key "apple" and get the updated length of the hashtable.
//	length := newHashtable.DeleteLength("apple")
//	// After deletion, the length of the hashtable is 1.
//	// The current length returned: 1
func (hashtable *Hashtable[K, V]) DeleteLength(key K) int {
	return hashtable.Delete(key).Length()
}

// DeleteMany removes multiple key-value pairs from the hashtable based on the provided keys. If a key exists in the hashtable,
// it is deleted.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//
//	// Delete key-value pairs with the keys "apple" and "banana".
//	newHashtable.DeleteMany("apple", "banana")
//	fmt.Println(newHashtable) // &map[]
func (hashtable *Hashtable[K, V]) DeleteMany(keys ...K) *Hashtable[K, V] {
	for _, key := range keys {
		hashtable.Delete(key)
	}
	return hashtable
}

// DeleteManyOK removes multiple key-value pairs from the hashtable based on the provided keys. If a key exists in the hashtable,
// it is deleted, and true is appended to the result slice to indicate a successful deletion. If the key is not found, false is appended.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//
//	// Attempt to delete key-value pairs with the keys "apple" and "orange".
//	results := newHashtable.DeleteManyOK("apple", "orange")
//	// Results after deletion: []bool{true, false}
//	// The first deletion succeeded ("apple": 5 was deleted), and the second deletion failed as "orange" was not found.
func (hashtable *Hashtable[K, V]) DeleteManyOK(keys ...K) *slice.Slice[bool] {
	deletetions := make(slice.Slice[bool], 0)
	for _, key := range keys {
		deletetions.Append(hashtable.DeleteOK(key))
	}
	return &deletetions
}

// DeleteManyValues removes key-value pairs from the hashtable based on the provided values. If a value exists in the hashtable,
// the corresponding key-value pair is deleted.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//
//	// Delete key-value pairs with the values 5 and 10.
//	newHashtable.DeleteManyValues(5, 10)
//	// Hashtable after deletion: {"banana": 3}
func (hashtable *Hashtable[K, V]) DeleteManyValues(values ...V) *Hashtable[K, V] {
	for key, value := range *hashtable {
		for _, v := range values {
			if reflect.DeepEqual(v, value) {
				hashtable.Delete(key)
			}
		}
	}
	return hashtable
}

// DeleteOK removes a key-value pair from the hashtable based on the provided key. If the key exists in the hashtable,
// it is deleted, and true is returned to indicate a successful deletion. If the key is not found, the hashtable remains unchanged,
// and false is returned to indicate that the deletion failed.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//
//	// Attempt to delete the key-value pair with the key "apple".
//	success := newHashtable.DeleteOK("apple")
//	// After deletion, the key "apple" is not found in the hashtable.
//	// Deletion success: true
//
//	// Attempt to delete a non-existing key.
//	success = newHashtable.DeleteOK("orange")
//	// The key "orange" does not exist in the hashtable.
//	// Deletion failed: false
func (hashtable *Hashtable[K, V]) DeleteOK(key K) bool {
	return !hashtable.Delete(key).Has(key)
}

// Each iterates over the key-value pairs in the hashtable and applies a function to each pair.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//
//	// Function to print all key-value pairs.
//	printKeyValue := func(key string, value int) {
//	    fmt.Println(key, value)
//	}
//
//	// Iterate over the hashtable and print all key-value pairs.
//	newHashtable.Each(printKeyValue)
//	// Output: "apple 5", "banana 3", "cherry 8"
func (hashtable *Hashtable[K, V]) Each(fn func(key K, value V)) *Hashtable[K, V] {
	return hashtable.EachBreak(func(key K, value V) bool {
		fn(key, value)
		return true
	})
}

// EachBreak applies the provided function to each key-value pair in the hashtable. The function is applied to key-value pairs
// in the hashtable until the provided function returns false. If the function returns false for any key-value pair,
// the iteration breaks early.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//
//	// Function to print key-value pairs until finding "banana".
//	stopPrinting := newHashtable.EachBreak(func(key string, value int) bool {
//	    fmt.Println(key, value)
//	    return key != "banana" // Continue printing until "banana" is encountered.
//	})
//	// Output: "apple 5", "banana 3"
func (hashtable *Hashtable[K, V]) EachBreak(fn func(key K, value V) bool) *Hashtable[K, V] {
	for key, value := range *hashtable {
		if !fn(key, value) {
			break
		}
	}
	return hashtable
}

// EachKey iterates over the keys in the hashtable and applies a function to each key.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//
//	// Function to print each key.
//	printKey := func(key string) {
//	    fmt.Println(key)
//	}
//
//	// Iterate over the hashtable and print each key.
//	newHashtable.EachKey(printKey)
//	// Output: "apple", "banana", "cherry"
func (hashtable *Hashtable[K, V]) EachKey(fn func(key K)) *Hashtable[K, V] {
	return hashtable.Each(func(key K, _ V) {
		fn(key)
	})
}

// EachKeyBreak iterates over the keys in the hashtable and applies a function to each key. It allows breaking the iteration early if the provided function returns false.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//
//	// Function to print each key and break the iteration if the key is "banana".
//	printAndBreak := func(key string) bool {
//	    fmt.Println(key)
//	    return key != "banana"
//	}
//
//	// Iterate over the hashtable keys, print them, and break when "banana" is encountered.
//	newHashtable.EachKeyBreak(printAndBreak)
//	// Output: "apple", "banana"
func (hashtable *Hashtable[K, V]) EachKeyBreak(fn func(key K) bool) *Hashtable[K, V] {
	return hashtable.EachBreak(func(key K, _ V) bool {
		return fn(key)
	})
}

// EachValue iterates over the values in the hashtable and applies a function to each value.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//
//	// Function to print each value.
//	printValue := func(value int) {
//	    fmt.Println(value)
//	}
//
//	// Iterate over the hashtable values and print them.
//	newHashtable.EachValue(printValue)
//	// Output: 5, 3, 8
func (hashtable *Hashtable[K, V]) EachValue(fn func(value V)) *Hashtable[K, V] {
	return hashtable.Each(func(_ K, value V) {
		fn(value)
	})
}

// EachValueBreak iterates over the values in the hashtable and applies a function to each value until the function returns false.
// If the provided function returns false, the iteration breaks early.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//
//	// Function to process each value. Returns false to break the iteration if the value is 3.
//	processValue := func(value int) bool {
//	    fmt.Println(value)
//	    return value != 3
//	}
//
//	// Iterate over the hashtable values and process them until the value is 3.
//	newHashtable.EachValueBreak(processValue)
//	// Output: 5, 3
func (hashtable *Hashtable[K, V]) EachValueBreak(fn func(value V) bool) *Hashtable[K, V] {
	return hashtable.EachBreak(func(_ K, value V) bool {
		return fn(value)
	})
}

// Get retrieves the value associated with the provided key from the hashtable.
// If the key exists, it returns the associated value and true. Otherwise, it returns the zero value for the value type and false.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	value, exists := newHashtable.Get("apple") // 5, true
//	value, exists = newHashtable.Get("orange")  // 0, false
func (hashtable *Hashtable[K, V]) Get(key K) (V, bool) {
	value, ok := (*hashtable)[key]
	return value, ok
}

// Filter applies the given function to each key-value pair in the hashtable and returns a new hashtable
// containing only the key-value pairs for which the function returns true. The original hashtable is not modified.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//
//	// Function to filter key-value pairs. Returns true if the value is greater than 4.
//	filterFunc := func(key string, value int) bool {
//		return value > 4
//	}
//
//	// Create a new hashtable containing key-value pairs where the value is greater than 4.
//	filteredHashtable := newHashtable.Filter(filterFunc)
func (hashtable *Hashtable[K, V]) Filter(fn func(key K, value V) bool) *Hashtable[K, V] {
	filteredHashtable := make(Hashtable[K, V], 0)
	hashtable.Each(func(key K, value V) {
		if fn(key, value) {
			filteredHashtable.Add(key, value)
		}
	})
	return &filteredHashtable
}

// GetMany retrieves the values associated with the provided keys from the hashtable. It accepts a variadic number of keys,
// and returns a slice containing the values corresponding to the keys found in the hashtable. If a key is not found in the hashtable,
// the corresponding position in the returned slice will be the zero value for the value type.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//
//	// Get values for specific keys.
//	values := newHashtable.GetMany("apple", "banana", "orange")
//
//	// The resulting values slice: {5, 3}
func (hashtable *Hashtable[K, V]) GetMany(keys ...K) *slice.Slice[V] {
	values := &slice.Slice[V]{}
	for _, key := range keys {
		if value, ok := hashtable.Get(key); ok {
			values.Append(value)
		}
	}
	return values
}

// Has checks if the provided key exists in the hashtable.
// It returns true if the key exists, and false otherwise.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	exists := newHashtable.Has("apple") // true
//	exists = newHashtable.Has("orange")  // false
func (hashtable *Hashtable[K, V]) Has(key K) bool {
	_, ok := (*hashtable)[key]
	return ok
}

// HasMany checks the existence of multiple keys in the hashtable and returns a slice of boolean values
// indicating whether each corresponding key exists in the hashtable.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//
//	// Check the existence of multiple keys.
//	keysToCheck := []string{"apple", "orange", "banana"}
//	results := newHashtable.HasMany(keysToCheck...)
//
//	// The resulting boolean slice: {true, false, true}
func (hashtable *Hashtable[K, V]) HasMany(keys ...K) *slice.Slice[bool] {
	values := make(slice.Slice[bool], len(keys))
	for i, key := range keys {
		if hashtable.Has(key) {
			values.Replace(i, true)
		}
	}
	return &values
}

// Keys returns a slice containing all the keys present in the hashtable.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//
//	// Get all keys from the hashtable.
//	keys := newHashtable.Keys() // Result: {"apple", "banana", "cherry"}
func (hashtable *Hashtable[K, V]) Keys() *slice.Slice[K] {
	keys := make(slice.Slice[K], 0)
	hashtable.EachKey(func(key K) {
		keys.Append(key)
	})
	return &keys
}

// KeysFunc applies the provided function to each key in the hashtable and returns a slice containing the keys for which the function returns true.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//
//	// Get keys from the hashtable where the key length is greater than 5.
//	keys := newHashtable.KeysFunc(func(key string) bool {
//	    return len(key) > 5
//	})
//	// Result: {"banana"}
func (hashtable *Hashtable[K, V]) KeysFunc(fn func(key K) bool) *slice.Slice[K] {
	keys := make(slice.Slice[K], 0)
	hashtable.EachKey(func(key K) {
		if fn(key) {
			keys.Append(key)
		}
	})
	return &keys
}

// Length returns the number of key-value pairs in the hashtable.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//
//	length := newHashtable.Length() // Result: 3
func (hashtable *Hashtable[K, V]) Length() int {
	return len(*hashtable)
}

// Map applies the provided function to each key-value pair in the hashtable and returns a new hashtable containing the mapped key-value pairs.
// The original hashtable remains unchanged.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//
//	// Function to double the values of each key-value pair.
//	doubleValues := func(key string, value int) int {
//		return value * 2
//	}
//
//	// Create a new hashtable with doubled values.
//	newHashtable := newHashtable.Map(doubleValues)
//	// New hashtable: {"apple": 10, "banana": 6}
func (hashtable *Hashtable[K, V]) Map(fn func(key K, value V) V) *Hashtable[K, V] {
	return hashtable.MapBreak(func(key K, value V) (V, bool) {
		return fn(key, value), true
	})
}

// MapBreak applies the provided function to each key-value pair in the hashtable. It creates a new hashtable containing the mapped key-value pairs
// until the function returns false for any pair, at which point the mapping breaks early. The original hashtable remains unchanged.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Add("cherry", 8)
//
//	// Create a new hashtable with doubled values until a value is greater than or equal to 10.
//	newHashtable := newHashtable.MapBreak(func(key string, value int) (int, bool) {
//		newValue := value * 2
//		return newValue, newValue < 10
//	})
//	// New hashtable: {"apple": 10, "banana": 6}
func (hashtable *Hashtable[K, V]) MapBreak(fn func(key K, value V) (V, bool)) *Hashtable[K, V] {
	newHashtable := make(Hashtable[K, V])
	for key, value := range *hashtable {
		value, ok := fn(key, value)
		if !ok {
			break
		}
		newHashtable.Add(key, value)
	}
	return &newHashtable
}

// Not checks if the given key is not present in the hashtable.
// It returns true if the key is not found, and false if the key exists in the hashtable.
//
//	ht := make(Hashtable[string, int])
//	result := ht.Not("apple")  // Returns true if "apple" is not in the hashtable, false otherwise
func (hashtable *Hashtable[K, V]) Not(key K) bool {
	return !hashtable.Has(key)
}

// Pop removes a key-value pair from the hashtable based on the provided key.
// It returns the removed value and a boolean indicating whether the key was found and removed successfully.
// If the key is present in the hashtable, the corresponding value is returned, and the key-value pair is deleted.
// If the key is not found, it returns the zero value for the value type and false.
//
//	ht := make(Hashtable[string, int])
//	ht.Add("apple", 5)
//	removedValue, ok := ht.Pop("apple")  // Removes the key "apple" and returns its associated value 5, ok is true
//	removedValue, ok = ht.Pop("banana")   // Key "banana" not found, removedValue is 0 and ok is false
func (hashtable *Hashtable[K, V]) Pop(key K) (V, bool) {
	value, ok := hashtable.Get(key)
	if ok {
		ok = hashtable.DeleteOK(key)
	}
	return value, ok
}

// PopMany removes multiple key-value pairs from the hashtable based on the provided keys.
// It takes a variadic number of keys as input and removes the corresponding key-value pairs from the hashtable.
// It returns a slice containing the removed values and does not guarantee any specific order of values in the result.
// If a key is not found in the hashtable, the corresponding value in the result slice will be the zero value for the value type.
//
//	ht := make(Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	removedValues := ht.PopMany("apple", "orange")  // Removes "apple", returns a slice containing [5, 0]
func (hashtable *Hashtable[K, V]) PopMany(keys ...K) *slice.Slice[V] {
	values := make(slice.Slice[V], 0)
	for _, key := range keys {
		value, ok := hashtable.Pop(key)
		if ok {
			values.Append(value)
		}
	}
	return &values
}

// Update iterates over the key-value pairs in the hashtable and applies the provided function to each pair.
// The function can modify the value and return a boolean indicating whether the update should be performed.
// If the function returns true, the key-value pair is updated in the same hashtable with the modified value.
// If the function returns false, the key-value pair is not modified.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("banana", 3)
//	newHashtable.Update(func(key string, value int) (int, bool) {
//	    if key == "banana" {
//	        return value * 2, true // Modify the value for the "banana" key
//	    }
//	    return value, false // Leave other values unchanged
//	})
//	// newHashtable: {"apple": 5, "banana": 6}
func (hashtable *Hashtable[K, V]) Update(fn func(key K, value V) (V, bool)) *Hashtable[K, V] {
	for key, value := range *hashtable {
		if updatedValue, ok := fn(key, value); ok {
			hashtable.Add(key, updatedValue)
		}
	}
	return hashtable
}

// Values returns a slice containing all the values present in the hashtable.
// It iterates over the hashtable and collects all the values in the order of insertion.
//
//	newHashtable := make(hashtable.Hashtable[string, int])
//	newHashtable.Add("apple", 5)
//	newHashtable.Add("orange", 10)
//	values := newHashtable.Values()  // Returns a slice containing [5, 10]
func (hashtable *Hashtable[K, V]) Values() *slice.Slice[V] {
	i := 0
	values := make(slice.Slice[V], hashtable.Length())
	hashtable.EachValue(func(value V) {
		values.Replace(i, value)
		i++
	})
	return &values
}

// ValuesFunc returns a slice containing values from the hashtable that satisfy the given condition function.
// The condition function takes a key-value pair as input and returns true if the pair meets the condition, false otherwise.
// It iterates over the hashtable and includes the values in the returned slice for which the condition function evaluates to true.
//
//	ht := make(Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("orange", 10)
//	filterFunc := func(key string, value int) bool {
//		return value > 7  // Include values greater than 7 in the result
//	}
//	values := ht.ValuesFunc(filterFunc)  // Returns a slice containing [10]
func (hashtable *Hashtable[K, V]) ValuesFunc(fn func(key K, value V) bool) *slice.Slice[V] {
	values := make(slice.Slice[V], 0)
	hashtable.Each(func(key K, value V) {
		if fn(key, value) {
			values.Append(value)
		}
	})
	return &values
}
