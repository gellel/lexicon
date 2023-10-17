package hashtable_test

import (
	"fmt"
	"testing"

	"github.com/lindsaygelle/hashtable"
)

// TestHashtable tests Hashtable.
func TestHashtable(t *testing.T) {
	table := hashtable.Hashtable[string, string]{
		"hello": "world"}
	t.Log(table)
}

// TestAdd tests Hashtable.Add.
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

// TestAddOK tests Hashtable.AddOK.
func TestAddOK(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])

	// Test adding a new key-value pair
	added := ht.AddOK("apple", 5)
	if !added {
		t.Error("Expected 'apple' to be added, but it was not.")
	}

	// Check if the key-value pair is added correctly
	value, exists := ht["apple"]
	if !exists || value != 5 {
		t.Errorf("Expected key 'apple' with value '5', but got value '%d'", value)
	}

	// Test adding an existing key
	reAdded := ht.AddOK("apple", 10)
	if reAdded {
		t.Error("Expected 'apple' to not be re-added, but it was.")
	}

	// Check if the value for 'apple' remains unchanged
	value, exists = ht["apple"]
	if !exists || value != 5 {
		t.Errorf("Expected key 'apple' to have value '5' after re-adding attempt, but got value '%d'", value)
	}

	// Test adding another new key-value pair
	addedNew := ht.AddOK("banana", 3)
	if !addedNew {
		t.Error("Expected 'banana' to be added, but it was not.")
	}

	// Check if the key-value pair for 'banana' is added correctly
	value, exists = ht["banana"]
	if !exists || value != 3 {
		t.Errorf("Expected key 'banana' with value '3', but got value '%d'", value)
	}
}

// TestDelete tests Hashtable.Delete.
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

// TestDeleteFunc tests Hashtable.DeleteFunc.
func TestDeleteFunc(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["orange"] = 7
	ht["kiwi"] = 6

	// Delete key-value pairs where the value is 7
	ht.DeleteFunc(func(key string, value int) bool {
		return value == 6
	})

	// Check if the key-value pair with key "kiwi" is removed
	_, exists := ht["kiwi"]
	if exists {
		t.Error("Expected key 'kiwi' to be removed, but it still exists.")
	}

	// Check if other key-value pairs are still present
	_, exists = ht["apple"]
	if !exists {
		t.Error("Expected key 'apple' to be present, but it is not.")
	}

	_, exists = ht["orange"]
	if !exists {
		t.Error("Expected key 'orange' to be present, but it is not.")
	}
}

// TestDeleteMany tests Hashtable.DeleteMany.
func TestDeleteMany(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["orange"] = 7
	ht["kiwi"] = 6

	// Delete key-value pairs with keys "apple" and "kiwi"
	ht.DeleteMany("apple", "kiwi")

	// Check if the key-value pair with key "apple" is removed
	_, exists := ht["apple"]
	if exists {
		t.Error("Expected key 'apple' to be removed, but it still exists.")
	}

	// Check if the key-value pair with key "kiwi" is removed
	_, exists = ht["kiwi"]
	if exists {
		t.Error("Expected key 'kiwi' to be removed, but it still exists.")
	}

	// Check if other key-value pairs are still present
	_, exists = ht["orange"]
	if !exists {
		t.Error("Expected key 'orange' to be present, but it is not.")
	}
}

// TestEach tests Hashtable.Each.
func TestEach(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Function to print all key-value pairs.
	var printedKeys []string
	printKeyValue := func(key string, value int) {
		printedKeys = append(printedKeys, key)
		fmt.Println(key, value)
	}

	// Iterate over the hashtable and print all key-value pairs.
	ht.Each(printKeyValue)
	// Output: "apple 5", "banana 3", "cherry 8"

	// Check if the printed keys are correct
	expectedKeys := []string{"apple", "banana", "cherry"}
	for i, key := range expectedKeys {
		if printedKeys[i] != key {
			t.Errorf("Expected key %s, but got %s", key, printedKeys[i])
		}
	}

	// Check if other keys are not printed
	if len(printedKeys) > 3 {
		t.Errorf("Expected only 'apple', 'banana', and 'cherry' to be printed, but got more keys.")
	}
}

// TestEachBreak tests Hashtable.EachBreak.
func TestEachBreak(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Function to print key-value pairs until finding "banana".
	var printedKeys []string
	ht.EachBreak(func(key string, value int) bool {
		printedKeys = append(printedKeys, key)
		fmt.Println(key, value)
		return key != "banana" // Continue printing until "banana" is encountered.
	})
	// Output: "apple 5", "banana 3"

	// Check if the printed keys are correct
	expectedKeys := []string{"apple", "banana"}
	for i, key := range expectedKeys {
		if printedKeys[i] != key {
			t.Errorf("Expected key %s, but got %s", key, printedKeys[i])
		}
	}

	// Check if other keys are not printed
	if len(printedKeys) > 2 {
		t.Errorf("Expected only 'apple' and 'banana' to be printed, but got more keys.")
	}
}

// TestEachKey tests Hashtable.EachKey.
func TestEachKey(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Function to collect each key in a slice.
	var collectedKeys []string
	collectKey := func(key string) {
		collectedKeys = append(collectedKeys, key)
	}

	// Iterate over the hashtable and collect each key.
	ht.EachKey(collectKey)

	// Check if the collected keys are correct
	expectedKeys := []string{"apple", "banana", "cherry"}
	if len(collectedKeys) != len(expectedKeys) {
		t.Errorf("Expected %d keys, but got %d", len(expectedKeys), len(collectedKeys))
	}
	for i, key := range expectedKeys {
		if collectedKeys[i] != key {
			t.Errorf("Expected key %s, but got %s", key, collectedKeys[i])
		}
	}
}

func TestEachKeyBreak(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Function to print each key and break the iteration if the key is "banana".
	var collectedKeys []string
	printAndBreak := func(key string) bool {
		collectedKeys = append(collectedKeys, key)
		fmt.Println(key)
		return key != "banana"
	}

	// Iterate over the hashtable keys, print them, and break when "banana" is encountered.
	ht.EachKeyBreak(printAndBreak)

	// Check if the collected keys and the printed output are correct.
	expectedKeys := []string{"apple", "banana"}
	if len(collectedKeys) != len(expectedKeys) {
		t.Errorf("Expected %d keys, but got %d", len(expectedKeys), len(collectedKeys))
	}
	for i, key := range expectedKeys {
		if collectedKeys[i] != key {
			t.Errorf("Expected key %s, but got %s", key, collectedKeys[i])
		}
	}
}

// TestGet tests Hashtable.Get.

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
