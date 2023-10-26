package hashtable_test

import (
	"math"
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
	// Test case 1: Add key-value pairs to an empty hashtable.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	ht.Add("banana", 3)

	// Verify the values using Get method.
	value, ok := ht.Get("apple")
	if !ok || value != 5 {
		t.Errorf("Expected key 'apple' to have value 5, but got %v", value)
	}
	value, ok = ht.Get("banana")
	if !ok || value != 3 {
		t.Errorf("Expected key 'banana' to have value 3, but got %v", value)
	}

	// Test case 2: Replace existing key-value pair.
	ht.Add("banana", 10) // Updates the value for the key "banana" to 10.
	value, ok = ht.Get("banana")
	if !ok || value != 10 {
		t.Errorf("Expected key 'banana' to have updated value 10, but got %v", value)
	}

	// Test case 3: Add a new key-value pair.
	ht.Add("cherry", 8)
	value, ok = ht.Get("cherry")
	if !ok || value != 8 {
		t.Errorf("Expected key 'cherry' to have value 8, but got %v", value)
	}

	// Verify that other keys are not affected.
	value, ok = ht.Get("grape")
	if ok {
		t.Errorf("Expected key 'grape' to be absent, but it was found with value %v", value)
	}
}

// TestAddLength tests Hashtable.AddLength.
func TestAddLength(t *testing.T) {
	// Test case 1: Add key-value pairs to an empty hashtable.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	length := ht.AddLength("apple", 5)
	if length != 1 {
		t.Errorf("Expected length of hashtable after adding 'apple' with value 5 to be 1, but got %v", length)
	}

	// Test case 2: Replace an existing key-value pair.
	length = ht.AddLength("apple", 10)
	if length != 1 {
		t.Errorf("Expected length of hashtable after updating 'apple' with value 10 to be 1, but got %v", length)
	}

	// Test case 3: Add a new key-value pair.
	length = ht.AddLength("banana", 3)
	if length != 2 {
		t.Errorf("Expected length of hashtable after adding 'banana' with value 3 to be 2, but got %v", length)
	}

	// Test case 4: Add another new key-value pair.
	length = ht.AddLength("cherry", 8)
	if length != 3 {
		t.Errorf("Expected length of hashtable after adding 'cherry' with value 8 to be 3, but got %v", length)
	}

	// Verify that other keys are not affected.
	value, ok := ht.Get("grape")
	if ok {
		t.Errorf("Expected key 'grape' to be absent, but it was found with value %v", value)
	}
}

// TestAddMany tests Hashtable.AddMany.

func TestAddMany(t *testing.T) {
	// Test case 1: Add key-value pairs to an empty hashtable.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.AddMany(map[string]int{"orange": 7, "grape": 4}, map[string]int{"kiwi": 6, "pear": 9})

	// Verify the added key-value pairs.
	expected := map[string]int{"orange": 7, "grape": 4, "kiwi": 6, "pear": 9}
	for key, expectedValue := range expected {
		value, ok := ht.Get(key)
		if !ok {
			t.Errorf("Expected key '%s' to be present, but it was not found", key)
		}
		if value != expectedValue {
			t.Errorf("Expected value for key '%s' to be %d, but got %d", key, expectedValue, value)
		}
	}

	// Test case 2: Replace existing key-value pairs and add new ones.
	ht.AddMany(map[string]int{"orange": 10, "grape": 3}, map[string]int{"apple": 5, "banana": 8})

	// Verify the updated and added key-value pairs.
	expected = map[string]int{"orange": 10, "grape": 3, "kiwi": 6, "pear": 9, "apple": 5, "banana": 8}
	for key, expectedValue := range expected {
		value, ok := ht.Get(key)
		if !ok {
			t.Errorf("Expected key '%s' to be present, but it was not found", key)
		}
		if value != expectedValue {
			t.Errorf("Expected value for key '%s' to be %d, but got %d", key, expectedValue, value)
		}
	}

	// Test case 3: Add new key-value pairs.
	ht.AddMany(map[string]int{"watermelon": 12, "cherry": 7})

	// Verify the added key-value pairs.
	expected = map[string]int{"orange": 10, "grape": 3, "kiwi": 6, "pear": 9, "apple": 5, "banana": 8, "watermelon": 12, "cherry": 7}
	for key, expectedValue := range expected {
		value, ok := ht.Get(key)
		if !ok {
			t.Errorf("Expected key '%s' to be present, but it was not found", key)
		}
		if value != expectedValue {
			t.Errorf("Expected value for key '%s' to be %d, but got %d", key, expectedValue, value)
		}
	}

	// Verify that other keys are not affected.
	_, ok := ht.Get("pineapple")
	if ok {
		t.Errorf("Expected key 'pineapple' to be absent, but it was found")
	}
}

