# Hashtable
Hashtable is a [Go](https://github.com/golang/go) package that provides a generic hashtable with extended functionality. It abstracts common map operations, such as adding, deleting, iterating, and more, making it easier to work with maps in Go.

![Hashtable]()

[![PkgGoDev](https://pkg.go.dev/badge/github.com/lindsaygelle/hashtable)](https://pkg.go.dev/github.com/lindsaygelle/hashtable)
[![Go Report Card](https://goreportcard.com/badge/github.com/lindsaygelle/hashtable)](https://goreportcard.com/report/github.com/lindsaygelle/hashtable)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/lindsaygelle/hashtable)](https://github.com/lindsaygelle/hashtable/releases)
[![GitHub](https://img.shields.io/github/license/lindsaygelle/hashtable)](LICENSE.txt)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)

## Features

## Installation
You can install it in your Go project using `go get`:

```sh
go get github.com/lindsaygelle/hashtable
```

## Usage
Import the package into your Go code:


```Go
import (
	"github.com/lindsaygelle/hashtable"
)
```

## Methods
Provided methods for `&hashtable.Hashtable[K]V`.

### Add
Adds a key-value pair to the hash table and returns the updated hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{}
myHashtable.Add("key1", 1)
myHashtable.Add("key2", 2)
fmt.Println(myHashtable) // &map[key1:1 key2:2]
```

### AddLength
Adds a key-value pair to the hash table and returns the new length of the hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{}
length := myHashtable.AddLength("key1", 1)
fmt.Println(length) // 1
```

### AddMany
Adds multiple key-value pairs to the hash table and returns the updated hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{}
myHashtable.AddMany(map[string]int{"key1": 1, "key2": 2})
fmt.Println(myHashtable) // &map[key1:1 key2:2]
```

### AddManyFunc
Adds key-value pairs from a slice of maps to the hash table using a custom function and returns the updated hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{}
myMaps := []map[string]int{{"key1": 1}, {"key2": 2}}
myHashtable.AddManyFunc(myMaps, func(i int, key string, value int) bool {
    return true // Add all key-value pairs
})
fmt.Println(myHashtable) // &map[key1:1 key2:2]
```

### AddManyOK
Adds multiple key-value pairs to the hash table and returns a slice indicating successful additions.

```Go
myHashtable := &hashtable.Hashtable[string, int]{}
results := myHashtable.AddManyOK(map[string]int{"key1": 1, "key2": 2})
fmt.Println(results) // &[true, true]
```

### AddOK
Adds a key-value pair to the hash table and returns true if the addition was successful, false otherwise.

```Go
myHashtable := &hashtable.Hashtable[string, int]{}
added := myHashtable.AddOK("key1", 1)
fmt.Println(added) // true
```

### Contains
Checks if the given value is present in the hash table and returns the corresponding key along with a boolean indicating existence.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
key, exists := myHashtable.Contains(2)
fmt.Println(key, exists) // key2 true
```

### Delete
Removes the specified key and its associated value from the hash table and returns the updated hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable.Delete("key1")
fmt.Println(myHashtable) // &map[key2:2]
```

### DeleteLength
Removes the specified key and its associated value from the hash table and returns the new length of the hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
length := myHashtable.DeleteLength("key1")
fmt.Println(length) // 1
```

### DeleteMany
Removes multiple keys and their associated values from the hash table and returns the updated hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable.DeleteMany("key1", "key2")
fmt.Println(myHashtable) // &map[]
```

### DeleteManyFunc
Removes key-value pairs from the hash table using a custom function and returns the updated hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable.DeleteManyFunc(func(key string, value int) bool {
    return key == "key1"
})
fmt.Println(myHashtable) // &map[key2:2]
```

### DeleteManyOK
Removes multiple keys and their associated values from the hash table and returns a slice indicating successful deletions.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
results := myHashtable.DeleteManyOK("key1", "key2")
fmt.Println(results) // &[true, true]
```

### DeleteManyValues
Removes multiple values from the hash table and returns the updated hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable.DeleteManyValues(1, 2)
fmt.Println(myHashtable) // &map[]
```

### DeleteOK
Removes the specified key and its associated value from the hash table and returns true if the deletion was successful, false otherwise.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
deleted := myHashtable.DeleteOK("key1")
fmt.Println(deleted) // true
```

### Each
Applies the given function to each key-value pair in the hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable.Each(func(key string, value int) {
    fmt.Println(key, value)
})
// Output:
// key1 1
// key2 2
```

### EachBreak
Applies the given function to each key-value pair in the hash table and breaks the iteration if the function returns false.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable.EachBreak(func(key string, value int) bool {
    fmt.Println(key, value)
    return key != "key1"
})
// Output:
// key1 1
```

### EachKey
Applies the given function to each key in the hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable.EachKey(func(key string) {
    fmt.Println(key)
})
// Output:
// key1
// key2
```

### EachKeyBreak
Applies the given function to each key in the hash table and breaks the iteration if the function returns false.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable.EachKeyBreak(func(key string) bool {
    fmt.Println(key)
    return key != "key1"
})
// Output:
// key1
```

### EachValue
Applies the given function to each value in the hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable.EachValue(func(value int) {
    fmt.Println(value)
})
// Output:
// 1
// 2
```

### EachValueBreak
Applies the given function to each value in the hash table and breaks the iteration if the function returns false.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable.EachValueBreak(func(value int) bool {
    fmt.Println(value)
    return value != 1
})
// Output:
// 1
```

### EmptyInto
Empties the current hash table and inserts its content into another hash table. It returns the updated destination hash table.

```Go
source := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
destination := &hashtable.Hashtable[string, int]{"key3": 3}
destination = source.EmptyInto(destination)
fmt.Println(destination) // &map[key1:1 key2:2]
```

### Equal
Checks if the current hash table is equal to another hash table.

```Go
myHashtable1 := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable2 := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
isEqual := myHashtable1.Equal(myHashtable2)
fmt.Println(isEqual) // true
```

### EqualFunc
Checks if the current hash table is equal to another hash table using a custom comparison function.

```Go
myHashtable1 := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable2 := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 4}
isEqual := myHashtable1.EqualFunc(myHashtable2, func(a int, b int) bool {
    return a == b || a%2 == 0 && b%2 == 0
})
fmt.Println(isEqual) // true (custom comparison allows for even values)
```

### EqualLength
Checks if the current hash table is equal in length to another hash table.

```Go
myHashtable1 := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable2 := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2, "key3": 3}
isEqualLength := myHashtable1.EqualLength(myHashtable2)
fmt.Println(isEqualLength) // false (different lengths)
```

### Fetch
Retrieves the value associated with the given key from the hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
value := myHashtable.Fetch("key1")
fmt.Println(value) // 1
```

### Filter
Applies the given function to each key-value pair in the hash table and retains pairs for which the function returns true.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2, "key3": 3}
filteredHashtable := myHashtable.Filter(func(key string, value int) bool {
    return value%2 == 0
})
fmt.Println(filteredHashtable) // &map[key2:2]
```

### Get
Retrieves the value associated with the given key from the hash table and returns it along with a boolean indicating existence.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
value, exists := myHashtable.Get("key2")
fmt.Println(value, exists) // 2 true
```

### GetMany
Retrieves values associated with multiple keys from the hash table and returns them in a slice.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2, "key3": 3}
values := myHashtable.GetMany("key1", "key3")
fmt.Println(values) // &[1 3]
```

### Has
Checks if the given key is present in the hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
exists := myHashtable.Has("key1")
fmt.Println(exists) // true
```

### HasMany
Checks if multiple keys are present in the hash table and returns a slice indicating their existence.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2, "key3": 3}
exists := myHashtable.HasMany("key1", "key4")
fmt.Println(exists) // &[true false]
```

### Intersection
Returns a new hash table containing key-value pairs that are present in both the current and another hash table.

```Go
myHashtable1 := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable2 := &hashtable.Hashtable[string, int]{"key1": 1, "key3": 3}
intersection := myHashtable1.Intersection(myHashtable2)
fmt.Println(intersection) // &map[key1:1]
```

### IntersectionFunc
Returns a new hash table containing key-value pairs that are present in both the current and another hash table, determined by a custom function.

```Go
myHashtable1 := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable2 := &hashtable.Hashtable[string, int]{"key1": 1, "key3": 3}
intersection := myHashtable1.IntersectionFunc(myHashtable2, func(key string, a int, b int) bool {
    return a == b
})
fmt.Println(intersection) // &map[key1:1]
```

### IsEmpty
Checks if the hash table is empty.

```Go
myHashtable := &hashtable.Hashtable[string, int]{}
isEmpty := myHashtable.IsEmpty()
fmt.Println(isEmpty) // true
```

### IsPopulated
Checks if the hash table contains key-value pairs.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
isPopulated := myHashtable.IsPopulated()
fmt.Println(isPopulated) // true
```

### Keys
Returns a slice containing all keys in the hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
keys := myHashtable.Keys()
fmt.Println(keys) // &["key1" "key2"]
```

### KeysFunc
Returns a slice containing keys that satisfy the given function.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2, "key3": 3}
keys := myHashtable.KeysFunc(func(key string) bool {
    return key != "key3"
})
fmt.Println(keys) // &["key1" "key2"]
```

### Length
Returns the number of key-value pairs in the hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
length := myHashtable.Length()
fmt.Println(length) // 2
```

### Map
Applies the given function to each key-value pair in the hash table and replaces the value with the result of the function.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable.Map(func(key string, value int) int {
    return value * 2
})
fmt.Println(myHashtable) // &map[key1:2 key2:4]
```

### MapBreak
Applies the given function to each key-value pair in the hash table and replaces the value with the result of the function. It breaks the iteration if the function returns false.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable.MapBreak(func(key string, value int) (int, bool) {
    return value * 2, key != "key2"
})
fmt.Println(myHashtable) // &map[key1:2 key2:2]
```

### Merge
Merges the current hash table with another hash table and returns the updated hash table.

```Go
myHashtable1 := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable2 := &hashtable.Hashtable[string, int]{"key2": 3, "key3": 4}
mergedHashtable := myHashtable1.Merge(myHashtable2)
fmt.Println(mergedHashtable) // &map[key1:1 key2:3 key3:4]
```

### MergeFunc
Merges the current hash table with another hash table using a custom function and returns the updated hash table.

```Go
myHashtable1 := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable2 := &hashtable.Hashtable[string, int]{"key2": 3, "key3": 4}
mergedHashtable := myHashtable1.MergeFunc(myHashtable2, func(key string, value1 int, value2 int) bool {
    return value1 > value2
})
fmt.Println(mergedHashtable) // &map[key1:1 key2:2 key3:4]
```

### MergeMany
Merges the current hash table with multiple other hash tables and returns the updated hash table.

```Go
myHashtable1 := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable2 := &hashtable.Hashtable[string, int]{"key2": 3, "key3": 4}
mergedHashtable := myHashtable1.MergeMany(myHashtable2, &hashtable.Hashtable[string, int]{"key4": 5})
fmt.Println(mergedHashtable) // &map[key1:1 key2:3 key3:4 key4:5]
```

### MergeManyFunc
Merges the current hash table with multiple other hash tables using a custom function and returns the updated hash table.

```Go
myHashtable1 := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
myHashtable2 := &hashtable.Hashtable[string, int]{"key2": 3, "key3": 4}
mergedHashtable := myHashtable1.MergeManyFunc([]*hashtable.Hashtable[string, int]{myHashtable2}, func(i int, key string, value int) bool {
    return key != "key3"
})
fmt.Println(mergedHashtable) // &map[key1:1 key2:3]
```

### Not
Checks if the given key is not present in the hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
notPresent := myHashtable.Not("key3")
fmt.Println(notPresent) // true
```

### NotMany
Checks if multiple keys are not present in the hash table and returns a slice indicating their absence.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
absentKeys := myHashtable.NotMany("key3", "key4")
fmt.Println(absentKeys) // &[true true]
```

### Pop
Removes the specified key and its associated value from the hash table and returns the value.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
value := myHashtable.Pop("key1")
fmt.Println(value) // 1
fmt.Println(myHashtable) // &map[key2:2]
```

### PopMany
Removes multiple keys and their associated values from the hash table and returns the values in a slice.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2, "key3": 3}
values := myHashtable.PopMany("key1", "key3")
fmt.Println(values) // &[1 3]
fmt.Println(myHashtable) // &map[key2:2]
```

### PopManyFunc
Removes key-value pairs from the hash table using a custom function and returns the values in a slice.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2, "key3": 3}
values := myHashtable.PopManyFunc(func(key string, value int) bool {
    return key != "key2"
})
fmt.Println(values) // &[1 3]
fmt.Println(myHashtable) // &map[key2:2]
```

### PopOK
Removes the specified key and its associated value from the hash table and returns the value along with a boolean indicating existence.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
value, exists := myHashtable.PopOK("key1")
fmt.Println(value, exists) // 1 true
fmt.Println(myHashtable)  // &map[key2:2]
```

### ReplaceMany
Applies the given function to each key-value pair in the hash table and replaces the value if the function returns true. It returns the updated hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2, "key3": 3}
myHashtable.ReplaceMany(func(key string, value int) (int, bool) {
    if key == "key2" {
        return value * 2, true
    }
    return value, false
})
fmt.Println(myHashtable) // &map[key1:1 key2:4 key3:3]
```

### TakeFrom
Empties the current hash table and inserts its content into another hash table. It returns the updated destination hash table.

```Go
source := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
destination := &hashtable.Hashtable[string, int]{"key3": 3}
destination = source.TakeFrom(destination)
fmt.Println(destination) // &map[key1:1 key2:2]
```

### Values
Returns a slice containing all values in the hash table.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2}
values := myHashtable.Values()
fmt.Println(values) // &[1 2]
```

### ValuesFunc
Returns a slice containing values that satisfy the given function.

```Go
myHashtable := &hashtable.Hashtable[string, int]{"key1": 1, "key2": 2, "key3": 3}
values := myHashtable.ValuesFunc(func(key string, value int) bool {
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
	"github.com/lindsaygelle/hashtable" // Import the hashtable package
)

type Animal struct {
	Age     int
	Name    string
	Species string
}

func main() {
	// Create a new Animal
	fluffy := Animal{Age: 1, Name: "Fluffy", Species: "Cat"}

	// Create a new Hashtable of Animals
	animals := &hashtable.Hashtable[string, Animal]{}

	// Add an Animal to the Hashtable using Add function
	animals.Add(fluffy.Name, fluffy)

	// Check if the Hashtable contains a specific key
	if animals.Has("Fluffy") {
		fmt.Println("Fluffy is in the Hashtable!")
	}

	// Get the value associated with a specific key
	if animal, exists := animals.Get("Fluffy"); exists {
		fmt.Println("Found Fluffy in the Hashtable:", animal)
	}

	// Update the age of Fluffy
	updatedFluffy := Animal{Age: 2, Name: "Fluffy", Species: "Cat"}
	animals.Add("Fluffy", updatedFluffy)

	// Delete an entry from the Hashtable
	animals.Delete("Fluffy")

	// Add more animals after deleting Fluffy
	animals.Add("Buddy", Animal{Age: 3, Name: "Buddy", Species: "Dog"})
	animals.Add("Whiskers", Animal{Age: 2, Name: "Whiskers", Species: "Rabbit"})

	// Iterate over the Hashtable and print each key-value pair
	animals.Each(func(key string, value Animal) {
		fmt.Println("Key:", key, "Value:", value)
	})

	// Check if the Hashtable is empty
	if animals.IsEmpty() {
		fmt.Println("Hashtable is empty.")
	} else {
		fmt.Println("Hashtable is not empty.")
	}

	// Get all keys from the Hashtable
	keys := animals.Keys().Slice()
	fmt.Println("Keys in Hashtable:", keys)

	// Get all values from the Hashtable
	values := animals.Values().Slice()
	fmt.Println("Values in Hashtable:", values)
}
```

### Chaining
Chaining methods together.


```Go
package main

import (
	"fmt"
	"github.com/lindsaygelle/hashtable" // Import the hashtable package
)

type Animal struct {
	Age     int
	Name    string
	Species string
}

func main() {
	// Create a new Hashtable of Animals and add animals using method chaining
	animals := &hashtable.Hashtable[string, Animal]{}.
		Add("Fluffy", Animal{Age: 1, Name: "Fluffy", Species: "Cat"}).
		Add("Buddy", Animal{Age: 3, Name: "Buddy", Species: "Dog"}).
		Add("Whiskers", Animal{Age: 2, Name: "Whiskers", Species: "Rabbit"})

	// Print the number of animals in the Hashtable using the Length() method
	fmt.Println("Number of animals in the Hashtable:", animals.Length())

	// Check if a specific animal is in the Hashtable using method chaining
	if exists := animals.Has("Fluffy"); exists {
		fmt.Println("Fluffy is in the Hashtable!")
	}

	// Retrieve and print the age of a specific animal using method chaining
	if age, exists := animals.Get("Buddy"); exists {
		fmt.Println("Buddy's age is:", age.Age)
	}

	// Iterate over the Hashtable and print each key-value pair using method chaining
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
docker build . -t hashtable
```

### Running
Developing and running Go within the Docker container:
```sh
docker run -it --rm --name hashtable hashtable
```

## Docker Compose
A [docker-compose](./docker-compose.yml) file has also been included for convenience:
### Running
Running the compose file.
```sh
docker-compose up -d
```

## Contributing
We warmly welcome contributions to Hashtable. Whether you have innovative ideas, bug reports, or enhancements in mind, please share them with us by submitting GitHub issues or creating pull requests. For substantial contributions, it's a good practice to start a discussion by creating an issue to ensure alignment with the project's goals and direction. Refer to the [CONTRIBUTING](./CONTRIBUTING.md) file for comprehensive details.

## Branching
For a smooth collaboration experience, we have established branch naming conventions and guidelines. Please consult the [BRANCH_NAMING_CONVENTION](./BRANCH_NAMING_CONVENTION.md) document for comprehensive information and best practices.

## License
Hashtable is released under the MIT License, granting you the freedom to use, modify, and distribute the code within this repository in accordance with the terms of the license. For additional information, please review the [LICENSE](./LICENSE) file.

## Security
If you discover a security vulnerability within this project, please consult the [SECURITY](./SECURITY.md) document for information and next steps.

## Code Of Conduct
This project has adopted the [Amazon Open Source Code of Conduct](https://aws.github.io/code-of-conduct). For additional information, please review the [CODE_OF_CONDUCT](./CODE_OF_CONDUCT.md) file.

## Acknowledgements
