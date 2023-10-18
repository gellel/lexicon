package hashtable_test

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/lindsaygelle/hashtable"
	"github.com/lindsaygelle/slice"
)

// TestHashtable tests Hashtable.
func TestHashtable(t *testing.T) {
	table := hashtable.Hashtable[string, string]{
		"hello": "world"}
	t.Log(table)
}

// TestAdd tests Hashtable.Add.
func TestAdd(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])

	// Add key-value pairs to the hashtable.
	ht.Add("apple", 5)
	ht.Add("banana", 3)
	ht.Add("cherry", 8)

	// Verify that the key-value pairs have been added correctly.
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

	// Update the value associated with the "banana" key.
	ht.Add("banana", 10)

	// Verify that the value has been updated correctly.
	if ht["banana"] != 10 {
		t.Fatalf("Expected value for key 'banana' to be updated to 10, but got %d", ht["banana"])
	}
}

// TestAddFunc tests Hashtable.AddFunc.
func TestAddFunc(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])

	// Test case 1: Add key-value pairs where the value is greater than 1.
	ht.AddFunc([]map[string]int{{"apple": 1, "orange": 2, "banana": 3}}, func(key string, value int) bool {
		return value > 1
	})

	// Check if expected key-value pairs are added.
	expected := map[string]int{"orange": 2, "banana": 3}
	for key, expectedValue := range expected {
		value, exists := ht.Get(key)
		if !exists || value != expectedValue {
			t.Fatalf("Expected key '%s' with value '%d', but got value '%d'", key, expectedValue, value)
		}
	}

	// Test case 2: Add all key-value pairs without any validation.
	ht.AddFunc([]map[string]int{{"kiwi": 0, "pear": 4}}, func(key string, value int) bool {
		return true
	})

	// Check if all key-value pairs are added.
	allValues := map[string]int{"orange": 2, "banana": 3, "kiwi": 0, "pear": 4}
	for key, expectedValue := range allValues {
		value, exists := ht.Get(key)
		if !exists || value != expectedValue {
			t.Fatalf("Expected key '%s' with value '%d', but got value '%d'", key, expectedValue, value)
		}
	}
}

// TestAddLength tests Hashtable.AddLength.
func TestAddLength(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])

	// Add key-value pairs to the hashtable and get the length.
	length := ht.AddLength("apple", 5)

	// Expected length after adding the first key-value pair: 1
	expectedLength := 1

	// Verify that the obtained length matches the expected length.
	if length != expectedLength {
		t.Fatalf("Expected length: %d, but got: %d", expectedLength, length)
	}

	// Add another key-value pair and get the updated length.
	length = ht.AddLength("banana", 3)

	// Expected length after adding the second key-value pair: 2
	expectedLength = 2

	// Verify that the obtained length matches the expected length.
	if length != expectedLength {
		t.Fatalf("Expected length: %d, but got: %d", expectedLength, length)
	}

	// Update an existing key-value pair and get the same length.
	length = ht.AddLength("apple", 10)

	// Length should remain 2 after updating the existing key "apple".
	// Expected length: 2
	expectedLength = 2

	// Verify that the obtained length matches the expected length.
	if length != expectedLength {
		t.Fatalf("Expected length: %d, but got: %d", expectedLength, length)
	}
}

// TestAddMany tests Hashtable.AddMany.

func TestAddMany(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])

	// Test case 1.
	ht.AddMany(map[string]int{"orange": 7, "grape": 4})
	expected1 := 7
	if val, ok := ht["orange"]; !ok || val != expected1 {
		t.Fatalf("Expected %d, but got %d for key 'orange'", expected1, val)
	}
	expected2 := 4
	if val, ok := ht["grape"]; !ok || val != expected2 {
		t.Fatalf("Expected %d, but got %d for key 'grape'", expected2, val)
	}

	// Test case 2.
	ht.AddMany(map[string]int{"kiwi": 6, "pear": 9})
	expected3 := 6
	if val, ok := ht["kiwi"]; !ok || val != expected3 {
		t.Fatalf("Expected %d, but got %d for key 'kiwi'", expected3, val)
	}
	expected4 := 9
	if val, ok := ht["pear"]; !ok || val != expected4 {
		t.Fatalf("Expected %d, but got %d for key 'pear'", expected4, val)
	}
}

