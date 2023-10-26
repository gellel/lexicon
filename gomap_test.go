package gomap_test

import (
	"math"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/lindsaygelle/gomap"
	"github.com/lindsaygelle/slice"
)

// TestHasnewMapable tests Map.
func TestHasnewMapable(t *testing.T) {
	table := gomap.Map[string, string]{
		"hello": "world"}
	t.Log(table)
}

// TestAdd tests Map.Add.
func TestAdd(t *testing.T) {
	// Test case 1: Add key-value pairs to an empty gomap.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	newMap.Add("banana", 3)

	// Verify the values using Get method.
	value, ok := newMap.Get("apple")
	if !ok || value != 5 {
		t.Errorf("Expected key 'apple' to have value 5, but got %v", value)
	}
	value, ok = newMap.Get("banana")
	if !ok || value != 3 {
		t.Errorf("Expected key 'banana' to have value 3, but got %v", value)
	}

	// Test case 2: Replace existing key-value pair.
	newMap.Add("banana", 10) // Updates the value for the key "banana" to 10.
	value, ok = newMap.Get("banana")
	if !ok || value != 10 {
		t.Errorf("Expected key 'banana' to have updated value 10, but got %v", value)
	}

	// Test case 3: Add a new key-value pair.
	newMap.Add("cherry", 8)
	value, ok = newMap.Get("cherry")
	if !ok || value != 8 {
		t.Errorf("Expected key 'cherry' to have value 8, but got %v", value)
	}

	// Verify that other keys are not affected.
	value, ok = newMap.Get("grape")
	if ok {
		t.Errorf("Expected key 'grape' to be absent, but it was found with value %v", value)
	}
}

// TestAddLength tests Map.AddLength.
func TestAddLength(t *testing.T) {
	// Test case 1: Add key-value pairs to an empty gomap.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	length := newMap.AddLength("apple", 5)
	if length != 1 {
		t.Errorf("Expected length of gomap after adding 'apple' with value 5 to be 1, but got %v", length)
	}

	// Test case 2: Replace an existing key-value pair.
	length = newMap.AddLength("apple", 10)
	if length != 1 {
		t.Errorf("Expected length of gomap after updating 'apple' with value 10 to be 1, but got %v", length)
	}

	// Test case 3: Add a new key-value pair.
	length = newMap.AddLength("banana", 3)
	if length != 2 {
		t.Errorf("Expected length of gomap after adding 'banana' with value 3 to be 2, but got %v", length)
	}

	// Test case 4: Add another new key-value pair.
	length = newMap.AddLength("cherry", 8)
	if length != 3 {
		t.Errorf("Expected length of gomap after adding 'cherry' with value 8 to be 3, but got %v", length)
	}

	// Verify that other keys are not affected.
	value, ok := newMap.Get("grape")
	if ok {
		t.Errorf("Expected key 'grape' to be absent, but it was found with value %v", value)
	}
}

// TestAddMany tests Map.AddMany.

func TestAddMany(t *testing.T) {
	// Test case 1: Add key-value pairs to an empty gomap.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.AddMany(map[string]int{"orange": 7, "grape": 4}, map[string]int{"kiwi": 6, "pear": 9})

	// Verify the added key-value pairs.
	expected := map[string]int{"orange": 7, "grape": 4, "kiwi": 6, "pear": 9}
	for key, expectedValue := range expected {
		value, ok := newMap.Get(key)
		if !ok {
			t.Errorf("Expected key '%s' to be present, but it was not found", key)
		}
		if value != expectedValue {
			t.Errorf("Expected value for key '%s' to be %d, but got %d", key, expectedValue, value)
		}
	}

	// Test case 2: Replace existing key-value pairs and add new ones.
	newMap.AddMany(map[string]int{"orange": 10, "grape": 3}, map[string]int{"apple": 5, "banana": 8})

	// Verify the updated and added key-value pairs.
	expected = map[string]int{"orange": 10, "grape": 3, "kiwi": 6, "pear": 9, "apple": 5, "banana": 8}
	for key, expectedValue := range expected {
		value, ok := newMap.Get(key)
		if !ok {
			t.Errorf("Expected key '%s' to be present, but it was not found", key)
		}
		if value != expectedValue {
			t.Errorf("Expected value for key '%s' to be %d, but got %d", key, expectedValue, value)
		}
	}

	// Test case 3: Add new key-value pairs.
	newMap.AddMany(map[string]int{"watermelon": 12, "cherry": 7})

	// Verify the added key-value pairs.
	expected = map[string]int{"orange": 10, "grape": 3, "kiwi": 6, "pear": 9, "apple": 5, "banana": 8, "watermelon": 12, "cherry": 7}
	for key, expectedValue := range expected {
		value, ok := newMap.Get(key)
		if !ok {
			t.Errorf("Expected key '%s' to be present, but it was not found", key)
		}
		if value != expectedValue {
			t.Errorf("Expected value for key '%s' to be %d, but got %d", key, expectedValue, value)
		}
	}

	// Verify that other keys are not affected.
	_, ok := newMap.Get("pineapple")
	if ok {
		t.Errorf("Expected key 'pineapple' to be absent, but it was found")
	}
}

// TestAddManyFunc tests Map.AddManyFunc.
func TestAddManyFunc(t *testing.T) {
	// Test case 1: Add key-value pairs with values greater than 0 to the gomap.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	values := []map[string]int{{"apple": 5, "orange": -3, "banana": 10}}
	condition := func(i int, key string, value int) bool {
		return value > 0
	}
	newMap.AddManyFunc(values, condition)

	// Verify that only the key-value pairs with values greater than 0 are added to the gomap.
	expected := &gomap.Map[string, int]{"apple": 5, "banana": 10}
	if !reflect.DeepEqual(newMap, expected) {
		t.Errorf("Expected %v, but got %v", expected, newMap)
	}

	// Test case 2: Add all key-value pairs to the gomap.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	values = []map[string]int{{"apple": 5, "orange": -3, "banana": 10}}
	condition = func(i int, key string, value int) bool {
		return true // Add all key-value pairs
	}
	newMap.AddManyFunc(values, condition)

	// Verify that all key-value pairs are added to the gomap.
	expected = &gomap.Map[string, int]{"apple": 5, "orange": -3, "banana": 10}
	if !reflect.DeepEqual(newMap, expected) {
		t.Errorf("Expected %v, but got %v", expected, newMap)
	}
}