// TestAddManyFunc tests Hashtable.AddManyFunc.
func TestAddManyFunc(t *testing.T) {
	// Test case 1: Add key-value pairs with values greater than 0 to the hashtable.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	values := []map[string]int{{"apple": 5, "orange": -3, "banana": 10}}
	condition := func(i int, key string, value int) bool {
		return value > 0
	}
	ht.AddManyFunc(values, condition)

	// Verify that only the key-value pairs with values greater than 0 are added to the hashtable.
	expected := &hashtable.Hashtable[string, int]{"apple": 5, "banana": 10}
	if !reflect.DeepEqual(ht, expected) {
		t.Errorf("Expected %v, but got %v", expected, ht)
	}

	// Test case 2: Add all key-value pairs to the hashtable.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	values = []map[string]int{{"apple": 5, "orange": -3, "banana": 10}}
	condition = func(i int, key string, value int) bool {
		return true // Add all key-value pairs
	}
	ht.AddManyFunc(values, condition)

	// Verify that all key-value pairs are added to the hashtable.
	expected = &hashtable.Hashtable[string, int]{"apple": 5, "orange": -3, "banana": 10}
	if !reflect.DeepEqual(ht, expected) {
		t.Errorf("Expected %v, but got %v", expected, ht)
	}
}

// TestAddManyOK tests Hashtable.AddManyOK.
func TestAddManyOK(t *testing.T) {
	// Test case 1: Add key-value pairs to an empty hashtable.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	results := ht.AddManyOK(
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
		value, ok := ht.Get(key)
		if !ok {
			t.Errorf("Expected key '%s' to be present, but it was not found", key)
		}
		if value != expectedValue {
			t.Errorf("Expected value for key '%s' to be %d, but got %d", key, expectedValue, value)
		}
	}

	// Test case 2: Add new key-value pairs.
	results = ht.AddManyOK(
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
		value, ok := ht.Get(key)
		if !ok {
			t.Errorf("Expected key '%s' to be present, but it was not found", key)
		}
		if value != expectedValue {
			t.Errorf("Expected value for key '%s' to be %d, but got %d", key, expectedValue, value)
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

// TestContains tests Hashtable.Contains.
func TestContains(t *testing.T) {
	// Test case 1: Check for a value in an empty hashtable.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	key, found := ht.Contains(5)              // Check if value 5 is in the hashtable.

	// Since the hashtable is empty, the result should be ("", false).
	if key != "" || found {
		t.Errorf("Expected ('', false), but got (%s, %t)", key, found)
	}

	// Test case 2: Check for a value in a non-empty hashtable.
	ht = &hashtable.Hashtable[string, int]{} // Create a new hashtable.
	ht.Add("apple", 5)
	ht.Add("banana", 10)
	ht.Add("orange", 5)

	// Check if value 5 is in the hashtable.
	key, found = ht.Contains(5)

	// Since value 5 exists in the hashtable, the result should be ("apple", true).
	if !found {
		t.Errorf("Expected ('%s', true), but got (%s, %t)", key, key, found)
	}

	// Test case 3: Check for a value that doesn't exist in the hashtable.
	key, found = ht.Contains(15) // Check if value 15 is in the hashtable.

	// Since value 15 doesn't exist in the hashtable, the result should be ("", false).
	if key != "" || found {
		t.Errorf("Expected ('', false), but got (%s, %t)", key, found)
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

	// Expected length after deleting "apple": initial length - 1.
	expectedLength := initialLength - 1

	// Verify that the obtained length matches the expected length.
	if lengthAfterDelete != expectedLength {
		t.Fatalf("Expected length: %d, but got: %d", expectedLength, lengthAfterDelete)
	}

	// Attempt to delete a non-existing key from the hashtable.
	lengthAfterNonExistingDelete := ht.DeleteLength("grape")

	// Length should remain the same after attempting to delete a non-existing key.
	// Expected length: initial length.
	expectedLengthNonExisting := len(ht)

	// Verify that the obtained length matches the expected length.
	if lengthAfterNonExistingDelete != expectedLengthNonExisting {
		t.Fatalf("Expected length: %d, but got: %d", expectedLengthNonExisting, lengthAfterNonExistingDelete)
	}
}

// TestDeleteMany tests Hashtable.DeleteMany.
func TestDeleteMany(t *testing.T) {
	// Test case 1: Delete keys from an empty hashtable.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.DeleteMany("apple", "banana")          // Attempt to delete keys "apple" and "banana".

	// The hashtable should remain empty.
	if !ht.IsEmpty() {
		t.Errorf("Expected hashtable to be empty, but got non-empty hashtable: %v", ht)
	}

	// Test case 2: Delete keys from a non-empty hashtable.
	ht = &hashtable.Hashtable[string, int]{} // Create a new hashtable.
	ht.Add("apple", 5)
	ht.Add("banana", 10)
	ht.Add("orange", 3)

	// Delete keys "apple" and "banana".
	ht.DeleteMany("apple", "banana")

	// Verify that "apple" and "banana" are deleted, and "orange" remains in the hashtable.
	expected := &slice.Slice[string]{"orange"}
	result := ht.Keys()

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected keys %v after deleting 'apple' and 'banana', but got keys %v", expected, result)
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

	expectedResults := []bool{true, true} // Expected results for "apple" (exists) and "grape" (does not exist).

	// Check if results match the expected results.
	for i, result := range *results {
		if result != expectedResults[i] {
			t.Fatalf("Expected deletion of key %s to be %v but got %v", keysToDelete[i], expectedResults[i], result)
		}
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
		t.Fatalf("Expected deletion of 'apple' to be successful")
	}

	// Attempt to delete a key that does not exist.
	notDeleted := ht.DeleteOK("grape")
	if !notDeleted {
		t.Fatalf("Expected deletion of 'grape' to be successful because the key does not exist")
	}

	// Attempt to delete a key that has already been deleted.
	alreadyDeleted := ht.DeleteOK("apple")
	if !alreadyDeleted {
		t.Fatalf("Expected deletion of 'apple' to be successful even though it was already deleted")
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

	// Expected output: "apple", "banana", "cherry".
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

// TestEmptyInto tests Hashtable.EmptyInto.
func TestEmptyInto(t *testing.T) {
	// Test case 1: Transfer from an empty hashtable to another empty hashtable.
	ht1 := &hashtable.Hashtable[string, int]{} // Create an empty source hashtable.
	ht2 := &hashtable.Hashtable[string, int]{} // Create an empty destination hashtable.
	ht1.EmptyInto(ht2)                         // Transfer from ht1 to ht2.

	// Verify that ht1 is empty.
	if ht1.Length() != 0 {
		t.Errorf("Expected source hashtable to be empty after transfer, but it has %d items", ht1.Length())
	}

	// Verify that ht2 is still empty.
	if ht2.Length() != 0 {
		t.Errorf("Expected destination hashtable to be empty after transfer, but it has %d items", ht2.Length())
	}

	// Test case 2: Transfer from a non-empty hashtable to an empty hashtable.
	ht1 = &hashtable.Hashtable[string, int]{} // Create an empty source hashtable.
	ht1.Add("apple", 5)
	ht2 = &hashtable.Hashtable[string, int]{} // Create an empty destination hashtable.
	ht1.EmptyInto(ht2)                        // Transfer from ht1 to ht2.

	// Verify that ht1 is empty.
	if ht1.Length() != 0 {
		t.Errorf("Expected source hashtable to be empty after transfer, but it has %d items", ht1.Length())
	}

	// Verify that ht2 contains the transferred key-value pair.
	expectedValue := 5
	transferredValue, ok := ht2.Get("apple")
	if !ok || transferredValue != expectedValue {
		t.Errorf("Expected destination hashtable to contain 'apple': %d after transfer, but it contains '%d'", expectedValue, transferredValue)
	}
}

// TestEqual tests Hashtable.Equal.
func TestEqual(t *testing.T) {
	// Test case 1: Compare equal hashtables.
	ht1 := &hashtable.Hashtable[string, int]{} // Create a new hashtable.
	ht1.Add("apple", 5)
	ht1.Add("orange", 10)

	ht2 := &hashtable.Hashtable[string, int]{} // Create another hashtable with similar values.
	ht2.Add("apple", 5)
	ht2.Add("orange", 10)

	// Check if the two hashtables are equal.
	equal := ht1.Equal(ht2)

	// Since ht1 and ht2 have the same key-value pairs, they are considered equal.
	if !equal {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Compare unequal hashtables.
	ht3 := &hashtable.Hashtable[string, int]{} // Create a new hashtable.
	ht3.Add("apple", 5)
	ht3.Add("orange", 10)

	ht4 := &hashtable.Hashtable[string, int]{} // Create another hashtable with different values for "orange".
	ht4.Add("apple", 5)
	ht4.Add("orange", 12)

	// Check if the two hashtables are equal.
	equal = ht3.Equal(ht4)

	// Since ht3 and ht4 have different values for "orange", they are not considered equal.
	if equal {
		t.Errorf("Expected false, but got true")
	}
}

// TestEqualFunc tests Hashtable.EqualFunc.
func TestEqualFunc(t *testing.T) {
	// Custom comparison function to check if two integers are equal when their difference is less than or equal to 1
	compareFunc := func(a, b int) bool {
		return math.Abs(float64(a-b)) <= 1
	}

	// Test case 1: Compare equal hashtables based on the custom comparison function.
	ht1 := &hashtable.Hashtable[string, int]{} // Create a new hashtable.
	ht1.Add("apple", 5)
	ht1.Add("orange", 10)

	ht2 := &hashtable.Hashtable[string, int]{} // Create another hashtable with similar values.
	ht2.Add("apple", 5)
	ht2.Add("orange", 11) // The difference between 10 and 11 is within the allowed range according to compareFunc.

	// Check if the two hashtables are equal based on the custom comparison function.
	equal := ht1.EqualFunc(ht2, compareFunc)

	// Since the values for "orange" (10 and 11) have a difference of 1, within the allowed range,
	// the hashtables are considered equal according to the custom comparison function.
	if !equal {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Compare unequal hashtables based on the custom comparison function.
	ht3 := &hashtable.Hashtable[string, int]{} // Create a new hashtable.
	ht3.Add("apple", 5)
	ht3.Add("orange", 10)

	ht4 := &hashtable.Hashtable[string, int]{} // Create another hashtable with different values for "orange".
	ht4.Add("apple", 5)
	ht4.Add("orange", 12) // The difference between 10 and 12 is greater than the allowed range according to compareFunc.

	// Check if the two hashtables are equal based on the custom comparison function.
	equal = ht3.EqualFunc(ht4, compareFunc)

	// Since the difference between 10 and 12 is greater than the allowed range according to compareFunc,
	// the hashtables are not considered equal based on the custom comparison function.
	if equal {
		t.Errorf("Expected false, but got true")
	}
}

// TestEqualLength tests Hashtable.EqualLength.
func TestEqualLength(t *testing.T) {
	// Test case 1: Compare hashtables with equal length.
	ht1 := &hashtable.Hashtable[string, int]{} // Create a new hashtable.
	ht1.Add("apple", 5)
	ht1.Add("orange", 10)

	ht2 := &hashtable.Hashtable[string, int]{} // Create another hashtable with the same number of key-value pairs.
	ht2.Add("apple", 5)
	ht2.Add("orange", 7)

	// Check if the two hashtables have equal length.
	equalLength := ht1.EqualLength(ht2)

	// Since ht1 and ht2 have the same number of key-value pairs, they are considered equal in length.
	if !equalLength {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Compare hashtables with unequal length.
	ht3 := &hashtable.Hashtable[string, int]{} // Create a new hashtable.
	ht3.Add("apple", 5)

	ht4 := &hashtable.Hashtable[string, int]{} // Create another hashtable with a different number of key-value pairs.
	ht4.Add("apple", 5)
	ht4.Add("orange", 7)

	// Check if the two hashtables have equal length.
	equalLength = ht3.EqualLength(ht4)

	// Since ht3 and ht4 have different numbers of key-value pairs, they are not considered equal in length.
	if equalLength {
		t.Errorf("Expected false, but got true")
	}
}

// TestFetch tests Hashtable.Fetch.
func TestFetch(t *testing.T) {
	// Test case 1: Fetch value for an existing key.
	ht := &hashtable.Hashtable[string, int]{} // Create a new hashtable.
	ht.Add("apple", 5)
	ht.Add("orange", 10)

	// Fetch the value associated with the key "apple".
	fetchedValue := ht.Fetch("apple")

	// Since "apple" is in the hashtable, the fetched value should be 5.
	expectedValue := 5
	if fetchedValue != expectedValue {
		t.Errorf("Expected %d, but got %d", expectedValue, fetchedValue)
	}

	// Test case 2: Fetch value for a non-existing key.
	// Fetch the value associated with the key "banana", which is not in the hashtable.
	fetchedValue = ht.Fetch("banana")

	// Since "banana" is not in the hashtable, the fetched value should be the zero value for int, which is 0.
	expectedValue = 0
	if fetchedValue != expectedValue {
		t.Errorf("Expected %d, but got %d", expectedValue, fetchedValue)
	}
}

// TestFilter tests Hashtable.Filter.
func TestFilter(t *testing.T) {
	// Test case 1: Filter with an empty hashtable and a function that never selects any pairs.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	filterFunc := func(key string, value int) bool {
		return false // Never select any values
	}
	filtered := ht.Filter(filterFunc)
	expected := &hashtable.Hashtable[string, int]{} // Expected empty hashtable.
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Expected %v, but got %v", expected, filtered)
	}

	// Test case 2: Filter with a non-empty hashtable and a function that never selects any pairs.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	ht.Add("orange", 10)
	filterFunc = func(key string, value int) bool {
		return false // Never select any values
	}
	filtered = ht.Filter(filterFunc)
	expected = &hashtable.Hashtable[string, int]{} // Expected empty hashtable.
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Expected %v, but got %v", expected, filtered)
	}

	// Test case 3: Filter with a non-empty hashtable and a function that selects certain pairs.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	ht.Add("orange", 10)
	ht.Add("banana", 3)
	filterFunc = func(key string, value int) bool {
		return value > 4 // Select pairs where value is greater than 4
	}
	filtered = ht.Filter(filterFunc)
	expected = &hashtable.Hashtable[string, int]{"apple": 5, "orange": 10} // Expected filtered hashtable.
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Expected %v, but got %v", expected, filtered)
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

	// The expected boolean slice: {true, false, true}.
	expectedResults := &slice.Slice[bool]{true, false, true}

	// Verify that the obtained results match the expected results.
	if !reflect.DeepEqual(results, expectedResults) {
		t.Fatalf("Expected results: %v, but got: %v", expectedResults, results)
	}
}

// TestIntersectionFunc tests Hashtable.IntersectionFunc.
func TestIntersectionFunc(t *testing.T) {
	// Test case: Check intersection of two hashtables with common key-value pairs.
	ht1 := &hashtable.Hashtable[string, int]{
		"apple":  5,
		"orange": 8,
	}
	ht2 := &hashtable.Hashtable[string, int]{
		"orange": 8,
		"banana": 6,
	}

	// Condition function to check if values are equal.
	conditionFunc := func(key string, a, b int) bool {
		return a == b
	}

	newHashtable := ht1.IntersectionFunc(ht2, conditionFunc)

	expectedHashtable := &hashtable.Hashtable[string, int]{
		"orange": 8,
	}

	if !reflect.DeepEqual(expectedHashtable, newHashtable) {
		t.Errorf("Expected intersection result to be %v, but got %v", expectedHashtable, newHashtable)
	}

	// Test case: Check intersection of two hashtables with no common key-value pairs.
	ht1 = &hashtable.Hashtable[string, int]{
		"apple":  5,
		"orange": 8,
	}
	ht2 = &hashtable.Hashtable[string, int]{
		"banana": 10,
		"grape":  7,
	}

	newHashtable = ht1.IntersectionFunc(ht2, conditionFunc)

	expectedHashtable = &hashtable.Hashtable[string, int]{}

	if !reflect.DeepEqual(expectedHashtable, newHashtable) {
		t.Errorf("Expected intersection result to be %v, but got %v", expectedHashtable, newHashtable)
	}

	// Test case: Check intersection of empty hashtables.
	ht1 = &hashtable.Hashtable[string, int]{}
	ht2 = &hashtable.Hashtable[string, int]{}

	newHashtable = ht1.IntersectionFunc(ht2, conditionFunc)

	expectedHashtable = &hashtable.Hashtable[string, int]{}

	if !reflect.DeepEqual(expectedHashtable, ht1) {
		t.Errorf("Expected intersection result to be %v, but got %v", expectedHashtable, newHashtable)
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

	// The expected keys slice: {"apple", "banana", "cherry"}.
	expectedKeys := &slice.Slice[string]{"apple", "banana", "cherry"}

	// Sort the keys for consistent iteration order.
	sort.Strings(*expectedKeys)

	// Verify that the obtained keys match the expected keys.
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Fatalf("Expected keys: %v, but got: %v", expectedKeys, keys)
	}
}

// TestKeysFunc tests Hashtable.Keys.
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

	// The expected keys slice: {"banana"}.
	expectedKeys := &slice.Slice[string]{"banana"}

	// Verify that the obtained keys match the expected keys.
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Fatalf("Expected keys: %v, but got: %v", expectedKeys, keys)
	}
}

// TestLength tests Hashtable.Length.
func TestLength(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Get the length of the hashtable.
	length := ht.Length()

	// Expected length: 3.
	expectedLength := 3

	// Verify that the obtained length matches the expected length.
	if length != expectedLength {
		t.Fatalf("Expected length: %d, but got: %d", expectedLength, length)
	}
}

func TestMap(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])

	// Add key-value pairs to the hashtable.
	ht["apple"] = 5
	ht["banana"] = 3
	ht["cherry"] = 8

	// Define a function to double the values.
	doubleValue := func(key string, value int) int {
		return value * 2
	}

	// Apply the function to double the values in the hashtable.
	doubledHT := ht.Map(doubleValue)

	// Expected doubled values.
	expectedValues := map[string]int{"apple": 10, "banana": 6, "cherry": 16}
	for key, expectedValue := range expectedValues {
		value, exists := (*doubledHT)[key]
		if !exists || value != expectedValue {
			t.Fatalf("Expected value %d for key %s, but got %d", expectedValue, key, value)
		}
	}

	// Ensure the original hashtable remains unchanged.
	for key, expectedValue := range expectedValues {
		value, exists := ht[key]
		if !exists || value != expectedValue/2 {
			t.Fatalf("Expected original value %d for key %s, but got %d", expectedValue/2, key, value)
		}
	}
}

// TestMapBreak tests Hashtable.MapBreak.
func TestMapBreak(t *testing.T) {
	// Create a new hashtable.
	ht := make(hashtable.Hashtable[string, int])

	// Add key-value pairs to the hashtable.
	ht["banana"] = 3

	// Apply the MapBreak function to modify values and break the iteration at "banana".
	ht.MapBreak(func(key string, value int) (int, bool) {
		if key == "banana" {
			return value * 2, false // Break the iteration when key is "banana".
		}
		return value * 2, true // Continue iterating for other keys and double the values.
	})

	// Check if values are not modified as expected.
	expectedValues := map[string]int{"banana": 3}
	for key, expectedValue := range expectedValues {
		value, exists := ht.Get(key)
		if !exists || value != expectedValue {
			t.Fatalf("Expected value %d for key %s, but got %d", expectedValue, key, value)
		}
	}
}

// TestMerge tests Hashtable.Merge.
func TestMerge(t *testing.T) {
	// Test case: Merge all key-value pairs from another hashtable.
	ht1 := &hashtable.Hashtable[string, int]{} // Create a new hashtable.
	ht1.Add("apple", 5)

	ht2 := &hashtable.Hashtable[string, int]{} // Create another hashtable.
	ht2.Add("orange", 10)

	// Merge all key-value pairs from ht2 into ht1.
	ht1.Merge(ht2)

	// After merging, ht1 should contain: {"apple": 5, "orange": 10}
	expectedHashtable := &hashtable.Hashtable[string, int]{
		"apple":  5,
		"orange": 10,
	}

	// Verify that ht1 is equal to the expected hashtable.
	if !reflect.DeepEqual(expectedHashtable, ht1) {
		t.Errorf("Merge did not produce the expected result. Got: %v, Expected: %v", ht1, expectedHashtable)
	}
}

// TestMergeFunc tests Hashtable.MergeFunc.
func TestMergeFunc(t *testing.T) {
	// Test case: Merge key-value pairs based on the condition function.
	ht1 := &hashtable.Hashtable[string, int]{} // Create a new hashtable.
	ht1.Add("apple", 5)
	ht1.Add("orange", 10)

	ht2 := &hashtable.Hashtable[string, int]{} // Create another hashtable.
	ht2.Add("orange", 8)
	ht2.Add("banana", 6)

	// Condition function to merge pairs where the value in ht2 is greater than 7.
	conditionFunc := func(key string, value int) bool {
		return value > 7
	}

	// Merge key-value pairs from ht2 into ht1 based on the condition function.
	ht1.MergeFunc(ht2, conditionFunc)

	// After merging, ht1 should contain: {"apple": 5, "orange": 8}
	expectedHashtable := &hashtable.Hashtable[string, int]{
		"apple":  5,
		"orange": 8,
	}

	// Verify that ht1 is equal to the expected hashtable.
	if !ht1.Equal(expectedHashtable) {
		t.Errorf("MergeFunc did not produce the expected result. Got: %v, Expected: %v", ht1, expectedHashtable)
	}
}

// TestMergeMany tests Hashtable.MergeMany.
func TestMergeMany(t *testing.T) {
	// Create hashtables for merging.
	ht1 := &hashtable.Hashtable[string, int]{
		"apple":  5,
		"orange": 10,
	}
	ht2 := &hashtable.Hashtable[string, int]{
		"orange": 15,
		"banana": 7,
	}
	ht3 := &hashtable.Hashtable[string, int]{
		"grape": 8,
		"melon": 12,
	}

	// Merge key-value pairs from ht2 and ht3 into ht1.
	ht1.MergeMany(ht2, ht3)

	// Expected merged hashtable.
	expectedHashtable := &hashtable.Hashtable[string, int]{
		"apple":  5,
		"orange": 15,
		"banana": 7,
		"grape":  8,
		"melon":  12,
	}

	// Check if the merged hashtable matches the expected hashtable.
	if !reflect.DeepEqual(expectedHashtable, ht1) {
		t.Errorf("Merged hashtable does not match the expected hashtable. Got: %v, Expected: %v", ht1, expectedHashtable)
	}

	// Test case for merging an empty hashtable.
	emptyHashtable := &hashtable.Hashtable[string, int]{}
	ht1.MergeMany(emptyHashtable)

	// Merged hashtable should remain unchanged.
	if !reflect.DeepEqual(expectedHashtable, ht1) {
		t.Errorf("Merged hashtable should remain unchanged after merging with an empty hashtable. Got: %v, Expected: %v", ht1, expectedHashtable)
	}
}

// TestMergeManyFunc tests Hashtable.MergeManyFunc.
func TestMergeManyFunc(t *testing.T) {
	// Test case: Merge key-value pairs based on a condition function.
	ht1 := &hashtable.Hashtable[string, int]{} // Create an empty destination hashtable.
	ht2 := &hashtable.Hashtable[string, int]{} // Create the first source hashtable.
	ht2.Add("apple", 5)
	ht3 := &hashtable.Hashtable[string, int]{} // Create the second source hashtable.
	ht3.Add("orange", 10)
	ht4 := &hashtable.Hashtable[string, int]{} // Create the third source hashtable.
	ht4.Add("banana", 7)

	// Condition function to include pairs only if the value is greater than 7.
	conditionFunc := func(i int, key string, value int) bool {
		return value >= 7
	}

	// Merge key-value pairs based on the condition function.
	mergedHashtable := ht1.MergeManyFunc([]*hashtable.Hashtable[string, int]{ht2, ht3, ht4}, conditionFunc)

	// Verify that the merged hashtable contains the expected key-value pairs.
	expectedPairs := map[string]int{
		"orange": 10,
		"banana": 7,
	}
	for key, expectedValue := range expectedPairs {
		value, ok := mergedHashtable.Get(key)
		if !ok || value != expectedValue {
			t.Errorf("Expected merged hashtable to contain key '%s': %d, but it contains '%d'", key, expectedValue, value)
		}
	}

	// Verify that unwanted pairs are not present in the merged hashtable.
	unwantedPairs := map[string]int{
		"apple": 5,
	}
	for key := range unwantedPairs {
		_, ok := mergedHashtable.Get(key)
		if ok {
			t.Errorf("Expected merged hashtable not to contain key '%s', but it was found", key)
		}
	}
}

// TestNot tests Hashtable.Not.
func TestNot(t *testing.T) {
	// Test case 1: Check if a key is not present in an empty hashtable.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	result := ht.Not("apple")                 // Check if "apple" is not in the hashtable.
	expected := true                          // "apple" is not present in the empty hashtable.

	if result != expected {
		t.Errorf("Expected result to be %v for key 'apple', but got %v", expected, result)
	}

	// Test case 2: Check if a key is not present in a non-empty hashtable.
	ht.Add("orange", 5)
	ht.Add("banana", 10)
	result = ht.Not("banana") // Check if "banana" is not in the hashtable.
	expected = false          // "banana" is present in the hashtable.

	if result != expected {
		t.Errorf("Expected result to be %v for key 'banana', but got %v", expected, result)
	}

	// Test case 3: Check if a key is not present after removing it from the hashtable.
	ht.Delete("banana") // Delete "banana" from the hashtable.
	result = ht.Not("banana")
	expected = true // "banana" is not present after removal.

	if result != expected {
		t.Errorf("Expected result to be %v for key 'banana' after removal, but got %v", expected, result)
	}
}

// TestPop tests Hashtable.Pop.
func TestPop(t *testing.T) {
	// Test case 1: Pop from an empty hashtable.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	removedValue := ht.Pop("apple")
	expectedValue := 0 // No key "apple" in the empty hashtable.

	if removedValue != expectedValue {
		t.Errorf("Expected removed value to be %d, but got %d", expectedValue, removedValue)
	}

	// Test case 2: Pop from a non-empty hashtable where the key is present.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	removedValue = ht.Pop("apple")
	expectedValue = 5 // Key "apple" exists with value 5.

	if removedValue != expectedValue {
		t.Errorf("Expected removed value to be %d, but got %d", expectedValue, removedValue)
	}
	// Verify that the key is removed.
	_, ok := ht.Get("apple")
	if ok {
		t.Errorf("Expected key 'apple' to be removed, but it was found")
	}

	// Test case 3: Pop from a non-empty hashtable where the key is not present.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	ht.Add("orange", 10)
	removedValue = ht.Pop("banana")
	expectedValue = 0 // No key "banana" in the hashtable.

	if removedValue != expectedValue {
		t.Errorf("Expected removed value to be %d, but got %d", expectedValue, removedValue)
	}
}

// TestPopOK tests Hashtable.PopOK.
func TestPopOK(t *testing.T) {
	// Test case 1: Pop from an empty hashtable.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	removedValue, ok := ht.PopOK("apple")
	if ok || removedValue != 0 {
		t.Errorf("Expected (0, false), but got (%d, %v)", removedValue, ok)
	}

	// Test case 2: Pop from a non-empty hashtable where the key is present.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	removedValue, ok = ht.PopOK("apple")
	if !ok || removedValue != 5 {
		t.Errorf("Expected (5, true), but got (%d, %v)", removedValue, ok)
	}
	// Verify that the key is removed.
	_, ok = ht.Get("apple")
	if ok {
		t.Errorf("Expected key 'apple' to be removed, but it was found")
	}

	// Test case 3: Pop from a non-empty hashtable where the key is not present.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	ht.Add("orange", 10)
	removedValue, ok = ht.PopOK("banana")
	if ok || removedValue != 0 {
		t.Errorf("Expected (0, false), but got (%d, %v)", removedValue, ok)
	}
}

func TestPopMany(t *testing.T) {
	// Test case 1: PopMany from an empty hashtable.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	removedValues := ht.PopMany("apple", "orange")
	expectedValues := &slice.Slice[int]{}
	if !reflect.DeepEqual(removedValues, expectedValues) {
		t.Errorf("Expected %v, but got %v", expectedValues, removedValues)
	}

	// Test case 2: PopMany from a non-empty hashtable where some keys are present.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	ht.Add("banana", 3)
	ht.Add("cherry", 8)
	removedValues = ht.PopMany("apple", "orange", "cherry", "grape")
	expectedValues = &slice.Slice[int]{5, 8}
	if !reflect.DeepEqual(removedValues, expectedValues) {
		t.Errorf("Expected %v, but got %v", expectedValues, removedValues)
	}
	// Verify that the keys are removed.
	_, ok := ht.Get("apple")
	if ok {
		t.Errorf("Expected key 'apple' to be removed, but it was found")
	}
	_, ok = ht.Get("orange")
	if ok {
		t.Errorf("Expected key 'orange' to be removed, but it was found")
	}
	_, ok = ht.Get("cherry")
	if ok {
		t.Errorf("Expected key 'cherry' to be removed, but it was found")
	}
	_, ok = ht.Get("grape")
	if ok {
		t.Errorf("Expected key 'grape' to be removed, but it was found")
	}
}

// TestPopManyFunc tests Hashtable.PopManyFunc.
func TestPopManyFunc(t *testing.T) {
	// Test case 1: Pop values greater than 7 from the hashtable.
	ht := &hashtable.Hashtable[string, int]{
		"apple":  5,
		"orange": 10,
		"banana": 8,
		"grape":  12,
	}

	removeCondition := func(key string, value int) bool {
		return value > 7
	}

	removedValues := ht.PopManyFunc(removeCondition)

	sort.Ints(*removedValues)

	expectedRemovedValues := &slice.Slice[int]{8, 10, 12}

	if !reflect.DeepEqual(expectedRemovedValues, removedValues) {
		t.Errorf("Expected removed values to be %v, but got %v", expectedRemovedValues, removedValues)
	}

	// Test case 2: Pop values when condition does not match any key-value pairs.
	ht = &hashtable.Hashtable[string, int]{
		"apple":  5,
		"orange": 3,
		"banana": 8,
	}

	removeCondition = func(key string, value int) bool {
		return value > 10
	}

	removedValues = ht.PopManyFunc(removeCondition)

	expectedRemovedValues = &slice.Slice[int]{} // No values match the condition.

	if !reflect.DeepEqual(expectedRemovedValues, removedValues) {
		t.Errorf("Expected removed values to be %v, but got %v", expectedRemovedValues, removedValues)
	}

	// Test case 3: Pop values from an empty hashtable.
	ht = &hashtable.Hashtable[string, int]{}

	removeCondition = func(key string, value int) bool {
		return value > 0
	}

	removedValues = ht.PopManyFunc(removeCondition)

	expectedRemovedValues = &slice.Slice[int]{} // No values to remove from an empty hashtable.

	if !reflect.DeepEqual(expectedRemovedValues, removedValues) {
		t.Errorf("Expected removed values to be %v, but got %v", expectedRemovedValues, removedValues)
	}
}

// TestReplaceMany tests Hashtable.ReplaceMany.
func TestUpdate(t *testing.T) {
	// Test case 1: Replace with an empty hashtable and a function that never modifies any pairs.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	replaceFunc := func(key string, value int) (int, bool) {
		return value, false // Never modify any values
	}
	ht.ReplaceMany(replaceFunc)
	expected := &hashtable.Hashtable[string, int]{} // Expected empty hashtable.
	if !reflect.DeepEqual(ht, expected) {
		t.Errorf("Expected %v, but got %v", expected, ht)
	}

	// Test case 2: Replace with a non-empty hashtable and a function that never modifies any pairs.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	ht.Add("orange", 10)
	replaceFunc = func(key string, value int) (int, bool) {
		return value, false // Never modify any values
	}
	ht.ReplaceMany(replaceFunc)
	expected = &hashtable.Hashtable[string, int]{"apple": 5, "orange": 10} // Expected same hashtable.
	if !reflect.DeepEqual(ht, expected) {
		t.Errorf("Expected %v, but got %v", expected, ht)
	}

	// Test case 3: Replace with a non-empty hashtable and a function that modifies certain pairs.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	ht.Add("orange", 10)
	replaceFunc = func(key string, value int) (int, bool) {
		if key == "apple" {
			return value * 2, true // Modify the value for the "apple" key
		}
		return value, false // Leave other values unchanged
	}
	ht.ReplaceMany(replaceFunc)
	expected = &hashtable.Hashtable[string, int]{"apple": 10, "orange": 10} // Expected modified hashtable.
	if !reflect.DeepEqual(ht, expected) {
		t.Errorf("Expected %v, but got %v", expected, ht)
	}
}