// TestAddManyOK tests Hashtable.AddManyOK.
func TestAddManyOK(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])

	// Attempt to add multiple key-value pairs and get the results indicating success.
	results := ht.AddManyOK(
		map[string]int{"apple": 5, "banana": 3, "cherry": 8},
	)

	// Expected results: [true, true, true] indicating successful insertions for "apple", "banana" and "cherry".
	expectedResults := []bool{true, true, true}

	// Verify that the obtained results match the expected results.
	for i, result := range *results {
		if result != expectedResults[i] {
			t.Errorf("Expected result: %v, but got: %v", expectedResults[i], result)
		}
	}

	// Attempt to add multiple key-value pairs and get the results indicating success.
	results = ht.AddManyOK(
		map[string]int{"apple": 5, "banana": 3, "cherry": 8},
	)

	// Expected results: [false, false, false] indicating unsuccessful insertions for "apple", "banana" and "cherry" due to existing key.
	expectedResults = []bool{false, false, false}

	// Verify that the obtained results match the expected results.
	for i, result := range *results {
		if result != expectedResults[i] {
			t.Errorf("Expected result: %v, but got: %v", expectedResults[i], result)
		}
	}
}

// TestAddOK tests Hashtable.AddOK.
func TestAddOK(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])

	// Test adding a new key-value pair.
	added := ht.AddOK("apple", 5)
	if !added {
		t.Error("Expected 'apple' to be added, but it was not.")
	}

	// Check if the key-value pair is added correctly.
	value, exists := ht["apple"]
	if !exists || value != 5 {
		t.Fatalf("Expected key 'apple' with value '5', but got value '%d'", value)
	}

	// Test adding an existing key.
	reAdded := ht.AddOK("apple", 10)
	if reAdded {
		t.Error("Expected 'apple' to not be re-added, but it was.")
	}

	// Check if the value for 'apple' remains unchanged.
	value, exists = ht["apple"]
	if !exists || value != 5 {
		t.Fatalf("Expected key 'apple' to have value '5' after re-adding attempt, but got value '%d'", value)
	}

	// Test adding another new key-value pair.
	addedNew := ht.AddOK("banana", 3)
	if !addedNew {
		t.Error("Expected 'banana' to be added, but it was not.")
	}

	// Check if the key-value pair for 'banana' is added correctly.
	value, exists = ht["banana"]
	if !exists || value != 3 {
		t.Fatalf("Expected key 'banana' with value '3', but got value '%d'", value)
	}
}

// TestDelete tests Hashtable.Delete.
func TestDelete(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3

	// Test case 1: Delete an existing key.
	ht.Delete("apple")
	if _, ok := ht["apple"]; ok {
		t.Fatalf("Expected key 'apple' to be deleted, but it still exists in the hashtable")
	}

	// Test case 2: Delete a non-existing key.
	ht.Delete("nonexistent")
	if _, ok := ht["nonexistent"]; ok {
		t.Fatalf("Expected key 'nonexistent' to not exist, but it was found in the hashtable")
	}

	// Test case 3: Delete a key after adding it again.
	ht["apple"] = 10
	ht.Delete("apple")
	if _, ok := ht["apple"]; ok {
		t.Fatalf("Expected key 'apple' to be deleted, but it still exists in the hashtable")
	}
}

// TestDeleteFunc tests Hashtable.DeleteFunc.
func TestDeleteFunc(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["orange"] = 7
	ht["kiwi"] = 6

	// Delete key-value pairs where the value is 7.
	ht.DeleteFunc(func(key string, value int) bool {
		return value == 6
	})

	// Check if the key-value pair with key "kiwi" is removed.
	_, exists := ht["kiwi"]
	if exists {
		t.Error("Expected key 'kiwi' to be removed, but it still exists.")
	}

	// Check if other key-value pairs are still present.
	_, exists = ht["apple"]
	if !exists {
		t.Error("Expected key 'apple' to be present, but it is not.")
	}

	_, exists = ht["orange"]
	if !exists {
		t.Error("Expected key 'orange' to be present, but it is not.")
	}
}