// TestAddManyOK tests Map.AddManyOK.
func TestAddManyOK(t *testing.T) {
	// Test case 1: Add key-value pairs to an empty gomap.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	results := newMap.AddManyOK(
		map[string]int{"apple": 5},
		map[string]int{"banana": 3},
		map[string]int{"banana": 10},
		map[string]int{"cherry": 8},
	)

	// Verify the success status of insertions.
	expectedResults := &slice.Slice[bool]{true, true, false, true}
	if !reflect.DeepEqual(results, expectedResults) {
		t.Errorf("Expected insertion results %v, but got %v", expectedResults, *results)
	}

	// Verify the added and updated key-value pairs.
	expected := map[string]int{"apple": 5, "cherry": 8}
	for key, expectedValue := range expected {
		value, ok := newMap.Get(key)
		if !ok {
			t.Errorf("Expected key '%s' to be present, but it was not found", key)
		}
		if value != expectedValue {
			t.Errorf("Expected value for key '%s' to be %d, but got %d", key, expectedValue, value)
		}
	}

	// Test case 2: Add new key-value pairs.
	results = newMap.AddManyOK(
		map[string]int{"watermelon": 12, "pineapple": 7},
	)

	// Verify the success status of insertions.
	expectedResults = &slice.Slice[bool]{true, true}
	if !reflect.DeepEqual(results, expectedResults) {
		t.Errorf("Expected insertion results %v, but got %v", expectedResults, *results)
	}

	// Verify the added key-value pairs.
	expected = map[string]int{"apple": 5, "cherry": 8, "watermelon": 12, "pineapple": 7}
	for key, expectedValue := range expected {
		value, ok := newMap.Get(key)
		if !ok {
			t.Errorf("Expected key '%s' to be present, but it was not found", key)
		}
		if value != expectedValue {
			t.Errorf("Expected value for key '%s' to be %d, but got %d", key, expectedValue, value)
		}
	}
}

// TestAddOK tests Map.AddOK.
func TestAddOK(t *testing.T) {
	newMap := make(gomap.Map[string, int])

	// Test adding a new key-value pair.
	added := newMap.AddOK("apple", 5)
	if !added {
		t.Error("Expected 'apple' to be added, but it was not.")
	}

	// Check if the key-value pair is added correctly.
	value, exists := newMap["apple"]
	if !exists || value != 5 {
		t.Fatalf("Expected key 'apple' with value '5', but got value '%d'", value)
	}

	// Test adding an existing key.
	reAdded := newMap.AddOK("apple", 10)
	if reAdded {
		t.Error("Expected 'apple' to not be re-added, but it was.")
	}

	// Check if the value for 'apple' remains unchanged.
	value, exists = newMap["apple"]
	if !exists || value != 5 {
		t.Fatalf("Expected key 'apple' to have value '5' after re-adding attempt, but got value '%d'", value)
	}

	// Test adding another new key-value pair.
	addedNew := newMap.AddOK("banana", 3)
	if !addedNew {
		t.Error("Expected 'banana' to be added, but it was not.")
	}

	// Check if the key-value pair for 'banana' is added correctly.
	value, exists = newMap["banana"]
	if !exists || value != 3 {
		t.Fatalf("Expected key 'banana' with value '3', but got value '%d'", value)
	}
}

// TestContains tests Map.Contains.
func TestContains(t *testing.T) {
	// Test case 1: Check for a value in an empty gomap.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	key, found := newMap.Contains(5)    // Check if value 5 is in the gomap.

	// Since the gomap is empty, the result should be ("", false).
	if key != "" || found {
		t.Errorf("Expected ('', false), but got (%s, %t)", key, found)
	}

	// Test case 2: Check for a value in a non-empty gomap.
	newMap = &gomap.Map[string, int]{} // Create a new gomap.
	newMap.Add("apple", 5)
	newMap.Add("banana", 10)
	newMap.Add("orange", 5)

	// Check if value 5 is in the gomap.
	key, found = newMap.Contains(5)

	// Since value 5 exists in the gomap, the result should be ("apple", true).
	if !found {
		t.Errorf("Expected ('%s', true), but got (%s, %t)", key, key, found)
	}

	// Test case 3: Check for a value that doesn't exist in the gomap.
	key, found = newMap.Contains(15) // Check if value 15 is in the gomap.

	// Since value 15 doesn't exist in the gomap, the result should be ("", false).
	if key != "" || found {
		t.Errorf("Expected ('', false), but got (%s, %t)", key, found)
	}
}

// TestDelete tests Map.Delete.
func TestDelete(t *testing.T) {
	newMap := make(gomap.Map[string, int])
	newMap["apple"] = 5
	newMap["banana"] = 3

	// Test case 1: Delete an existing key.
	newMap.Delete("apple")
	if _, ok := newMap["apple"]; ok {
		t.Fatalf("Expected key 'apple' to be deleted, but it still exists in the gomap")
	}

	// Test case 2: Delete a non-existing key.
	newMap.Delete("nonexistent")
	if _, ok := newMap["nonexistent"]; ok {
		t.Fatalf("Expected key 'nonexistent' to not exist, but it was found in the gomap")
	}

	// Test case 3: Delete a key after adding it again.
	newMap["apple"] = 10
	newMap.Delete("apple")
	if _, ok := newMap["apple"]; ok {
		t.Fatalf("Expected key 'apple' to be deleted, but it still exists in the gomap")
	}
}

// TestDeleteLength tests Map.DeleteLength.
func TestDeleteLength(t *testing.T) {
	// Create a new gomap.
	newMap := make(gomap.Map[string, int])

	// Add key-value pairs to the gomap and get the initial length.
	newMap["apple"] = 5
	newMap["banana"] = 3
	initialLength := len(newMap)

	// Delete an existing key from the gomap and get the updated length.
	lengthAfterDelete := newMap.DeleteLength("apple")

	// Expected length after deleting "apple": initial length - 1.
	expectedLength := initialLength - 1

	// Verify that the obtained length matches the expected length.
	if lengthAfterDelete != expectedLength {
		t.Fatalf("Expected length: %d, but got: %d", expectedLength, lengthAfterDelete)
	}

	// Attempt to delete a non-existing key from the gomap.
	lengthAfterNonExistingDelete := newMap.DeleteLength("grape")

	// Length should remain the same after attempting to delete a non-existing key.
	// Expected length: initial length.
	expectedLengthNonExisting := len(newMap)

	// Verify that the obtained length matches the expected length.
	if lengthAfterNonExistingDelete != expectedLengthNonExisting {
		t.Fatalf("Expected length: %d, but got: %d", expectedLengthNonExisting, lengthAfterNonExistingDelete)
	}
}

// TestDeleteMany tests Map.DeleteMany.
func TestDeleteMany(t *testing.T) {
	// Test case 1: Delete keys from an empty gomap.
	newMap := &gomap.Map[string, int]{}  // Create an empty gomap.
	newMap.DeleteMany("apple", "banana") // Attempt to delete keys "apple" and "banana".

	// The gomap should remain empty.
	if !newMap.IsEmpty() {
		t.Errorf("Expected gomap to be empty, but got non-empty gomap: %v", newMap)
	}

	// Test case 2: Delete keys from a non-empty gomap.
	newMap = &gomap.Map[string, int]{} // Create a new gomap.
	newMap.Add("apple", 5)
	newMap.Add("banana", 10)
	newMap.Add("orange", 3)

	// Delete keys "apple" and "banana".
	newMap.DeleteMany("apple", "banana")

	// Verify that "apple" and "banana" are deleted, and "orange" remains in the gomap.
	expected := &slice.Slice[string]{"orange"}
	result := newMap.Keys()

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected keys %v after deleting 'apple' and 'banana', but got keys %v", expected, result)
	}
}

