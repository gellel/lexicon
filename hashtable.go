package hashtable

// Hashtable represents a generic hash table that maps keys of type K to values of type V. It provides efficient key-value storage and retrieval operations.
//
// Example:
//   // Create a new Hashtable with string keys and integer values.
//   ht := make(hashtable.Hashtable[string, int])
//
//   // Insert key-value pairs into the hashtable.
//   ht["apple"] = 5
//   ht["banana"] = 3
//   ht["cherry"] = 8
//
//   // Retrieve the value associated with a specific key.
//   value := ht["banana"] // 3
//
//   // Check if a key exists in the hashtable.
//   _, exists := ht["grape"]
//
//   // Delete a key-value pair from the hashtable.
//   delete(ht, "cherry")
type Hashtable[K comparable, V any] map[K]V

// Add inserts a key-value pair into the hashtable. If the key already exists, its associated value is updated.
//
// Example:
//   ht := make(hashtable.Hashtable[string, int])
//   ht.Add("apple", 5)
//   ht.Add("banana", 3)
//   ht.Add("cherry", 8)
//   ht.Add("banana", 10) // Updates the value associated with the "banana" key.
func (hashtable *Hashtable[K, V]) Add(key K, value V) *Hashtable[K, V] {
	(*hashtable)[key] = value
	return hashtable
}

// AddFunc inserts key-value pairs into the hashtable based on the provided maps and a custom validation function.
// The validation function should return true for key-value pairs that should be added to the hashtable, and false otherwise.
//
// Example:
//   ht := make(hashtable.Hashtable[string, int])
//   ht.AddFunc([]map[string]int{{"apple": 1, "orange": 2}}, func(key string, value int) bool {
//       // Only add key-value pairs where the value is greater than 1.
//       return value > 1
//   })
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

// AddMany inserts multiple key-value pairs into the hashtable from the provided maps.
// If keys already exist, their associated values are updated.
//
// Example:
//   ht := make(hashtable.Hashtable[string, int])
//   ht.AddMany(map[string]int{"orange": 7, "grape": 4}, map[string]int{"kiwi": 6, "pear": 9})
func (hashtable *Hashtable[K, V]) AddMany(values ...map[K]V) *Hashtable[K, V] {
	for _, item := range values {
		for key, value := range item {
			hashtable.Add(key, value)
		}
	}
	return hashtable
}

// AddOK inserts a new key-value pair into the hashtable if the key does not already exist.
// It returns a boolean value indicating whether the key was added successfully (true) or if the key already existed (false).
//
// Example:
//   ht := make(hashtable.Hashtable[string, int])
//
//   // Attempt to add key-value pairs.
//   added := ht.AddOK("apple", 5)    // added is true, "apple" is added with value 5.
//   reAdded := ht.AddOK("apple", 10)  // reAdded is false, "apple" already exists with value 5, no change is made.
//   addedNew := ht.AddOK("banana", 3) // addedNew is true, "banana" is added with value 3.
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
//   ht := make(hashtable.Hashtable[string, int])
//   ht.Add("apple", 5)
//   ht.Add("banana", 3)
//   ht.Delete("apple") // Removes the key-value pair with key "apple" from the hashtable.
func (hashtable *Hashtable[K, V]) Delete(key K) *Hashtable[K, V] {
	delete(*hashtable, key)
	return hashtable
}

// DeleteFunc removes key-value pairs from the hashtable based on the evaluation performed by the provided function.
// For each key-value pair in the hashtable, the function fn is called with the key and value.
// If the function returns true, the key-value pair is removed from the hashtable.
//
// Example:
//   ht := make(hashtable.Hashtable[string, int])
//   ht.Add("apple", 5)
//   ht.Add("orange", 7)
//   ht.Add("kiwi", 6)
//
//   // Delete key-value pairs where the value is less than 7.
//   ht.DeleteFunc(func(key string, value int) bool {
//       return value < 7
//   }) // Removes the key-value pair with key "kiwi" from the hashtable.
func (hashtable *Hashtable[K, V]) DeleteFunc(fn func(key K, value V) bool) *Hashtable[K, V] {
	for key, value := range *hashtable {
		if fn(key, value) {
			hashtable.Delete(key)
		}
	}
	return hashtable
}

