package hashtable

import (
	"reflect"

	"github.com/lindsaygelle/slice"
)

// Hashtable represents a generic hash table that maps keys of type K to values of type V. It provides efficient key-value storage and retrieval operations.
//
// Example:
//
//	// Create a new Hashtable with string keys and integer values.
//	ht := make(hashtable.Hashtable[string, int])
//
//	// Insert key-value pairs into the hashtable.
//	ht["apple"] = 5
//	ht["banana"] = 3
//	ht["cherry"] = 8
//
//	// Retrieve the value associated with a specific key.
//	value := ht["banana"] // 3
//
//	// Check if a key exists in the hashtable.
//	_, exists := ht["grape"]
//
//	// Delete a key-value pair from the hashtable.
//	delete(ht, "cherry")
type Hashtable[K comparable, V any] map[K]V

// Add inserts a key-value pair into the hashtable. If the key already exists, its associated value is updated.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Add("cherry", 8)
//	ht.Add("banana", 10) // Updates the value associated with the "banana" key.
func (hashtable *Hashtable[K, V]) Add(key K, value V) *Hashtable[K, V] {
	(*hashtable)[key] = value
	return hashtable
}

// AddFunc inserts key-value pairs into the hashtable based on the provided maps and a custom validation function.
// The validation function should return true for key-value pairs that should be added to the hashtable, and false otherwise.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.AddFunc([]map[string]int{{"apple": 1, "orange": 2}}, func(key string, value int) bool {
//	    // Only add key-value pairs where the value is greater than 1.
//	    return value > 1
//	})
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

// AddLength inserts a key-value pair into the hashtable and returns the new length of the hashtable.
// If the key already exists, its associated value is updated.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	length := ht.AddLength("apple", 5) // length is 1
//	length = ht.AddLength("banana", 3)  // length is 2
//	length = ht.AddLength("apple", 10)  // length remains 2 (key "apple" is updated)
func (hashtable *Hashtable[K, V]) AddLength(key K, value V) int {
	return hashtable.Add(key, value).Length()
}

// AddMany inserts multiple key-value pairs into the hashtable from the provided maps.
// If keys already exist, their associated values are updated.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.AddMany(map[string]int{"orange": 7, "grape": 4}, map[string]int{"kiwi": 6, "pear": 9})
func (hashtable *Hashtable[K, V]) AddMany(values ...map[K]V) *Hashtable[K, V] {
	for _, item := range values {
		for key, value := range item {
			hashtable.Add(key, value)
		}
	}
	return hashtable
}

// AddManyOK inserts multiple key-value pairs into the hashtable and returns a slice of booleans indicating
// whether each insertion was successful. If a key already exists, it is not updated, and the corresponding
// boolean value is set to false in the returned slice.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	results := ht.AddManyOK(
//	    map[string]int{"apple": 5, "banana": 3},
//	    map[string]int{"banana": 10, "cherry": 8},
//	)
//	// results contains [true, false, true] indicating successful insertions for "apple" and "cherry"
//	// and unsuccessful insertion for "banana" due to existing key.
func (hashtable *Hashtable[K, V]) AddManyOK(values ...map[K]V) *slice.Slice[bool] {
	successfulInsertions := make(slice.Slice[bool], 0)
	for _, item := range values {
		for key, value := range item {
			successfulInsertions.Append(hashtable.AddOK(key, value))
		}
	}
	return &successfulInsertions
}

// AddOK inserts a new key-value pair into the hashtable if the key does not already exist.
// It returns a boolean value indicating whether the key was added successfully (true) or if the key already existed (false).
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//
//	// Attempt to add key-value pairs.
//	added := ht.AddOK("apple", 5)    // added is true, "apple" is added with value 5.
//	reAdded := ht.AddOK("apple", 10)  // reAdded is false, "apple" already exists with value 5, no change is made.
//	addedNew := ht.AddOK("banana", 3) // addedNew is true, "banana" is added with value 3.
func (hashtable *Hashtable[K, V]) AddOK(key K, value V) bool {
	ok := !hashtable.Has(key)
	if ok {
		hashtable.Add(key, value)
	}
	return ok
}

// Delete removes a key-value pair from the hashtable based on the provided key.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Delete("apple") // Removes the key-value pair with key "apple" from the hashtable.
func (hashtable *Hashtable[K, V]) Delete(key K) *Hashtable[K, V] {
	delete(*hashtable, key)
	return hashtable
}

// DeleteFunc removes key-value pairs from the hashtable based on the evaluation performed by the provided function.
// For each key-value pair in the hashtable, the function fn is called with the key and value.
// If the function returns true, the key-value pair is removed from the hashtable.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("orange", 7)
//	ht.Add("kiwi", 6)
//
//	// Delete key-value pairs where the value is less than 7.
//	ht.DeleteFunc(func(key string, value int) bool {
//	    return value < 7
//	}) // Removes the key-value pair with key "kiwi" from the hashtable.
func (hashtable *Hashtable[K, V]) DeleteFunc(fn func(key K, value V) bool) *Hashtable[K, V] {
	for key, value := range *hashtable {
		if fn(key, value) {
			hashtable.Delete(key)
		}
	}
	return hashtable
}