// TestDeleteManyOK tests Map.DeleteManyOK.
func TestDeleteManyOK(t *testing.T) {
	// Create a new gomap.
	newMap := make(gomap.Map[string, int])

	// Add key-value pairs to the gomap.
	newMap.Add("apple", 5)
	newMap.Add("banana", 3)

	// Specify keys to delete.
	keysToDelete := []string{"apple", "grape"}

	// Attempt to delete keys and check if deletion is successful.
	results := newMap.DeleteManyOK(keysToDelete...)

	expectedResults := []bool{true, true} // Expected results for "apple" (exists) and "grape" (does not exist).

	// Check if results match the expected results.
	for i, result := range *results {
		if result != expectedResults[i] {
			t.Fatalf("Expected deletion of key %s to be %v but got %v", keysToDelete[i], expectedResults[i], result)
		}
	}
}

// TestDeleteManyValues tests Map.DeleteManyValues.
func TestDeleteManyValues(t *testing.T) {
	// Create a new gomap.
	newMap := make(gomap.Map[string, int])
	newMap.Add("apple", 5)
	newMap.Add("banana", 3)
	newMap.Add("cherry", 8)

	// Delete key-value pairs where the value is 3 or 8.
	newMap.DeleteManyValues(3, 8)

	// Verify that the gomap only contains the expected key-value pair.
	expected := gomap.Map[string, int]{"apple": 5}
	if !reflect.DeepEqual(newMap, expected) {
		t.Fatalf("Expected gomap: %v, but got: %v", expected, newMap)
	}
}

// TestDeleteOK tests Map.DeleteOK.
func TestDeleteOK(t *testing.T) {
	// Create a new gomap.
	newMap := make(gomap.Map[string, int])

	// Add key-value pairs to the gomap.
	newMap.Add("apple", 5)
	newMap.Add("banana", 3)

	// Delete keys and check if deletion is successful.
	deleted := newMap.DeleteOK("apple")
	if !deleted {
		t.Fatalf("Expected deletion of 'apple' to be successful")
	}

	// Attempt to delete a key that does not exist.
	notDeleted := newMap.DeleteOK("grape")
	if !notDeleted {
		t.Fatalf("Expected deletion of 'grape' to be successful because the key does not exist")
	}

	// Attempt to delete a key that has already been deleted.
	alreadyDeleted := newMap.DeleteOK("apple")
	if !alreadyDeleted {
		t.Fatalf("Expected deletion of 'apple' to be successful even though it was already deleted")
	}
}

// TestEach tests Map.Each.
func TestEach(t *testing.T) {
	newMap := make(gomap.Map[string, int])

	// Add key-value pairs to the gomap.
	newMap["apple"] = 5
	newMap["banana"] = 3
	newMap["cherry"] = 8

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
	newMap.Each(printKeyValue)

	// Check if all expected keys have been processed.
	if len(expectedOutput) > 0 {
		t.Fatalf("Not all keys were processed: %v", expectedOutput)
	}
}

// TestEachBreak tests Map.EachBreak.
func TestEachBreak(t *testing.T) {
	var stopPrinting string
	newMap := make(gomap.Map[string, int])

	// Add key-value pairs to the gomap.
	newMap.Add("apple", 5)
	newMap.Add("banana", 3)
	newMap.Add("cherry", 8)

	// Define a function to stop iteration when key is "banana".
	newMap.EachBreak(func(key string, value int) bool {
		t.Logf("%s %d", key, value)
		stopPrinting = key
		return key != "banana"
	})

	// Ensure that iteration stopped at "banana".
	if stopPrinting != "banana" {
		t.Fatalf("Iteration did not stop at 'banana'.")
	}
}

// TestEachKey tests Map.EachKey.
func TestEachKey(t *testing.T) {
	newMap := make(gomap.Map[string, int])

	// Add key-value pairs to the gomap.
	newMap["apple"] = 5
	newMap["banana"] = 3
	newMap["cherry"] = 8

	// Define a function to print each key.
	var printedKeys []string
	printKey := func(key string) {
		printedKeys = append(printedKeys, key)
	}

	// Iterate over the keys and print each key.
	newMap.EachKey(printKey)

	// Sort the printed values for consistent comparison.
	sort.Strings(printedKeys)

	// Expected output: "apple", "banana", "cherry".
	expectedKeys := []string{"apple", "banana", "cherry"}
	for i, key := range printedKeys {
		if key != expectedKeys[i] {
			t.Fatalf("Expected key %s at index %d, but got %s", expectedKeys[i], i, key)
		}
	}
}

// TestEachKeyBreak tests Map.EachKeyBreak.
func TestEachKeyBreak(t *testing.T) {
	newMap := make(gomap.Map[string, int])

	// Add key-value pairs to the gomap.
	newMap["apple"] = 5
	newMap["banana"] = 3
	newMap["cherry"] = 8

	var keyToBreak string
	newMap.EachBreak(func(key string, value int) bool {
		keyToBreak = key
		return key != "banana"
	})

	if keyToBreak != "banana" {
		t.Fatalf("Expect keyToBreak to equal 'banana', but got %s", keyToBreak)
	}
}

