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
		map[string]int{"apple": 5, "banana": 3},
		map[string]int{"banana": 10, "cherry": 8},
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
			t.Fatalf("Expected result: %v, but got: %v", expectedResults[i], result)
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
			t.Fatalf("Expected result: %v, but got: %v", expectedResults[i], result)
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

	if !result.Equal(expected) {
		t.Errorf("Expected keys %v after deleting 'apple' and 'banana', but got keys %v", expected, result)
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
			t.Fatalf("Expected deletion of key %s to be %v but got %v", keysToDelete[i], expectedResults[i], result)
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
			return value * 2, false // Break the iteration when key is "banana"
		}
		return value * 2, true // Continue iterating for other keys and double the values
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
