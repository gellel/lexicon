package hashtable_test

import (
	"sort"
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

	// Add key-value pairs to the hashtable.
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Define a map to store the expected output.
	expectedOutput := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 8,
	}

	// Define a function to compare the actual output with the expected output.
	printKeyValue := func(key string, value int) {
		if expected, ok := expectedOutput[key]; ok {
			if value != expected {
				t.Errorf("Expected %s: %d, but got %d", key, expected, value)
			}
			delete(expectedOutput, key)
		} else {
			t.Errorf("Unexpected key: %s", key)
		}
	}

	// Call the Each function with the printKeyValue function.
	ht.Each(printKeyValue)

	// Check if all expected keys have been processed.
	if len(expectedOutput) > 0 {
		t.Errorf("Not all keys were processed: %v", expectedOutput)
	}
}

// TestEachBreak tests Hashtable.EachBreak.
func TestEachBreak(t *testing.T) {
	var stopPrinting string
	ht := make(hashtable.Hashtable[string, int])

	// Add key-value pairs to the hashtable.
	ht.Add("apple", 5)
	ht.Add("banana", 3)
	ht.Add("cherry", 8)

	// Define a function to stop iteration when key is "banana".
	ht.EachBreak(func(key string, value int) bool {
		t.Logf("%s %d", key, value)
		stopPrinting = key
		return key != "banana"
	})

	// Ensure that iteration stopped at "banana".
	if stopPrinting != "banana" {
		t.Errorf("Iteration did not stop at 'banana'.")
	}
}

// TestEachKey tests Hashtable.EachKey.
func TestEachKey(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])

	// Add key-value pairs to the hashtable.
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Define a function to print each key.
	var printedKeys []string
	printKey := func(key string) {
		printedKeys = append(printedKeys, key)
	}

	// Iterate over the keys and print each key.
	ht.EachKey(printKey)

	// Sort the printed values for consistent comparison.
	sort.Strings(printedKeys)

	// Expected output: "apple", "banana", "cherry"
	expectedKeys := []string{"apple", "banana", "cherry"}
	for i, key := range printedKeys {
		if key != expectedKeys[i] {
			t.Errorf("Expected key %s at index %d, but got %s", expectedKeys[i], i, key)
		}
	}
}

// TestEachKeyBreak tests Hashtable.EachKeyBreak.
func TestEachKeyBreak(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])

	// Add key-value pairs to the hashtable.
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Define a function to print each key and break if the key is "banana".
	var printedKeys []string
	printAndBreak := func(key string) bool {
		printedKeys = append(printedKeys, key)
		return key != "banana"
	}

	// Iterate over the keys and print each key, breaking if the key is "banana".
	ht.EachKeyBreak(printAndBreak)

	// Sort the printed values for consistent comparison.
	sort.Strings(printedKeys)

	// Expected output: "apple", "banana"
	expectedKeys := []string{"apple", "banana"}
	for i, key := range printedKeys {
		if key != expectedKeys[i] {
			t.Errorf("Expected key %s at index %d, but got %s", expectedKeys[i], i, key)
		}
	}
}

// TestEachValue tests Hashtable.EachValue.
func TestEachValue(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])

	// Add key-value pairs to the hashtable.
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Define a function to print each value.
	var printedValues []int
	printValue := func(value int) {
		printedValues = append(printedValues, value)
	}

	// Iterate over the hashtable values and print them.
	ht.EachValue(printValue)

	// Sort the printed values for consistent comparison.
	sort.Ints(printedValues)

	// Expected output: 3, 5, 8
	expectedValues := []int{3, 5, 8}

	if len(printedValues) != len(expectedValues) {
		t.Errorf("Expected %d values, but got %d", len(expectedValues), len(printedValues))
		return
	}

	for i, value := range printedValues {
		if value != expectedValues[i] {
			t.Errorf("Expected value %d at index %d, but got %d", expectedValues[i], i, value)
		}
	}
}

// TestEachValueBreak tests Hashtable.EachValueBreak.

func TestEachValueBreak(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])

	// Add key-value pairs to the hashtable.
	ht.Add("apple", 5)
	ht.Add("banana", 3)
	ht.Add("cherry", 8)

	// Sort the keys for consistent iteration order.
	keys := make([]string, 0, len(ht))
	for key := range ht {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Define a function to process each value. It returns false to break the iteration if the value is 3.
	var processedValues []int
	processValue := func(value int) bool {
		processedValues = append(processedValues, value)
		return value != 3
	}

	// Iterate over the hashtable values and process them until the value is 3.
	for _, key := range keys {
		value, _ := ht.Get(key)
		if !processValue(value) {
			break
		}
	}

	// Expected output: 5, 3
	expectedValues := []int{5, 3}
	for i, value := range processedValues {
		if value != expectedValues[i] {
			t.Errorf("Expected value %d at index %d, but got %d", expectedValues[i], i, value)
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