// TestEachValue tests Map.EachValue.
func TestEachValue(t *testing.T) {
	// Create a new gomap.
	newMap := make(gomap.Map[string, int])

	// Add key-value pairs to the gomap.
	newMap["apple"] = 5
	newMap["banana"] = 3
	newMap["cherry"] = 8

	// Define a function to print each value.
	var printedValues []int
	printValue := func(value int) {
		printedValues = append(printedValues, value)
	}

	// Iterate over the gomap values and print them.
	newMap.EachValue(printValue)

	// Sort the printed values for consistent comparison.
	sort.Ints(printedValues)

	// Expected output: 3, 5, 8.
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

// TestEachValueBreak tests Map.EachValueBreak.

func TestEachValueBreak(t *testing.T) {
	// Create a new gomap.
	newMap := make(gomap.Map[string, int])

	// Add key-value pairs to the gomap.
	newMap.Add("apple", 5)
	newMap.Add("banana", 3)
	newMap.Add("cherry", 8)

	keys := make([]string, 0, len(newMap))
	for key := range newMap {
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

	// Iterate over the gomap values and process them until the value is 3.
	for _, key := range keys {
		value, _ := newMap.Get(key)
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

// TestEmptyInto tests Map.EmptyInto.
func TestEmptyInto(t *testing.T) {
	// Test case 1: Transfer from an empty gomap to another empty gomap.
	newMap1 := &gomap.Map[string, int]{} // Create an empty source gomap.
	newMap2 := &gomap.Map[string, int]{} // Create an empty destination gomap.
	newMap1.EmptyInto(newMap2)           // Transfer from newMap1 to newMap2.

	// Verify that newMap1 is empty.
	if newMap1.Length() != 0 {
		t.Errorf("Expected source gomap to be empty after transfer, but it has %d items", newMap1.Length())
	}

	// Verify that newMap2 is still empty.
	if newMap2.Length() != 0 {
		t.Errorf("Expected destination gomap to be empty after transfer, but it has %d items", newMap2.Length())
	}

	// Test case 2: Transfer from a non-empty gomap to an empty gomap.
	newMap1 = &gomap.Map[string, int]{} // Create an empty source gomap.
	newMap1.Add("apple", 5)
	newMap2 = &gomap.Map[string, int]{} // Create an empty destination gomap.
	newMap1.EmptyInto(newMap2)          // Transfer from newMap1 to newMap2.

	// Verify that newMap1 is empty.
	if newMap1.Length() != 0 {
		t.Errorf("Expected source gomap to be empty after transfer, but it has %d items", newMap1.Length())
	}

	// Verify that newMap2 contains the transferred key-value pair.
	expectedValue := 5
	transferredValue, ok := newMap2.Get("apple")
	if !ok || transferredValue != expectedValue {
		t.Errorf("Expected destination gomap to contain 'apple': %d after transfer, but it contains '%d'", expectedValue, transferredValue)
	}
}

// TestEqual tests Map.Equal.
func TestEqual(t *testing.T) {
	// Test case 1: Compare equal hasnewMapables.
	newMap1 := &gomap.Map[string, int]{} // Create a new gomap.
	newMap1.Add("apple", 5)
	newMap1.Add("orange", 10)

	newMap2 := &gomap.Map[string, int]{} // Create another gomap with similar values.
	newMap2.Add("apple", 5)
	newMap2.Add("orange", 10)

	// Check if the two hasnewMapables are equal.
	equal := newMap1.Equal(newMap2)

	// Since newMap1 and newMap2 have the same key-value pairs, they are considered equal.
	if !equal {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Compare unequal hasnewMapables.
	newMap3 := &gomap.Map[string, int]{} // Create a new gomap.
	newMap3.Add("apple", 5)
	newMap3.Add("orange", 10)

	newMap4 := &gomap.Map[string, int]{} // Create another gomap with different values for "orange".
	newMap4.Add("apple", 5)
	newMap4.Add("orange", 12)

	// Check if the two hasnewMapables are equal.
	equal = newMap3.Equal(newMap4)

	// Since newMap3 and newMap4 have different values for "orange", they are not considered equal.
	if equal {
		t.Errorf("Expected false, but got true")
	}
}

// TestEqualFunc tests Map.EqualFunc.
func TestEqualFunc(t *testing.T) {
	// Custom comparison function to check if two integers are equal when their difference is less than or equal to 1
	compareFunc := func(a, b int) bool {
		return math.Abs(float64(a-b)) <= 1
	}

	// Test case 1: Compare equal hasnewMapables based on the custom comparison function.
	newMap1 := &gomap.Map[string, int]{} // Create a new gomap.
	newMap1.Add("apple", 5)
	newMap1.Add("orange", 10)

	newMap2 := &gomap.Map[string, int]{} // Create another gomap with similar values.
	newMap2.Add("apple", 5)
	newMap2.Add("orange", 11) // The difference between 10 and 11 is within the allowed range according to compareFunc.

	// Check if the two hasnewMapables are equal based on the custom comparison function.
	equal := newMap1.EqualFunc(newMap2, compareFunc)

	// Since the values for "orange" (10 and 11) have a difference of 1, within the allowed range,
	// the hasnewMapables are considered equal according to the custom comparison function.
	if !equal {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Compare unequal hasnewMapables based on the custom comparison function.
	newMap3 := &gomap.Map[string, int]{} // Create a new gomap.
	newMap3.Add("apple", 5)
	newMap3.Add("orange", 10)

	newMap4 := &gomap.Map[string, int]{} // Create another gomap with different values for "orange".
	newMap4.Add("apple", 5)
	newMap4.Add("orange", 12) // The difference between 10 and 12 is greater than the allowed range according to compareFunc.

	// Check if the two hasnewMapables are equal based on the custom comparison function.
	equal = newMap3.EqualFunc(newMap4, compareFunc)

	// Since the difference between 10 and 12 is greater than the allowed range according to compareFunc,
	// the hasnewMapables are not considered equal based on the custom comparison function.
	if equal {
		t.Errorf("Expected false, but got true")
	}
}

// TestEqualLength tests Map.EqualLength.
func TestEqualLength(t *testing.T) {
	// Test case 1: Compare hasnewMapables with equal length.
	newMap1 := &gomap.Map[string, int]{} // Create a new gomap.
	newMap1.Add("apple", 5)
	newMap1.Add("orange", 10)

	newMap2 := &gomap.Map[string, int]{} // Create another gomap with the same number of key-value pairs.
	newMap2.Add("apple", 5)
	newMap2.Add("orange", 7)

	// Check if the two hasnewMapables have equal length.
	equalLength := newMap1.EqualLength(newMap2)

	// Since newMap1 and newMap2 have the same number of key-value pairs, they are considered equal in length.
	if !equalLength {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Compare hasnewMapables with unequal length.
	newMap3 := &gomap.Map[string, int]{} // Create a new gomap.
	newMap3.Add("apple", 5)

	newMap4 := &gomap.Map[string, int]{} // Create another gomap with a different number of key-value pairs.
	newMap4.Add("apple", 5)
	newMap4.Add("orange", 7)

	// Check if the two hasnewMapables have equal length.
	equalLength = newMap3.EqualLength(newMap4)

	// Since newMap3 and newMap4 have different numbers of key-value pairs, they are not considered equal in length.
	if equalLength {
		t.Errorf("Expected false, but got true")
	}
}

// TestFetch tests Map.Fetch.
func TestFetch(t *testing.T) {
	// Test case 1: Fetch value for an existing key.
	newMap := &gomap.Map[string, int]{} // Create a new gomap.
	newMap.Add("apple", 5)
	newMap.Add("orange", 10)

	// Fetch the value associated with the key "apple".
	fetchedValue := newMap.Fetch("apple")

	// Since "apple" is in the gomap, the fetched value should be 5.
	expectedValue := 5
	if fetchedValue != expectedValue {
		t.Errorf("Expected %d, but got %d", expectedValue, fetchedValue)
	}

	// Test case 2: Fetch value for a non-existing key.
	// Fetch the value associated with the key "banana", which is not in the gomap.
	fetchedValue = newMap.Fetch("banana")

	// Since "banana" is not in the gomap, the fetched value should be the zero value for int, which is 0.
	expectedValue = 0
	if fetchedValue != expectedValue {
		t.Errorf("Expected %d, but got %d", expectedValue, fetchedValue)
	}
}

// TestFilter tests Map.Filter.
func TestFilter(t *testing.T) {
	// Test case 1: Filter with an empty gomap and a function that never selects any pairs.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	filterFunc := func(key string, value int) bool {
		return false // Never select any values
	}
	filtered := newMap.Filter(filterFunc)
	expected := &gomap.Map[string, int]{} // Expected empty gomap.
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Expected %v, but got %v", expected, filtered)
	}

	// Test case 2: Filter with a non-empty gomap and a function that never selects any pairs.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	newMap.Add("orange", 10)
	filterFunc = func(key string, value int) bool {
		return false // Never select any values
	}
	filtered = newMap.Filter(filterFunc)
	expected = &gomap.Map[string, int]{} // Expected empty gomap.
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Expected %v, but got %v", expected, filtered)
	}

	// Test case 3: Filter with a non-empty gomap and a function that selects certain pairs.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	newMap.Add("orange", 10)
	newMap.Add("banana", 3)
	filterFunc = func(key string, value int) bool {
		return value > 4 // Select pairs where value is greater than 4
	}
	filtered = newMap.Filter(filterFunc)
	expected = &gomap.Map[string, int]{"apple": 5, "orange": 10} // Expected filtered gomap.
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Expected %v, but got %v", expected, filtered)
	}
}

// TestGet tests Map.Get.

func TestGet(t *testing.T) {
	newMap := make(gomap.Map[string, int])
	newMap["apple"] = 5
	newMap["banana"] = 3

	// Test case 1: Get an existing key.
	value, exists := newMap.Get("apple")
	if !exists {
		t.Fatalf("Expected key 'apple' to exist, but it was not found in the gomap")
	}
	if value != 5 {
		t.Fatalf("Expected value for key 'apple' to be 5, but got %d", value)
	}

	// Test case 2: Get a non-existing key.
	value, exists = newMap.Get("orange")
	if exists {
		t.Fatalf("Expected key 'orange' to not exist, but it was found in the gomap with value %d", value)
	}
	if value != 0 {
		t.Fatalf("Expected default value for non-existing key 'orange' to be 0, but got %d", value)
	}
}

// TestGetMany tests Map.GetMany.
func TestGetMany(t *testing.T) {
	// Create a new gomap.
	newMap := make(gomap.Map[string, int])
	newMap["apple"] = 5
	newMap["banana"] = 3
	newMap["cherry"] = 8

	// Get values for specific keys.
	values := newMap.GetMany("apple", "banana", "orange")

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

// TestHas tests Map.Has.

func TestHas(t *testing.T) {
	newMap := make(gomap.Map[string, int])
	newMap["apple"] = 5
	newMap["banana"] = 3

	// Test case 1: Key exists in the gomap.
	if !newMap.Has("apple") {
		t.Fatalf("Expected key 'apple' to exist, but it was not found in the gomap")
	}

	// Test case 2: Key does not exist in the gomap.
	if newMap.Has("orange") {
		t.Fatalf("Expected key 'orange' to not exist, but it was found in the gomap")
	}
}

// TestHasMany test Map.HasMany.
func TestHasMany(t *testing.T) {
	// Create a new gomap.
	newMap := make(gomap.Map[string, int])
	newMap["apple"] = 5
	newMap["banana"] = 3
	newMap["cherry"] = 8

	// Keys to check existence.
	keysToCheck := []string{"apple", "orange", "banana"}

	// Check the existence of multiple keys.
	results := newMap.HasMany(keysToCheck...)

	// The expected boolean slice: {true, false, true}.
	expectedResults := &slice.Slice[bool]{true, false, true}

	// Verify that the obtained results match the expected results.
	if !reflect.DeepEqual(results, expectedResults) {
		t.Fatalf("Expected results: %v, but got: %v", expectedResults, results)
	}
}

// TestIntersectionFunc tests Map.IntersectionFunc.
func TestIntersectionFunc(t *testing.T) {
	// Test case: Check intersection of two hasnewMapables with common key-value pairs.
	newMap1 := &gomap.Map[string, int]{
		"apple":  5,
		"orange": 8,
	}
	newMap2 := &gomap.Map[string, int]{
		"orange": 8,
		"banana": 6,
	}

	// Condition function to check if values are equal.
	conditionFunc := func(key string, a, b int) bool {
		return a == b
	}

	newHasnewMapable := newMap1.IntersectionFunc(newMap2, conditionFunc)

	expectedHasnewMapable := &gomap.Map[string, int]{
		"orange": 8,
	}

	if !reflect.DeepEqual(expectedHasnewMapable, newHasnewMapable) {
		t.Errorf("Expected intersection result to be %v, but got %v", expectedHasnewMapable, newHasnewMapable)
	}

	// Test case: Check intersection of two hasnewMapables with no common key-value pairs.
	newMap1 = &gomap.Map[string, int]{
		"apple":  5,
		"orange": 8,
	}
	newMap2 = &gomap.Map[string, int]{
		"banana": 10,
		"grape":  7,
	}

	newHasnewMapable = newMap1.IntersectionFunc(newMap2, conditionFunc)

	expectedHasnewMapable = &gomap.Map[string, int]{}

	if !reflect.DeepEqual(expectedHasnewMapable, newHasnewMapable) {
		t.Errorf("Expected intersection result to be %v, but got %v", expectedHasnewMapable, newHasnewMapable)
	}

	// Test case: Check intersection of empty hasnewMapables.
	newMap1 = &gomap.Map[string, int]{}
	newMap2 = &gomap.Map[string, int]{}

	newHasnewMapable = newMap1.IntersectionFunc(newMap2, conditionFunc)

	expectedHasnewMapable = &gomap.Map[string, int]{}

	if !reflect.DeepEqual(expectedHasnewMapable, newMap1) {
		t.Errorf("Expected intersection result to be %v, but got %v", expectedHasnewMapable, newHasnewMapable)
	}
}

// TestKeys tests Map.Keys.
func TestKeys(t *testing.T) {
	// Create a new gomap.
	newMap := make(gomap.Map[string, int])
	newMap["apple"] = 5
	newMap["banana"] = 3
	newMap["cherry"] = 8

	// Get all keys from the gomap.
	keys := newMap.Keys()

	// Sort the keys for consistent iteration order.
	sort.Strings(*keys)

	// The expected keys slice: {"apple", "banana", "cherry"}.
	expectedKeys := &slice.Slice[string]{"apple", "banana", "cherry"}

	// Sort the keys for consistent iteration order.
	sort.Strings(*expectedKeys)

	// Verify that the obtained keys match the expected keys.
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Fatalf("Expected keys: %v, but got: %v", expectedKeys, keys)
	}
}

// TestKeysFunc tests Map.Keys.
func TestKeysFunc(t *testing.T) {
	// Create a new gomap.
	newMap := make(gomap.Map[string, int])
	newMap["apple"] = 5
	newMap["banana"] = 3
	newMap["cherry"] = 8

	// Get keys from the gomap where the key length is greater than 5.
	keys := newMap.KeysFunc(func(key string) bool {
		return strings.HasPrefix(key, "b")
	})

	// The expected keys slice: {"banana"}.
	expectedKeys := &slice.Slice[string]{"banana"}

	// Verify that the obtained keys match the expected keys.
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Fatalf("Expected keys: %v, but got: %v", expectedKeys, keys)
	}
}

// TestLength tests Map.Length.
func TestLength(t *testing.T) {
	// Create a new gomap.
	newMap := make(gomap.Map[string, int])
	newMap["apple"] = 5
	newMap["banana"] = 3
	newMap["cherry"] = 8

	// Get the length of the gomap.
	length := newMap.Length()

	// Expected length: 3.
	expectedLength := 3

	// Verify that the obtained length matches the expected length.
	if length != expectedLength {
		t.Fatalf("Expected length: %d, but got: %d", expectedLength, length)
	}
}

func TestMap(t *testing.T) {
	// Create a new gomap.
	newMap := make(gomap.Map[string, int])

	// Add key-value pairs to the gomap.
	newMap["apple"] = 5
	newMap["banana"] = 3
	newMap["cherry"] = 8

	// Define a function to double the values.
	doubleValue := func(key string, value int) int {
		return value * 2
	}

	// Apply the function to double the values in the gomap.
	doubledHT := newMap.Map(doubleValue)

	// Expected doubled values.
	expectedValues := map[string]int{"apple": 10, "banana": 6, "cherry": 16}
	for key, expectedValue := range expectedValues {
		value, exists := (*doubledHT)[key]
		if !exists || value != expectedValue {
			t.Fatalf("Expected value %d for key %s, but got %d", expectedValue, key, value)
		}
	}

	// Ensure the original gomap remains unchanged.
	for key, expectedValue := range expectedValues {
		value, exists := newMap[key]
		if !exists || value != expectedValue/2 {
			t.Fatalf("Expected original value %d for key %s, but got %d", expectedValue/2, key, value)
		}
	}
}

// TestMapBreak tests Map.MapBreak.
func TestMapBreak(t *testing.T) {
	// Create a new gomap.
	newMap := make(gomap.Map[string, int])

	// Add key-value pairs to the gomap.
	newMap["banana"] = 3

	// Apply the MapBreak function to modify values and break the iteration at "banana".
	newMap.MapBreak(func(key string, value int) (int, bool) {
		if key == "banana" {
			return value * 2, false // Break the iteration when key is "banana".
		}
		return value * 2, true // Continue iterating for other keys and double the values.
	})

	// Check if values are not modified as expected.
	expectedValues := map[string]int{"banana": 3}
	for key, expectedValue := range expectedValues {
		value, exists := newMap.Get(key)
		if !exists || value != expectedValue {
			t.Fatalf("Expected value %d for key %s, but got %d", expectedValue, key, value)
		}
	}
}

// TestMerge tests Map.Merge.
func TestMerge(t *testing.T) {
	// Test case: Merge all key-value pairs from another gomap.
	newMap1 := &gomap.Map[string, int]{} // Create a new gomap.
	newMap1.Add("apple", 5)

	newMap2 := &gomap.Map[string, int]{} // Create another gomap.
	newMap2.Add("orange", 10)

	// Merge all key-value pairs from newMap2 into newMap1.
	newMap1.Merge(newMap2)

	// After merging, newMap1 should contain: {"apple": 5, "orange": 10}
	expectedHasnewMapable := &gomap.Map[string, int]{
		"apple":  5,
		"orange": 10,
	}

	// Verify that newMap1 is equal to the expected gomap.
	if !reflect.DeepEqual(expectedHasnewMapable, newMap1) {
		t.Errorf("Merge did not produce the expected result. Got: %v, Expected: %v", newMap1, expectedHasnewMapable)
	}
}

// TestMergeFunc tests Map.MergeFunc.
func TestMergeFunc(t *testing.T) {
	// Test case: Merge key-value pairs based on the condition function.
	newMap1 := &gomap.Map[string, int]{} // Create a new gomap.
	newMap1.Add("apple", 5)
	newMap1.Add("orange", 10)

	newMap2 := &gomap.Map[string, int]{} // Create another gomap.
	newMap2.Add("orange", 8)
	newMap2.Add("banana", 6)

	// Condition function to merge pairs where the value in newMap2 is greater than 7.
	conditionFunc := func(key string, value int) bool {
		return value > 7
	}

	// Merge key-value pairs from newMap2 into newMap1 based on the condition function.
	newMap1.MergeFunc(newMap2, conditionFunc)

	// After merging, newMap1 should contain: {"apple": 5, "orange": 8}
	expectedHasnewMapable := &gomap.Map[string, int]{
		"apple":  5,
		"orange": 8,
	}

	// Verify that newMap1 is equal to the expected gomap.
	if !newMap1.Equal(expectedHasnewMapable) {
		t.Errorf("MergeFunc did not produce the expected result. Got: %v, Expected: %v", newMap1, expectedHasnewMapable)
	}
}

// TestMergeMany tests Map.MergeMany.
func TestMergeMany(t *testing.T) {
	// Create hasnewMapables for merging.
	newMap1 := &gomap.Map[string, int]{
		"apple":  5,
		"orange": 10,
	}
	newMap2 := &gomap.Map[string, int]{
		"orange": 15,
		"banana": 7,
	}
	newMap3 := &gomap.Map[string, int]{
		"grape": 8,
		"melon": 12,
	}

	// Merge key-value pairs from newMap2 and newMap3 into newMap1.
	newMap1.MergeMany(newMap2, newMap3)

	// Expected merged gomap.
	expectedHasnewMapable := &gomap.Map[string, int]{
		"apple":  5,
		"orange": 15,
		"banana": 7,
		"grape":  8,
		"melon":  12,
	}

	// Check if the merged gomap matches the expected gomap.
	if !reflect.DeepEqual(expectedHasnewMapable, newMap1) {
		t.Errorf("Merged gomap does not match the expected gomap. Got: %v, Expected: %v", newMap1, expectedHasnewMapable)
	}

	// Test case for merging an empty gomap.
	emptyHasnewMapable := &gomap.Map[string, int]{}
	newMap1.MergeMany(emptyHasnewMapable)

	// Merged gomap should remain unchanged.
	if !reflect.DeepEqual(expectedHasnewMapable, newMap1) {
		t.Errorf("Merged gomap should remain unchanged after merging with an empty gomap. Got: %v, Expected: %v", newMap1, expectedHasnewMapable)
	}
}

// TestMergeManyFunc tests Map.MergeManyFunc.
func TestMergeManyFunc(t *testing.T) {
	// Test case: Merge key-value pairs based on a condition function.
	newMap1 := &gomap.Map[string, int]{} // Create an empty destination gomap.
	newMap2 := &gomap.Map[string, int]{} // Create the first source gomap.
	newMap2.Add("apple", 5)
	newMap3 := &gomap.Map[string, int]{} // Create the second source gomap.
	newMap3.Add("orange", 10)
	newMap4 := &gomap.Map[string, int]{} // Create the third source gomap.
	newMap4.Add("banana", 7)

	// Condition function to include pairs only if the value is greater than 7.
	conditionFunc := func(i int, key string, value int) bool {
		return value >= 7
	}

	// Merge key-value pairs based on the condition function.
	mergedHasnewMapable := newMap1.MergeManyFunc([]*gomap.Map[string, int]{newMap2, newMap3, newMap4}, conditionFunc)

	// Verify that the merged gomap contains the expected key-value pairs.
	expectedPairs := map[string]int{
		"orange": 10,
		"banana": 7,
	}
	for key, expectedValue := range expectedPairs {
		value, ok := mergedHasnewMapable.Get(key)
		if !ok || value != expectedValue {
			t.Errorf("Expected merged gomap to contain key '%s': %d, but it contains '%d'", key, expectedValue, value)
		}
	}

	// Verify that unwanted pairs are not present in the merged gomap.
	unwantedPairs := map[string]int{
		"apple": 5,
	}
	for key := range unwantedPairs {
		_, ok := mergedHasnewMapable.Get(key)
		if ok {
			t.Errorf("Expected merged gomap not to contain key '%s', but it was found", key)
		}
	}
}

// TestNot tests Map.Not.
func TestNot(t *testing.T) {
	// Test case 1: Check if a key is not present in an empty gomap.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	result := newMap.Not("apple")       // Check if "apple" is not in the gomap.
	expected := true                    // "apple" is not present in the empty gomap.

	if result != expected {
		t.Errorf("Expected result to be %v for key 'apple', but got %v", expected, result)
	}

	// Test case 2: Check if a key is not present in a non-empty gomap.
	newMap.Add("orange", 5)
	newMap.Add("banana", 10)
	result = newMap.Not("banana") // Check if "banana" is not in the gomap.
	expected = false              // "banana" is present in the gomap.

	if result != expected {
		t.Errorf("Expected result to be %v for key 'banana', but got %v", expected, result)
	}

	// Test case 3: Check if a key is not present after removing it from the gomap.
	newMap.Delete("banana") // Delete "banana" from the gomap.
	result = newMap.Not("banana")
	expected = true // "banana" is not present after removal.

	if result != expected {
		t.Errorf("Expected result to be %v for key 'banana' after removal, but got %v", expected, result)
	}
}

// TestPop tests Map.Pop.
func TestPop(t *testing.T) {
	// Test case 1: Pop from an empty gomap.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	removedValue := newMap.Pop("apple")
	expectedValue := 0 // No key "apple" in the empty gomap.

	if removedValue != expectedValue {
		t.Errorf("Expected removed value to be %d, but got %d", expectedValue, removedValue)
	}

	// Test case 2: Pop from a non-empty gomap where the key is present.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	removedValue = newMap.Pop("apple")
	expectedValue = 5 // Key "apple" exists with value 5.

	if removedValue != expectedValue {
		t.Errorf("Expected removed value to be %d, but got %d", expectedValue, removedValue)
	}
	// Verify that the key is removed.
	_, ok := newMap.Get("apple")
	if ok {
		t.Errorf("Expected key 'apple' to be removed, but it was found")
	}

	// Test case 3: Pop from a non-empty gomap where the key is not present.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	newMap.Add("orange", 10)
	removedValue = newMap.Pop("banana")
	expectedValue = 0 // No key "banana" in the gomap.

	if removedValue != expectedValue {
		t.Errorf("Expected removed value to be %d, but got %d", expectedValue, removedValue)
	}
}

// TestPopOK tests Map.PopOK.
func TestPopOK(t *testing.T) {
	// Test case 1: Pop from an empty gomap.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	removedValue, ok := newMap.PopOK("apple")
	if ok || removedValue != 0 {
		t.Errorf("Expected (0, false), but got (%d, %v)", removedValue, ok)
	}

	// Test case 2: Pop from a non-empty gomap where the key is present.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	removedValue, ok = newMap.PopOK("apple")
	if !ok || removedValue != 5 {
		t.Errorf("Expected (5, true), but got (%d, %v)", removedValue, ok)
	}
	// Verify that the key is removed.
	_, ok = newMap.Get("apple")
	if ok {
		t.Errorf("Expected key 'apple' to be removed, but it was found")
	}

	// Test case 3: Pop from a non-empty gomap where the key is not present.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	newMap.Add("orange", 10)
	removedValue, ok = newMap.PopOK("banana")
	if ok || removedValue != 0 {
		t.Errorf("Expected (0, false), but got (%d, %v)", removedValue, ok)
	}
}

func TestPopMany(t *testing.T) {
	// Test case 1: PopMany from an empty gomap.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	removedValues := newMap.PopMany("apple", "orange")
	expectedValues := &slice.Slice[int]{}
	if !reflect.DeepEqual(removedValues, expectedValues) {
		t.Errorf("Expected %v, but got %v", expectedValues, removedValues)
	}

	// Test case 2: PopMany from a non-empty gomap where some keys are present.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	newMap.Add("banana", 3)
	newMap.Add("cherry", 8)
	removedValues = newMap.PopMany("apple", "orange", "cherry", "grape")
	expectedValues = &slice.Slice[int]{5, 8}
	if !reflect.DeepEqual(removedValues, expectedValues) {
		t.Errorf("Expected %v, but got %v", expectedValues, removedValues)
	}
	// Verify that the keys are removed.
	_, ok := newMap.Get("apple")
	if ok {
		t.Errorf("Expected key 'apple' to be removed, but it was found")
	}
	_, ok = newMap.Get("orange")
	if ok {
		t.Errorf("Expected key 'orange' to be removed, but it was found")
	}
	_, ok = newMap.Get("cherry")
	if ok {
		t.Errorf("Expected key 'cherry' to be removed, but it was found")
	}
	_, ok = newMap.Get("grape")
	if ok {
		t.Errorf("Expected key 'grape' to be removed, but it was found")
	}
}

// TestPopManyFunc tests Map.PopManyFunc.
func TestPopManyFunc(t *testing.T) {
	// Test case 1: Pop values greater than 7 from the gomap.
	newMap := &gomap.Map[string, int]{
		"apple":  5,
		"orange": 10,
		"banana": 8,
		"grape":  12,
	}

	removeCondition := func(key string, value int) bool {
		return value > 7
	}

	removedValues := newMap.PopManyFunc(removeCondition)

	sort.Ints(*removedValues)

	expectedRemovedValues := &slice.Slice[int]{8, 10, 12}

	if !reflect.DeepEqual(expectedRemovedValues, removedValues) {
		t.Errorf("Expected removed values to be %v, but got %v", expectedRemovedValues, removedValues)
	}

	// Test case 2: Pop values when condition does not match any key-value pairs.
	newMap = &gomap.Map[string, int]{
		"apple":  5,
		"orange": 3,
		"banana": 8,
	}

	removeCondition = func(key string, value int) bool {
		return value > 10
	}

	removedValues = newMap.PopManyFunc(removeCondition)

	expectedRemovedValues = &slice.Slice[int]{} // No values match the condition.

	if !reflect.DeepEqual(expectedRemovedValues, removedValues) {
		t.Errorf("Expected removed values to be %v, but got %v", expectedRemovedValues, removedValues)
	}

	// Test case 3: Pop values from an empty gomap.
	newMap = &gomap.Map[string, int]{}

	removeCondition = func(key string, value int) bool {
		return value > 0
	}

	removedValues = newMap.PopManyFunc(removeCondition)

	expectedRemovedValues = &slice.Slice[int]{} // No values to remove from an empty gomap.

	if !reflect.DeepEqual(expectedRemovedValues, removedValues) {
		t.Errorf("Expected removed values to be %v, but got %v", expectedRemovedValues, removedValues)
	}
}

// TestReplaceMany tests Map.ReplaceMany.
func TestUpdate(t *testing.T) {
	// Test case 1: Replace with an empty gomap and a function that never modifies any pairs.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	replaceFunc := func(key string, value int) (int, bool) {
		return value, false // Never modify any values
	}
	newMap.ReplaceMany(replaceFunc)
	expected := &gomap.Map[string, int]{} // Expected empty gomap.
	if !reflect.DeepEqual(newMap, expected) {
		t.Errorf("Expected %v, but got %v", expected, newMap)
	}

	// Test case 2: Replace with a non-empty gomap and a function that never modifies any pairs.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	newMap.Add("orange", 10)
	replaceFunc = func(key string, value int) (int, bool) {
		return value, false // Never modify any values
	}
	newMap.ReplaceMany(replaceFunc)
	expected = &gomap.Map[string, int]{"apple": 5, "orange": 10} // Expected same gomap.
	if !reflect.DeepEqual(newMap, expected) {
		t.Errorf("Expected %v, but got %v", expected, newMap)
	}

	// Test case 3: Replace with a non-empty gomap and a function that modifies certain pairs.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	newMap.Add("orange", 10)
	replaceFunc = func(key string, value int) (int, bool) {
		if key == "apple" {
			return value * 2, true // Modify the value for the "apple" key
		}
		return value, false // Leave other values unchanged
	}
	newMap.ReplaceMany(replaceFunc)
	expected = &gomap.Map[string, int]{"apple": 10, "orange": 10} // Expected modified gomap.
	if !reflect.DeepEqual(newMap, expected) {
		t.Errorf("Expected %v, but got %v", expected, newMap)
	}
}

// TestTakeFrom tests Map.TakeFrom.
func TestTakeFrom(t *testing.T) {
	// Test case 1: Transfer from an empty gomap to another empty gomap.
	newMap1 := &gomap.Map[string, int]{} // Create an empty destination gomap.
	newMap2 := &gomap.Map[string, int]{} // Create an empty source gomap.
	newMap1.TakeFrom(newMap2)            // Transfer from newMap2 to newMap1.

	// Verify that newMap1 is still empty.
	if newMap1.Length() != 0 {
		t.Errorf("Expected destination gomap to be empty after transfer, but it has %d items", newMap1.Length())
	}

	// Verify that newMap2 is still empty.
	if newMap2.Length() != 0 {
		t.Errorf("Expected source gomap to be empty after transfer, but it has %d items", newMap2.Length())
	}

	// Test case 2: Transfer from a non-empty gomap to an empty gomap.
	newMap1 = &gomap.Map[string, int]{} // Create an empty destination gomap.
	newMap2 = &gomap.Map[string, int]{} // Create a source gomap.
	newMap2.Add("orange", 10)
	newMap1.TakeFrom(newMap2) // Transfer from newMap2 to newMap1.

	// Verify that newMap1 contains the transferred key-value pair.
	expectedValue := 10
	transferredValue, ok := newMap1.Get("orange")
	if !ok || transferredValue != expectedValue {
		t.Errorf("Expected destination gomap to contain 'orange': %d after transfer, but it contains '%d'", expectedValue, transferredValue)
	}

	// Verify that newMap2 is empty after transfer.
	if newMap2.Length() != 0 {
		t.Errorf("Expected source gomap to be empty after transfer, but it has %d items", newMap2.Length())
	}
}

// TestValues tests Map.Values.
func TestValues(t *testing.T) {
	// Test case 1: Values of an empty gomap should be an empty slice.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	values := newMap.Values()
	expected := &slice.Slice[int]{} // Expected empty slice.
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, but got %v", expected, values)
	}

	// Test case 2: Values of a non-empty gomap.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	newMap.Add("orange", 10)
	values = newMap.Values()
	sort.Ints(*values)
	expected = &slice.Slice[int]{5, 10} // Expected slice containing [5, 10].
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, but got %v", expected, values)
	}

	// Test case 3: Values of a gomap with multiple entries.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	newMap.Add("orange", 10)
	newMap.Add("banana", 15)
	values = newMap.Values()

	sort.Ints(*values)
	expected = &slice.Slice[int]{5, 10, 15} // Expected slice containing [5, 10, 15].
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, but got %v", expected, values)
	}
}