// TestDeleteLength tests Hashtable.DeleteLength.
func TestDeleteLength(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])

	// Add key-value pairs to the hashtable and get the initial length.
	ht["apple"] = 5
	ht["banana"] = 3
	initialLength := len(ht)

	// Delete an existing key from the hashtable and get the updated length.
	lengthAfterDelete := ht.DeleteLength("apple")

	// Expected length after deleting "apple": initial length - 1
	expectedLength := initialLength - 1

	// Verify that the obtained length matches the expected length.
	if lengthAfterDelete != expectedLength {
		t.Fatalf("Expected length: %d, but got: %d", expectedLength, lengthAfterDelete)
	}

	// Attempt to delete a non-existing key from the hashtable.
	lengthAfterNonExistingDelete := ht.DeleteLength("grape")

	// Length should remain the same after attempting to delete a non-existing key.
	// Expected length: initial length
	expectedLengthNonExisting := len(ht)

	// Verify that the obtained length matches the expected length.
	if lengthAfterNonExistingDelete != expectedLengthNonExisting {
		t.Fatalf("Expected length: %d, but got: %d", expectedLengthNonExisting, lengthAfterNonExistingDelete)
	}
}

// TestDeleteMany tests Hashtable.DeleteMany.
func TestDeleteMany(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["orange"] = 7
	ht["kiwi"] = 6

	// Delete key-value pairs with keys "apple" and "kiwi".
	ht.DeleteMany("apple", "kiwi")

	// Check if the key-value pair with key "apple" is removed.
	_, exists := ht["apple"]
	if exists {
		t.Error("Expected key 'apple' to be removed, but it still exists.")
	}

	// Check if the key-value pair with key "kiwi" is removed.
	_, exists = ht["kiwi"]
	if exists {
		t.Error("Expected key 'kiwi' to be removed, but it still exists.")
	}

	// Check if other key-value pairs are still present.
	_, exists = ht["orange"]
	if !exists {
		t.Error("Expected key 'orange' to be present, but it is not.")
	}
}

// TestDeleteManyValues tests Hashtable.DeleteManyValues.
func TestDeleteManyValues(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])
	ht.Add("apple", 5)
	ht.Add("banana", 3)
	ht.Add("cherry", 8)

	// Delete key-value pairs where the value is 3 or 8.
	ht.DeleteManyValues(3, 8)

	// Verify that the hashtable only contains the expected key-value pair.
	expected := hashtable.Hashtable[string, int]{"apple": 5}
	if !reflect.DeepEqual(ht, expected) {
		t.Fatalf("Expected hashtable: %v, but got: %v", expected, ht)
	}
}

// TestDeleteManyOK tests Hashtable.DeleteManyOK.
func TestDeleteManyOK(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])

	// Add key-value pairs to the hashtable.
	ht.Add("apple", 5)
	ht.Add("banana", 3)

	// Specify keys to delete.
	keysToDelete := []string{"apple", "grape"}

	// Attempt to delete keys and check if deletion is successful.
	results := ht.DeleteManyOK(keysToDelete...)

	expectedResults := []bool{true, true} // Expected results for "apple" (exists) and "grape" (does not exist)

	// Check if results match the expected results.
	for i, result := range *results {
		if result != expectedResults[i] {
			t.Errorf("Expected deletion of key %s to be %v but got %v", keysToDelete[i], expectedResults[i], result)
		}
	}
}