// DeleteLength deletes a key from the hashtable and returns the new length of the hashtable.
// If the key does not exist, the length remains unchanged.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	length := ht.DeleteLength("apple") // length is 1 (key "apple" is deleted)
//	length = ht.DeleteLength("grape")   // length remains 1 (key "grape" does not exist)
func (hashtable *Hashtable[K, V]) DeleteLength(key K) int {
	return hashtable.Delete(key).Length()
}

// DeleteMany removes multiple key-value pairs from the hashtable based on the provided keys.
// If a key doesn't exist in the hashtable, it is ignored.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("orange", 7)
//	ht.Add("kiwi", 6)
//	ht.DeleteMany("apple", "kiwi") // Removes key-value pairs with keys "apple" and "kiwi" from the hashtable.
func (hashtable *Hashtable[K, V]) DeleteMany(keys ...K) *Hashtable[K, V] {
	for _, key := range keys {
		hashtable.Delete(key)
	}
	return hashtable
}

// DeleteManyOK deletes multiple keys from the hashtable and returns a slice of booleans indicating whether each deletion was successful.
// For each specified key, it checks if the key exists in the hashtable before attempting deletion. If the key does not exist,
// the deletion is considered unsuccessful for that key, and false is appended to the returned slice. If the key exists and is successfully
// deleted, true is appended; otherwise, false is appended.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	keysToDelete := []string{"apple", "grape"}
//	results := ht.DeleteManyOK(keysToDelete...)
//	// results contains [true, true], indicating successful deletion of "apple" (exists) and "grape" (does not exist)
func (hashtable *Hashtable[K, V]) DeleteManyOK(keys ...K) *slice.Slice[bool] {
	deletetions := make(slice.Slice[bool], 0)
	for _, key := range keys {
		deletetions.Append(hashtable.DeleteOK(key))
	}
	return &deletetions
}

// DeleteManyValues deletes key-value pairs from the hashtable where the value matches any of the specified values.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Add("cherry", 8)
//
//	// Delete key-value pairs where the value is 3 or 8.
//	ht.DeleteManyValues(3, 8)
//
//	// The hashtable after deletion: {"apple": 5}
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

// DeleteOK deletes the specified key from the hashtable and returns a boolean indicating whether the deletion was successful.
// If the key does not exist in the hashtable, it is considered a successful deletion, and true is returned.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	deleted := ht.DeleteOK("apple") // true, "apple" key is successfully deleted
//	notDeleted := ht.DeleteOK("grape") // true, "grape" key does not exist, deletion is considered successful
func (hashtable *Hashtable[K, V]) DeleteOK(key K) bool {
	return !hashtable.Delete(key).Has(key)
}

// Each iterates over the key-value pairs in the hashtable and applies a function to each pair.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Add("cherry", 8)
//
//	// Function to print all key-value pairs.
//	printKeyValue := func(key string, value int) {
//	    fmt.Println(key, value)
//	}
//
//	// Iterate over the hashtable and print all key-value pairs.
//	ht.Each(printKeyValue)
//	// Output: "apple 5", "banana 3", "cherry 8"
func (hashtable *Hashtable[K, V]) Each(fn func(key K, value V)) *Hashtable[K, V] {
	return hashtable.EachBreak(func(key K, value V) bool {
		fn(key, value)
		return true
	})
}

// EachBreak iterates over the key-value pairs in the hashtable and applies a function to each pair.
// If the function returns false at any point, the iteration breaks.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Add("cherry", 8)
//
//	// Function to print key-value pairs until finding "banana".
//	stopPrinting := ht.EachBreak(func(key string, value int) bool {
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
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Add("cherry", 8)
//
//	// Function to print each key.
//	printKey := func(key string) {
//	    fmt.Println(key)
//	}
//
//	// Iterate over the hashtable and print each key.
//	ht.EachKey(printKey)
//	// Output: "apple", "banana", "cherry"
func (hashtable *Hashtable[K, V]) EachKey(fn func(key K)) *Hashtable[K, V] {
	return hashtable.Each(func(key K, _ V) {
		fn(key)
	})
}

// EachKeyBreak iterates over the keys in the hashtable and applies a function to each key. It allows breaking the iteration early if the provided function returns false.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Add("cherry", 8)
//
//	// Function to print each key and break the iteration if the key is "banana".
//	printAndBreak := func(key string) bool {
//	    fmt.Println(key)
//	    return key != "banana"
//	}
//
//	// Iterate over the hashtable keys, print them, and break when "banana" is encountered.
//	ht.EachKeyBreak(printAndBreak)
//	// Output: "apple", "banana"
func (hashtable *Hashtable[K, V]) EachKeyBreak(fn func(key K) bool) *Hashtable[K, V] {
	return hashtable.EachBreak(func(key K, _ V) bool {
		return fn(key)
	})
}