// DeleteMany removes multiple key-value pairs from the hashtable based on the provided keys.
// If a key doesn't exist in the hashtable, it is ignored.
//
// Example:
//   ht := make(hashtable.Hashtable[string, int])
//   ht.Add("apple", 5)
//   ht.Add("orange", 7)
//   ht.Add("kiwi", 6)
//   ht.DeleteMany("apple", "kiwi") // Removes key-value pairs with keys "apple" and "kiwi" from the hashtable.
func (hashtable *Hashtable[K, V]) DeleteMany(keys ...K) *Hashtable[K, V] {
	for _, key := range keys {
		hashtable.Delete(key)
	}
	return hashtable
}

// Each iterates over the key-value pairs in the hashtable and applies a function to each pair.
//
// Example:
//   ht := make(hashtable.Hashtable[string, int])
//   ht.Add("apple", 5)
//   ht.Add("banana", 3)
//   ht.Add("cherry", 8)
//
//   // Function to print all key-value pairs.
//   printKeyValue := func(key string, value int) {
//       fmt.Println(key, value)
//   }
//
//   // Iterate over the hashtable and print all key-value pairs.
//   ht.Each(printKeyValue)
//   // Output: "apple 5", "banana 3", "cherry 8"
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
//   ht := make(hashtable.Hashtable[string, int])
//   ht.Add("apple", 5)
//   ht.Add("banana", 3)
//   ht.Add("cherry", 8)
//
//   // Function to print key-value pairs until finding "banana".
//   stopPrinting := ht.EachBreak(func(key string, value int) bool {
//       fmt.Println(key, value)
//       return key != "banana" // Continue printing until "banana" is encountered.
//   })
//   // Output: "apple 5", "banana 3"
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
//   ht := make(hashtable.Hashtable[string, int])
//   ht.Add("apple", 5)
//   ht.Add("banana", 3)
//   ht.Add("cherry", 8)
//
//   // Function to print each key.
//   printKey := func(key string) {
//       fmt.Println(key)
//   }
//
//   // Iterate over the hashtable and print each key.
//   ht.EachKey(printKey)
//   // Output: "apple", "banana", "cherry"
func (hashtable *Hashtable[K, V]) EachKey(fn func(key K)) *Hashtable[K, V] {
	return hashtable.Each(func(key K, _ V) {
		fn(key)
	})
}

// EachKeyBreak iterates over the keys in the hashtable and applies a function to each key. It allows breaking the iteration early if the provided function returns false.
//
// Example:
//   ht := make(hashtable.Hashtable[string, int])
//   ht.Add("apple", 5)
//   ht.Add("banana", 3)
//   ht.Add("cherry", 8)
//
//   // Function to print each key and break the iteration if the key is "banana".
//   printAndBreak := func(key string) bool {
//       fmt.Println(key)
//       return key != "banana"
//   }
//
//   // Iterate over the hashtable keys, print them, and break when "banana" is encountered.
//   ht.EachKeyBreak(printAndBreak)
//   // Output: "apple", "banana"
func (hashtable *Hashtable[K, V]) EachKeyBreak(fn func(key K) bool) *Hashtable[K, V] {
	return hashtable.EachBreak(func(key K, _ V) bool {
		return fn(key)
	})
}

// Get retrieves the value associated with the provided key from the hashtable.
// If the key exists, it returns the associated value and true. Otherwise, it returns the zero value for the value type and false.
//
// Example:
//   ht := make(hashtable.Hashtable[string, int])
//   ht.Add("apple", 5)
//   value, exists := ht.Get("apple") // 5, true
//   value, exists = ht.Get("orange")  // 0, false
func (hashtable *Hashtable[K, V]) Get(key K) (V, bool) {
	value, ok := (*hashtable)[key]
	return value, ok
}

// Has checks if the provided key exists in the hashtable.
// It returns true if the key exists, and false otherwise.
//
// Example:
//   ht := make(hashtable.Hashtable[string, int])
//   ht.Add("apple", 5)
//   exists := ht.Has("apple") // true
//   exists = ht.Has("orange")  // false
func (hashtable *Hashtable[K, V]) Has(key K) bool {
	_, ok := (*hashtable)[key]
	return ok
}
