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
