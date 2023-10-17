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