// TestValuesFunc tests Map.ValuesFunc.
func TestValuesFunc(t *testing.T) {
	// Test case 1: ValuesFunc with an empty gomap and a condition that never satisfies.
	newMap := &gomap.Map[string, int]{} // Create an empty gomap.
	filterFunc := func(key string, value int) bool {
		return value > 7 // Include values greater than 7 in the result
	}
	values := newMap.ValuesFunc(filterFunc)
	expected := &slice.Slice[int]{} // Expected empty slice.
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, but got %v", expected, values)
	}

	// Test case 2: ValuesFunc with a non-empty gomap and a condition that never satisfies.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	newMap.Add("orange", 2)
	filterFunc = func(key string, value int) bool {
		return value > 7 // Include values greater than 7 in the result
	}
	values = newMap.ValuesFunc(filterFunc)
	expected = &slice.Slice[int]{} // Expected empty slice.
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, but got %v", expected, values)
	}

	// Test case 3: ValuesFunc with a non-empty gomap and a condition that satisfies for some values.
	newMap = &gomap.Map[string, int]{} // Create an empty gomap.
	newMap.Add("apple", 5)
	newMap.Add("orange", 10)
	newMap.Add("banana", 15)
	filterFunc = func(key string, value int) bool {
		return value > 7 // Include values greater than 7 in the result
	}
	values = newMap.ValuesFunc(filterFunc)
	sort.Ints(*values)
	expected = &slice.Slice[int]{10, 15} // Expected slice containing [10, 15].
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, but got %v", expected, values)
	}
}
