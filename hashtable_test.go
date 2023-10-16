package hashtable_test

import (
	"testing"

	"github.com/lindsaygelle/hashtable"
)

// TestHashtable tests Hashtable.
func TestHashtable(t *testing.T) {
	table := hashtable.Hashtable[string, string]{
		"hello": "world"}
	t.Log(table)
}

// TestAdd tests Hashtable.Add
func TestAdd(t *testing.T) {
	// Create a new hashtable
	ht := make(hashtable.Hashtable[string, int])

	// Add key-value pairs to the hashtable
	ht.Add("apple", 5)
	ht.Add("banana", 3)
	ht.Add("cherry", 8)

	// Verify that the key-value pairs have been added correctly
	expected := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 8,
	}

	for key, expectedValue := range expected {
		actualValue, ok := ht[key]
		if !ok {
			t.Fatalf("Key %s not found in the hashtable", key)
		} else if actualValue != expectedValue {
			t.Fatalf("Expected value for key %s is %d, but got %d", key, expectedValue, actualValue)
		}
	}

	// Update the value associated with the "banana" key
	ht.Add("banana", 10)

	// Verify that the value has been updated correctly
	if ht["banana"] != 10 {
		t.Fatalf("Expected value for key 'banana' to be updated to 10, but got %d", ht["banana"])
	}
}

// TestAddFunc tests Hashtable.AddFunc.
func TestAddFunc(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])

	// Test case 1: Add key-value pairs where the value is greater than 1
	ht.AddFunc([]map[string]int{{"apple": 1, "orange": 2, "banana": 3}}, func(key string, value int) bool {
		return value > 1
	})

	// Check if expected key-value pairs are added
	expected := map[string]int{"orange": 2, "banana": 3}
	for key, expectedValue := range expected {
		value, exists := ht.Get(key)
		if !exists || value != expectedValue {
			t.Errorf("Expected key '%s' with value '%d', but got value '%d'", key, expectedValue, value)
		}
	}

	// Test case 2: Add all key-value pairs without any validation
	ht.AddFunc([]map[string]int{{"kiwi": 0, "pear": 4}}, func(key string, value int) bool {
		return true
	})

	// Check if all key-value pairs are added
	allValues := map[string]int{"orange": 2, "banana": 3, "kiwi": 0, "pear": 4}
	for key, expectedValue := range allValues {
		value, exists := ht.Get(key)
		if !exists || value != expectedValue {
			t.Errorf("Expected key '%s' with value '%d', but got value '%d'", key, expectedValue, value)
		}
	}
}

// TestAddMany tests Hashtable.AddMany

func TestAddMany(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])

	// Test case 1
	ht.AddMany(map[string]int{"orange": 7, "grape": 4})
	expected1 := 7
	if val, ok := ht["orange"]; !ok || val != expected1 {
		t.Errorf("Expected %d, but got %d for key 'orange'", expected1, val)
	}
	expected2 := 4
	if val, ok := ht["grape"]; !ok || val != expected2 {
		t.Errorf("Expected %d, but got %d for key 'grape'", expected2, val)
	}

	// Test case 2
	ht.AddMany(map[string]int{"kiwi": 6, "pear": 9})
	expected3 := 6
	if val, ok := ht["kiwi"]; !ok || val != expected3 {
		t.Errorf("Expected %d, but got %d for key 'kiwi'", expected3, val)
	}
	expected4 := 9
	if val, ok := ht["pear"]; !ok || val != expected4 {
		t.Errorf("Expected %d, but got %d for key 'pear'", expected4, val)
	}
}

// TestDelete tests Hashtable.Delete
func TestDelete(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3

	// Test case 1: Delete an existing key
	ht.Delete("apple")
	if _, ok := ht["apple"]; ok {
		t.Errorf("Expected key 'apple' to be deleted, but it still exists in the hashtable")
	}

	// Test case 2: Delete a non-existing key
	ht.Delete("nonexistent")
	if _, ok := ht["nonexistent"]; ok {
		t.Errorf("Expected key 'nonexistent' to not exist, but it was found in the hashtable")
	}

	// Test case 3: Delete a key after adding it again
	ht["apple"] = 10
	ht.Delete("apple")
	if _, ok := ht["apple"]; ok {
		t.Errorf("Expected key 'apple' to be deleted, but it still exists in the hashtable")
	}
}

// TestGet tests Hashtable.Get

func TestGet(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3

	// Test case 1: Get an existing key
	value, exists := ht.Get("apple")
	if !exists {
		t.Errorf("Expected key 'apple' to exist, but it was not found in the hashtable")
	}
	if value != 5 {
		t.Errorf("Expected value for key 'apple' to be 5, but got %d", value)
	}

	// Test case 2: Get a non-existing key
	value, exists = ht.Get("orange")
	if exists {
		t.Errorf("Expected key 'orange' to not exist, but it was found in the hashtable with value %d", value)
	}
	if value != 0 {
		t.Errorf("Expected default value for non-existing key 'orange' to be 0, but got %d", value)
	}
}

// TestHas tests Hashtable.Has

func TestHas(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3

	// Test case 1: Key exists in the hashtable
	if !ht.Has("apple") {
		t.Errorf("Expected key 'apple' to exist, but it was not found in the hashtable")
	}

	// Test case 2: Key does not exist in the hashtable
	if ht.Has("orange") {
		t.Errorf("Expected key 'orange' to not exist, but it was found in the hashtable")
	}
}