// TestTakeFrom tests Hashtable.TakeFrom.
func TestTakeFrom(t *testing.T) {
	// Test case 1: Transfer from an empty hashtable to another empty hashtable.
	ht1 := &hashtable.Hashtable[string, int]{} // Create an empty destination hashtable.
	ht2 := &hashtable.Hashtable[string, int]{} // Create an empty source hashtable.
	ht1.TakeFrom(ht2)                          // Transfer from ht2 to ht1.

	// Verify that ht1 is still empty.
	if ht1.Length() != 0 {
		t.Errorf("Expected destination hashtable to be empty after transfer, but it has %d items", ht1.Length())
	}

	// Verify that ht2 is still empty.
	if ht2.Length() != 0 {
		t.Errorf("Expected source hashtable to be empty after transfer, but it has %d items", ht2.Length())
	}

	// Test case 2: Transfer from a non-empty hashtable to an empty hashtable.
	ht1 = &hashtable.Hashtable[string, int]{} // Create an empty destination hashtable.
	ht2 = &hashtable.Hashtable[string, int]{} // Create a source hashtable.
	ht2.Add("orange", 10)
	ht1.TakeFrom(ht2) // Transfer from ht2 to ht1.

	// Verify that ht1 contains the transferred key-value pair.
	expectedValue := 10
	transferredValue, ok := ht1.Get("orange")
	if !ok || transferredValue != expectedValue {
		t.Errorf("Expected destination hashtable to contain 'orange': %d after transfer, but it contains '%d'", expectedValue, transferredValue)
	}

	// Verify that ht2 is empty after transfer.
	if ht2.Length() != 0 {
		t.Errorf("Expected source hashtable to be empty after transfer, but it has %d items", ht2.Length())
	}
}

