# Gomap
Gomap is a generic map implementation in [Go](https://github.com/golang/go), supporting dynamic key-value pairs of any data type. It offers essential operations like adding, deleting, and checking for key-value pairs, along with advanced functionalities such as iteration, merging, intersection, and conditional mapping. Its flexibility allows seamless manipulation and querying, making it a powerful tool for various applications.

![Gomap](https://repository-images.githubusercontent.com/192931263/82f0f49f-fe0c-46f7-b52d-d1ded90f4204)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/lindsaygelle/gomap)](https://pkg.go.dev/github.com/lindsaygelle/gomap)
[![Go Report Card](https://goreportcard.com/badge/github.com/lindsaygelle/gomap)](https://goreportcard.com/report/github.com/lindsaygelle/gomap)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/lindsaygelle/gomap)](https://github.com/lindsaygelle/gomap/releases)
[![GitHub](https://img.shields.io/github/license/lindsaygelle/gomap)](LICENSE.txt)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)

## Features

### 🔑 Key Operations

Effortlessly manage key-value pairs by adding, deleting, and replacing them. Employ custom functions to pop values and make replacements seamlessly.

### 🔍 Query Operations

Determine key existence, retrieve corresponding values, and identify specific values within the hash table

### 🔄 Iteration

Navigate through keys and values smoothly, applying functions and breaking based on specified conditions.

### 🔗 Set Operations

Effortlessly compute intersections with other hash tables. Merge hash tables with ease and apply custom merging functions to create cohesive datasets.

### 🎛️ Functional Operations

Filter data based on conditions and map values efficiently using custom functions. Compare hash tables for equality or define your own custom equality functions to ensure accurate results.

### 🛠️ Convenience Functions

Retrieve keys or values as slices for easy handling. Quickly check hash table emptiness, presence, and retrieve its length.

## Installation
You can install it in your Go project using `go get`:

```sh
go get github.com/lindsaygelle/gomap
```

## Usage
Import the package into your Go code:


```Go
import (
	"github.com/lindsaygelle/gomap"
)
```

## Methods
Provided methods for `&gomap.Map[K]V`.

### Add
Adds a key-value pair to the hash table and returns the updated hash table.

```Go
myMap := &gomap.Map[string, int]{}
myMap.Add("key1", 1)
myMap.Add("key2", 2)
fmt.Println(myMap) // &map[key1:1 key2:2]
```

### AddLength
Adds a key-value pair to the hash table and returns the new length of the hash table.

```Go
myMap := &gomap.Map[string, int]{}
length := myMap.AddLength("key1", 1)
fmt.Println(length) // 1
```

### AddMany
Adds multiple key-value pairs to the hash table and returns the updated hash table.

```Go
myMap := &gomap.Map[string, int]{}
myMap.AddMany(map[string]int{"key1": 1, "key2": 2})
fmt.Println(myMap) // &map[key1:1 key2:2]
```

### AddManyFunc
Adds key-value pairs from a slice of maps to the hash table using a custom function and returns the updated hash table.

```Go
myMap := &gomap.Map[string, int]{}
myMaps := []map[string]int{{"key1": 1}, {"key2": 2}}
myMap.AddManyFunc(myMaps, func(i int, key string, value int) bool {
    return true // Add all key-value pairs
})
fmt.Println(myMap) // &map[key1:1 key2:2]
```

### AddManyOK
Adds multiple key-value pairs to the hash table and returns a slice indicating successful additions.

```Go
myMap := &gomap.Map[string, int]{}
results := myMap.AddManyOK(map[string]int{"key1": 1, "key2": 2})
fmt.Println(results) // &[true, true]
```

### AddOK
Adds a key-value pair to the hash table and returns true if the addition was successful, false otherwise.

```Go
myMap := &gomap.Map[string, int]{}
added := myMap.AddOK("key1", 1)
fmt.Println(added) // true
```

### AddValueFunc
Adds a key-value pair to the map using a function to determine the key from the given value.

```Go
myMap := &gomap.Map[string, int]{}
myMap.AddValueFunc(5, func(value int) string {
	return strconv.Itoa(value)
})
fmt.Println(myMap) // &map[5:5]
```

### AddValuesFunc
Adds multiple key-value pairs to the map using a function to determine keys from the given values.

```Go
myMap := &gomap.Map[string, int]{}
myMap.AddValuesFunc([]int{5, 10, 15}, func(i int, value int) string {
	return strconv.Itoa(value) + strconv.Itoa(i)
})
fmt.Println(myMap) // &map[50:5 101:10 152:15]
```

### Contains
Checks if the given value is present in the hash table and returns the corresponding key along with a boolean indicating existence.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
key, exists := myMap.Contains(2)
fmt.Println(key, exists) // key2 true
```

### Delete
Removes the specified key and its associated value from the hash table and returns the updated hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap.Delete("key1")
fmt.Println(myMap) // &map[key2:2]
```

### DeleteLength
Removes the specified key and its associated value from the hash table and returns the new length of the hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
length := myMap.DeleteLength("key1")
fmt.Println(length) // 1
```

### DeleteMany
Removes multiple keys and their associated values from the hash table and returns the updated hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap.DeleteMany("key1", "key2")
fmt.Println(myMap) // &map[]
```

### DeleteManyFunc
Removes key-value pairs from the hash table using a custom function and returns the updated hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap.DeleteManyFunc(func(key string, value int) bool {
    return key == "key1"
})
fmt.Println(myMap) // &map[key2:2]
```

### DeleteManyOK
Removes multiple keys and their associated values from the hash table and returns a slice indicating successful deletions.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
results := myMap.DeleteManyOK("key1", "key2")
fmt.Println(results) // &[true, true]
```

### DeleteManyValues
Removes multiple values from the hash table and returns the updated hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap.DeleteManyValues(1, 2)
fmt.Println(myMap) // &map[]
```

### DeleteOK
Removes the specified key and its associated value from the hash table and returns true if the deletion was successful, false otherwise.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
deleted := myMap.DeleteOK("key1")
fmt.Println(deleted) // true
```

### Each
Applies the given function to each key-value pair in the hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap.Each(func(key string, value int) {
    fmt.Println(key, value)
})
// Output:
// key1 1
// key2 2
```

### EachBreak
Applies the given function to each key-value pair in the hash table and breaks the iteration if the function returns false.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap.EachBreak(func(key string, value int) bool {
    fmt.Println(key, value)
    return key != "key1"
})
// Output:
// key1 1
```

### EachKey
Applies the given function to each key in the hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap.EachKey(func(key string) {
    fmt.Println(key)
})
// Output:
// key1
// key2
```

### EachKeyBreak
Applies the given function to each key in the hash table and breaks the iteration if the function returns false.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap.EachKeyBreak(func(key string) bool {
    fmt.Println(key)
    return key != "key1"
})
// Output:
// key1
```

### EachValue
Applies the given function to each value in the hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap.EachValue(func(value int) {
    fmt.Println(value)
})
// Output:
// 1
// 2
```

### EachValueBreak
Applies the given function to each value in the hash table and breaks the iteration if the function returns false.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap.EachValueBreak(func(value int) bool {
    fmt.Println(value)
    return value != 1
})
// Output:
// 1
```

### EmptyInto
Empties the current hash table and inserts its content into another hash table. It returns the updated destination hash table.

```Go
source := &gomap.Map[string, int]{"key1": 1, "key2": 2}
destination := &gomap.Map[string, int]{"key3": 3}
destination = source.EmptyInto(destination)
fmt.Println(destination) // &map[key1:1 key2:2]
```

### Equal
Checks if the current hash table is equal to another hash table.

```Go
myMap1 := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap2 := &gomap.Map[string, int]{"key1": 1, "key2": 2}
isEqual := myMap1.Equal(myMap2)
fmt.Println(isEqual) // true
```

### EqualFunc
Checks if the current hash table is equal to another hash table using a custom comparison function.

```Go
myMap1 := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap2 := &gomap.Map[string, int]{"key1": 1, "key2": 4}
isEqual := myMap1.EqualFunc(myMap2, func(a int, b int) bool {
    return a == b || a%2 == 0 && b%2 == 0
})
fmt.Println(isEqual) // true (custom comparison allows for even values)
```

### EqualLength
Checks if the current hash table is equal in length to another hash table.

```Go
myMap1 := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap2 := &gomap.Map[string, int]{"key1": 1, "key2": 2, "key3": 3}
isEqualLength := myMap1.EqualLength(myMap2)
fmt.Println(isEqualLength) // false (different lengths)
```

### Fetch
Retrieves the value associated with the given key from the hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
value := myMap.Fetch("key1")
fmt.Println(value) // 1
```

### Filter
Applies the given function to each key-value pair in the hash table and retains pairs for which the function returns true.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2, "key3": 3}
filteredHashtable := myMap.Filter(func(key string, value int) bool {
    return value%2 == 0
})
fmt.Println(filteredHashtable) // &map[key2:2]
```

### Get
Retrieves the value associated with the given key from the hash table and returns it along with a boolean indicating existence.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
value, exists := myMap.Get("key2")
fmt.Println(value, exists) // 2 true
```

### GetMany
Retrieves values associated with multiple keys from the hash table and returns them in a slice.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2, "key3": 3}
values := myMap.GetMany("key1", "key3")
fmt.Println(values) // &[1 3]
```

### Has
Checks if the given key is present in the hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
exists := myMap.Has("key1")
fmt.Println(exists) // true
```

### HasMany
Checks if multiple keys are present in the hash table and returns a slice indicating their existence.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2, "key3": 3}
exists := myMap.HasMany("key1", "key4")
fmt.Println(exists) // &[true false]
```

### Intersection
Returns a new hash table containing key-value pairs that are present in both the current and another hash table.

```Go
myMap1 := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap2 := &gomap.Map[string, int]{"key1": 1, "key3": 3}
intersection := myMap1.Intersection(myMap2)
fmt.Println(intersection) // &map[key1:1]
```

### IntersectionFunc
Returns a new hash table containing key-value pairs that are present in both the current and another hash table, determined by a custom function.

```Go
myMap1 := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap2 := &gomap.Map[string, int]{"key1": 1, "key3": 3}
intersection := myMap1.IntersectionFunc(myMap2, func(key string, a int, b int) bool {
    return a == b
})
fmt.Println(intersection) // &map[key1:1]
```

### IsEmpty
Checks if the hash table is empty.

```Go
myMap := &gomap.Map[string, int]{}
isEmpty := myMap.IsEmpty()
fmt.Println(isEmpty) // true
```

### IsPopulated
Checks if the hash table contains key-value pairs.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
isPopulated := myMap.IsPopulated()
fmt.Println(isPopulated) // true
```

### Keys
Returns a slice containing all keys in the hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
keys := myMap.Keys()
fmt.Println(keys) // &["key1" "key2"]
```

### KeysFunc
Returns a slice containing keys that satisfy the given function.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2, "key3": 3}
keys := myMap.KeysFunc(func(key string) bool {
    return key != "key3"
})
fmt.Println(keys) // &["key1" "key2"]
```

### Length
Returns the number of key-value pairs in the hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
length := myMap.Length()
fmt.Println(length) // 2
```

### Map
Applies the given function to each key-value pair in the hash table and replaces the value with the result of the function.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap.Map(func(key string, value int) int {
    return value * 2
})
fmt.Println(myMap) // &map[key1:2 key2:4]
```

### MapBreak
Applies the given function to each key-value pair in the hash table and replaces the value with the result of the function. It breaks the iteration if the function returns false.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap.MapBreak(func(key string, value int) (int, bool) {
    return value * 2, key != "key2"
})
fmt.Println(myMap) // &map[key1:2 key2:2]
```

### Merge
Merges the current hash table with another hash table and returns the updated hash table.

```Go
myMap1 := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap2 := &gomap.Map[string, int]{"key2": 3, "key3": 4}
mergedHashtable := myMap1.Merge(myMap2)
fmt.Println(mergedHashtable) // &map[key1:1 key2:3 key3:4]
```

### MergeFunc
Merges the current hash table with another hash table using a custom function and returns the updated hash table.

```Go
myMap1 := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap2 := &gomap.Map[string, int]{"key2": 3, "key3": 4}
mergedHashtable := myMap1.MergeFunc(myMap2, func(key string, value1 int, value2 int) bool {
    return value1 > value2
})
fmt.Println(mergedHashtable) // &map[key1:1 key2:2 key3:4]
```

### MergeMany
Merges the current hash table with multiple other hash tables and returns the updated hash table.

```Go
myMap1 := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap2 := &gomap.Map[string, int]{"key2": 3, "key3": 4}
mergedHashtable := myMap1.MergeMany(myMap2, &gomap.Map[string, int]{"key4": 5})
fmt.Println(mergedHashtable) // &map[key1:1 key2:3 key3:4 key4:5]
```

### MergeManyFunc
Merges the current hash table with multiple other hash tables using a custom function and returns the updated hash table.

```Go
myMap1 := &gomap.Map[string, int]{"key1": 1, "key2": 2}
myMap2 := &gomap.Map[string, int]{"key2": 3, "key3": 4}
mergedHashtable := myMap1.MergeManyFunc([]*gomap.Map[string, int]{myMap2}, func(i int, key string, value int) bool {
    return key != "key3"
})
fmt.Println(mergedHashtable) // &map[key1:1 key2:3]
```

### Not
Checks if the given key is not present in the hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
notPresent := myMap.Not("key3")
fmt.Println(notPresent) // true
```

### NotMany
Checks if multiple keys are not present in the hash table and returns a slice indicating their absence.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
absentKeys := myMap.NotMany("key3", "key4")
fmt.Println(absentKeys) // &[true true]
```

### Pop
Removes the specified key and its associated value from the hash table and returns the value.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
value := myMap.Pop("key1")
fmt.Println(value) // 1
fmt.Println(myMap) // &map[key2:2]
```

### PopMany
Removes multiple keys and their associated values from the hash table and returns the values in a slice.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2, "key3": 3}
values := myMap.PopMany("key1", "key3")
fmt.Println(values) // &[1 3]
fmt.Println(myMap) // &map[key2:2]
```

### PopManyFunc
Removes key-value pairs from the hash table using a custom function and returns the values in a slice.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2, "key3": 3}
values := myMap.PopManyFunc(func(key string, value int) bool {
    return key != "key2"
})
fmt.Println(values) // &[1 3]
fmt.Println(myMap) // &map[key2:2]
```

### PopOK
Removes the specified key and its associated value from the hash table and returns the value along with a boolean indicating existence.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
value, exists := myMap.PopOK("key1")
fmt.Println(value, exists) // 1 true
fmt.Println(myMap)  // &map[key2:2]
```

### ReplaceMany
Applies the given function to each key-value pair in the hash table and replaces the value if the function returns true. It returns the updated hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2, "key3": 3}
myMap.ReplaceMany(func(key string, value int) (int, bool) {
    if key == "key2" {
        return value * 2, true
    }
    return value, false
})
fmt.Println(myMap) // &map[key1:1 key2:4 key3:3]
```

### TakeFrom
Empties the current hash table and inserts its content into another hash table. It returns the updated destination hash table.

```Go
source := &gomap.Map[string, int]{"key1": 1, "key2": 2}
destination := &gomap.Map[string, int]{"key3": 3}
destination = source.TakeFrom(destination)
fmt.Println(destination) // &map[key1:1 key2:2]
```

### Values
Returns a slice containing all values in the hash table.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2}
values := myMap.Values()
fmt.Println(values) // &[1 2]
```

### ValuesFunc
Returns a slice containing values that satisfy the given function.

```Go
myMap := &gomap.Map[string, int]{"key1": 1, "key2": 2, "key3": 3}
values := myMap.ValuesFunc(func(key string, value int) bool {
    return value%2 == 0
})
fmt.Println(values) // &[2]
```

## Examples

### Struct
Using a custom struct.

```Go
package main

import (
	"fmt"
	"github.com/lindsaygelle/gomap" // Import the gomap package
)

type Animal struct {
	Age     int
	Name    string
	Species string
}

func main() {
	// Create a new Animal
	fluffy := Animal{Age: 1, Name: "Fluffy", Species: "Cat"}

	// Create a new Map of Animals
	animals := &gomap.Map[string, Animal]{}

	// Add an Animal to the Map using Add function
	animals.Add(fluffy.Name, fluffy)

	// Check if the Map contains a specific key
	if animals.Has("Fluffy") {
		fmt.Println("Fluffy is in the Map!")
	}

	// Get the value associated with a specific key
	if animal, exists := animals.Get("Fluffy"); exists {
		fmt.Println("Found Fluffy in the Map:", animal)
	}

	// Update the age of Fluffy
	updatedFluffy := Animal{Age: 2, Name: "Fluffy", Species: "Cat"}
	animals.Add("Fluffy", updatedFluffy)

	// Delete an entry from the Map
	animals.Delete("Fluffy")

	// Add more animals after deleting Fluffy
	animals.Add("Buddy", Animal{Age: 3, Name: "Buddy", Species: "Dog"})
	animals.Add("Whiskers", Animal{Age: 2, Name: "Whiskers", Species: "Rabbit"})

	// Iterate over the Map and print each key-value pair
	animals.Each(func(key string, value Animal) {
		fmt.Println("Key:", key, "Value:", value)
	})

	// Check if the Map is empty
	if animals.IsEmpty() {
		fmt.Println("Map is empty.")
	} else {
		fmt.Println("Map is not empty.")
	}

	// Get all keys from the Map
	keys := animals.Keys().Slice()
	fmt.Println("Keys in Map:", keys)

	// Get all values from the Map
	values := animals.Values().Slice()
	fmt.Println("Values in Map:", values)
}
```

### Chaining
Chaining methods together.


```Go
package main

import (
	"fmt"
	"github.com/lindsaygelle/gomap" // Import the gomap package
)

type Animal struct {
	Age     int
	Name    string
	Species string
}

func main() {
	// Create a new Map of Animals and add animals using method chaining
	animals := &gomap.Map[string, Animal]{}.
		Add("Fluffy", Animal{Age: 1, Name: "Fluffy", Species: "Cat"}).
		Add("Buddy", Animal{Age: 3, Name: "Buddy", Species: "Dog"}).
		Add("Whiskers", Animal{Age: 2, Name: "Whiskers", Species: "Rabbit"})

	// Print the number of animals in the Map using the Length() method
	fmt.Println("Number of animals in the Map:", animals.Length())

	// Check if a specific animal is in the Map using method chaining
	if exists := animals.Has("Fluffy"); exists {
		fmt.Println("Fluffy is in the Map!")
	}

	// Retrieve and print the age of a specific animal using method chaining
	if age, exists := animals.Get("Buddy"); exists {
		fmt.Println("Buddy's age is:", age.Age)
	}

	// Iterate over the Map and print each key-value pair using method chaining
	animals.Each(func(key string, value Animal) {
		fmt.Println("Key:", key, "Value:", value)
	})

	// Chain filtering and mapping methods to filter out specific animals and update their ages
	animals.
		Filter(func(key string, value Animal) bool {
			return value.Age > 1
		}).
		Map(func(key string, value Animal) Animal {
			value.Age++
			return value
		})

	// Print the updated ages of animals after filtering and mapping using method chaining
	fmt.Println("Ages of animals after filtering and mapping:")
	animals.Each(func(key string, value Animal) {
		fmt.Println("Key:", key, "Value:", value)
	})
}
```

## Docker
A [Dockerfile](./Dockerfile) is provided for individuals that prefer containerized development.

### Building
Building the Docker container:
```sh
docker build . -t gomap
```

### Running
Developing and running Go within the Docker container:
```sh
docker run -it --rm --name gomap gomap
```

## Docker Compose
A [docker-compose](./docker-compose.yml) file has also been included for convenience:
### Running
Running the compose file.
```sh
docker-compose up -d
```

## Contributing
We warmly welcome contributions to Map. Whether you have innovative ideas, bug reports, or enhancements in mind, please share them with us by submitting GitHub issues or creating pull requests. For substantial contributions, it's a good practice to start a discussion by creating an issue to ensure alignment with the project's goals and direction. Refer to the [CONTRIBUTING](./CONTRIBUTING.md) file for comprehensive details.

## Branching
For a smooth collaboration experience, we have established branch naming conventions and guidelines. Please consult the [BRANCH_NAMING_CONVENTION](./BRANCH_NAMING_CONVENTION.md) document for comprehensive information and best practices.

## License
Map is released under the MIT License, granting you the freedom to use, modify, and distribute the code within this repository in accordance with the terms of the license. For additional information, please review the [LICENSE](./LICENSE) file.

## Security
If you discover a security vulnerability within this project, please consult the [SECURITY](./SECURITY.md) document for information and next steps.

## Code Of Conduct
This project has adopted the [Amazon Open Source Code of Conduct](https://aws.github.io/code-of-conduct). For additional information, please review the [CODE_OF_CONDUCT](./CODE_OF_CONDUCT.md) file.

## Acknowledgements
Big thanks to [egonelbre/gophers](https://github.com/egonelbre/gophers) for providing the delightful Gopher artwork used in the social preview. Don't hesitate to pay them a visit!