// TestDeleteOK tests Hashtable.DeleteOK.
func TestDeleteOK(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])

	// Add key-value pairs to the hashtable.
	ht.Add("apple", 5)
	ht.Add("banana", 3)

	// Delete keys and check if deletion is successful.
	deleted := ht.DeleteOK("apple")
	if !deleted {
		t.Errorf("Expected deletion of 'apple' to be successful")
	}

	// Attempt to delete a key that does not exist.
	notDeleted := ht.DeleteOK("grape")
	if !notDeleted {
		t.Errorf("Expected deletion of 'grape' to be successful because the key does not exist")
	}

	// Attempt to delete a key that has already been deleted.
	alreadyDeleted := ht.DeleteOK("apple")
	if !alreadyDeleted {
		t.Errorf("Expected deletion of 'apple' to be successful even though it was already deleted")
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
				t.Fatalf("Expected %s: %d, but got %d", key, expected, value)
			}
			delete(expectedOutput, key)
		} else {
			t.Fatalf("Unexpected key: %s", key)
		}
	}

	// Call the Each function with the printKeyValue function.
	ht.Each(printKeyValue)

	// Check if all expected keys have been processed.
	if len(expectedOutput) > 0 {
		t.Fatalf("Not all keys were processed: %v", expectedOutput)
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
		t.Fatalf("Iteration did not stop at 'banana'.")
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
			t.Fatalf("Expected key %s at index %d, but got %s", expectedKeys[i], i, key)
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

	var keyToBreak string
	ht.EachBreak(func(key string, value int) bool {
		keyToBreak = key
		return key != "banana"
	})

	if keyToBreak != "banana" {
		t.Fatalf("Expect keyToBreak to equal 'banana', but got %s", keyToBreak)
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
		t.Fatalf("Expected %d values, but got %d", len(expectedValues), len(printedValues))
		return
	}

	for i, value := range printedValues {
		if value != expectedValues[i] {
			t.Fatalf("Expected value %d at index %d, but got %d", expectedValues[i], i, value)
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

	keys := make([]string, 0, len(ht))
	for key := range ht {
		keys = append(keys, key)
	}

	// Sort the keys for consistent iteration order.
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

	// Expected output: 5, 3.
	expectedValues := []int{5, 3}
	for i, value := range processedValues {
		if value != expectedValues[i] {
			t.Fatalf("Expected value %d at index %d, but got %d", expectedValues[i], i, value)
		}
	}
}

// TestGet tests Hashtable.Get.

func TestGet(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3

	// Test case 1: Get an existing key.
	value, exists := ht.Get("apple")
	if !exists {
		t.Fatalf("Expected key 'apple' to exist, but it was not found in the hashtable")
	}
	if value != 5 {
		t.Fatalf("Expected value for key 'apple' to be 5, but got %d", value)
	}

	// Test case 2: Get a non-existing key.
	value, exists = ht.Get("orange")
	if exists {
		t.Fatalf("Expected key 'orange' to not exist, but it was found in the hashtable with value %d", value)
	}
	if value != 0 {
		t.Fatalf("Expected default value for non-existing key 'orange' to be 0, but got %d", value)
	}
}

// TestGetMany tests Hashtable.GetMany.
func TestGetMany(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Get values for specific keys.
	values := ht.GetMany("apple", "banana", "orange")

	// Sort the keys for consistent iteration order.
	sort.Ints(*values)

	// The expected values slice: {5, 3}.
	expectedValues := &slice.Slice[int]{5, 3}

	// Sort the keys for consistent iteration order.
	sort.Ints(*expectedValues)

	// Verify that the obtained values match the expected values.
	if values == expectedValues {
		t.Fatalf("Expected values: %v, but got: %v", expectedValues, values)
	}
}

// TestHas tests Hashtable.Has.

func TestHas(t *testing.T) {
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3

	// Test case 1: Key exists in the hashtable.
	if !ht.Has("apple") {
		t.Fatalf("Expected key 'apple' to exist, but it was not found in the hashtable")
	}

	// Test case 2: Key does not exist in the hashtable.
	if ht.Has("orange") {
		t.Fatalf("Expected key 'orange' to not exist, but it was found in the hashtable")
	}
}

// TestHasMany test Hashtable.HasMany.
func TestHasMany(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Keys to check existence.
	keysToCheck := []string{"apple", "orange", "banana"}

	// Check the existence of multiple keys.
	results := ht.HasMany(keysToCheck...)

	// The expected boolean slice: {true, false, true}
	expectedResults := &slice.Slice[bool]{true, false, true}

	// Verify that the obtained results match the expected results.
	if !reflect.DeepEqual(results, expectedResults) {
		t.Fatalf("Expected results: %v, but got: %v", expectedResults, results)
	}
}

// TestKeys tests Hashtable.Keys.
func TestKeys(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Get all keys from the hashtable.
	keys := ht.Keys()

	// Sort the keys for consistent iteration order.
	sort.Strings(*keys)

	// The expected keys slice: {"apple", "banana", "cherry"}
	expectedKeys := &slice.Slice[string]{"apple", "banana", "cherry"}

	// Sort the keys for consistent iteration order.
	sort.Strings(*expectedKeys)

	// Verify that the obtained keys match the expected keys.
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Fatalf("Expected keys: %v, but got: %v", expectedKeys, keys)
	}
}

func TestKeysFunc(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Get keys from the hashtable where the key length is greater than 5.
	keys := ht.KeysFunc(func(key string) bool {
		return strings.HasPrefix(key, "b")
	})

	// The expected keys slice: {"banana"}
	expectedKeys := &slice.Slice[string]{"banana"}

	// Verify that the obtained keys match the expected keys.
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Fatalf("Expected keys: %v, but got: %v", expectedKeys, keys)
	}
}

func TestLength(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Get the length of the hashtable.
	length := ht.Length()

	// Expected length: 3
	expectedLength := 3

	// Verify that the obtained length matches the expected length.
	if length != expectedLength {
		t.Fatalf("Expected length: %d, but got: %d", expectedLength, length)
	}
}