// EachValue iterates over the values in the hashtable and applies a function to each value.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Add("cherry", 8)
//
//	// Function to print each value.
//	printValue := func(value int) {
//	    fmt.Println(value)
//	}
//
//	// Iterate over the hashtable values and print them.
//	ht.EachValue(printValue)
//	// Output: 5, 3, 8
func (hashtable *Hashtable[K, V]) EachValue(fn func(value V)) *Hashtable[K, V] {
	return hashtable.Each(func(_ K, value V) {
		fn(value)
	})
}

// EachValueBreak iterates over the values in the hashtable and applies a function to each value until the function returns false.
// If the provided function returns false, the iteration breaks early.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Add("cherry", 8)
//
//	// Function to process each value. Returns false to break the iteration if the value is 3.
//	processValue := func(value int) bool {
//	    fmt.Println(value)
//	    return value != 3
//	}
//
//	// Iterate over the hashtable values and process them until the value is 3.
//	ht.EachValueBreak(processValue)
//	// Output: 5, 3
func (hashtable *Hashtable[K, V]) EachValueBreak(fn func(value V) bool) *Hashtable[K, V] {
	return hashtable.EachBreak(func(_ K, value V) bool {
		return fn(value)
	})
}

// Get retrieves the value associated with the provided key from the hashtable.
// If the key exists, it returns the associated value and true. Otherwise, it returns the zero value for the value type and false.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	value, exists := ht.Get("apple") // 5, true
//	value, exists = ht.Get("orange")  // 0, false
func (hashtable *Hashtable[K, V]) Get(key K) (V, bool) {
	value, ok := (*hashtable)[key]
	return value, ok
}

// GetMany retrieves values from the hashtable for the specified keys and returns them as a slice.
// If a key is not found in the hashtable, it is skipped in the result slice.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Add("cherry", 8)
//
//	// Get values for specific keys.
//	values := ht.GetMany("apple", "banana", "orange")
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
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	exists := ht.Has("apple") // true
//	exists = ht.Has("orange")  // false
func (hashtable *Hashtable[K, V]) Has(key K) bool {
	_, ok := (*hashtable)[key]
	return ok
}

// HasMany checks the existence of multiple keys in the hashtable and returns a slice of boolean values
// indicating whether each corresponding key exists in the hashtable.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Add("cherry", 8)
//
//	// Check the existence of multiple keys.
//	keysToCheck := []string{"apple", "orange", "banana"}
//	results := ht.HasMany(keysToCheck...)
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
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Add("cherry", 8)
//
//	// Get all keys from the hashtable.
//	keys := ht.Keys() // Result: {"apple", "banana", "cherry"}
func (hashtable *Hashtable[K, V]) Keys() *slice.Slice[K] {
	keys := make(slice.Slice[K], 0)
	hashtable.EachKey(func(key K) {
		keys.Append(key)
	})
	return &keys
}

// KeysFunc returns a slice containing the keys from the hashtable for which the provided function returns true.
// The provided function `fn` should accept a key of type `K` and return a boolean value.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Add("cherry", 8)
//
//	// Get keys from the hashtable where the key length is greater than 5.
//	keys := ht.KeysFunc(func(key string) bool {
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
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Add("cherry", 8)
//
//	length := ht.Length() // Result: 3
func (hashtable *Hashtable[K, V]) Length() int {
	return len(*hashtable)
}

// Map iterates over the key-value pairs in the hashtable and applies the provided function to each pair.
// The function can modify the value. The modified key-value pairs are updated in the same hashtable.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	ht.Map(func(key string, value int) int {
//	    if key == "banana" {
//	        return value * 2 // Modify the value for the "banana" key
//	    }
//	    return value // Leave other values unchanged
//	})
//	// ht: {"apple": 5, "banana": 6}
func (hashtable *Hashtable[K, V]) Map(fn func(key K, value V) V) *Hashtable[K, V] {
	return hashtable.MapBreak(func(key K, value V) (V, bool) {
		return fn(key, value), true
	})
}

// MapBreak iterates over the key-value pairs in the hashtable and applies the provided function to each pair.
// The function can modify the value and return a boolean indicating whether to continue the iteration.
// If the function returns false, the iteration breaks, and a new hashtable with modified key-value pairs is returned.
//
// Example:
//
//	ht := make(hashtable.Hashtable[string, int])
//	ht.Add("apple", 5)
//	ht.Add("banana", 3)
//	newHT := ht.MapBreak(func(key string, value int) (int, bool) {
//	    if key == "banana" {
//	        return value * 2, false // Break the iteration when key is "banana"
//	    }
//	    return value, true // Continue iterating for other keys
//	})
//	// newHT: {"apple": 5}
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
