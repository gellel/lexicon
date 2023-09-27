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