// TestValues tests Hashtable.Values.
func TestValues(t *testing.T) {
	// Test case 1: Values of an empty hashtable should be an empty slice.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	values := ht.Values()
	expected := &slice.Slice[int]{} // Expected empty slice.
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, but got %v", expected, values)
	}

	// Test case 2: Values of a non-empty hashtable.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	ht.Add("orange", 10)
	values = ht.Values()
	sort.Ints(*values)
	expected = &slice.Slice[int]{5, 10} // Expected slice containing [5, 10].
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, but got %v", expected, values)
	}

	// Test case 3: Values of a hashtable with multiple entries.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	ht.Add("orange", 10)
	ht.Add("banana", 15)
	values = ht.Values()

	sort.Ints(*values)
	expected = &slice.Slice[int]{5, 10, 15} // Expected slice containing [5, 10, 15].
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, but got %v", expected, values)
	}
}

// TestValuesFunc tests Hashtable.ValuesFunc.
func TestValuesFunc(t *testing.T) {
	// Test case 1: ValuesFunc with an empty hashtable and a condition that never satisfies.
	ht := &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	filterFunc := func(key string, value int) bool {
		return value > 7 // Include values greater than 7 in the result
	}
	values := ht.ValuesFunc(filterFunc)
	expected := &slice.Slice[int]{} // Expected empty slice.
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, but got %v", expected, values)
	}

	// Test case 2: ValuesFunc with a non-empty hashtable and a condition that never satisfies.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	ht.Add("orange", 2)
	filterFunc = func(key string, value int) bool {
		return value > 7 // Include values greater than 7 in the result
	}
	values = ht.ValuesFunc(filterFunc)
	expected = &slice.Slice[int]{} // Expected empty slice.
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, but got %v", expected, values)
	}

	// Test case 3: ValuesFunc with a non-empty hashtable and a condition that satisfies for some values.
	ht = &hashtable.Hashtable[string, int]{} // Create an empty hashtable.
	ht.Add("apple", 5)
	ht.Add("orange", 10)
	ht.Add("banana", 15)
	filterFunc = func(key string, value int) bool {
		return value > 7 // Include values greater than 7 in the result
	}
	values = ht.ValuesFunc(filterFunc)
	sort.Ints(*values)
	expected = &slice.Slice[int]{10, 15} // Expected slice containing [10, 15].
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Expected %v, but got %v", expected, values)
	}
}
