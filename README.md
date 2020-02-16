[![Build Status](https://travis-ci.org/gellel/lex.svg?branch=master)](https://travis-ci.org/gellel/map)
[![Apache V2 License](https://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/gellel/map/blob/master/LICENSE)

# Lex

Package lex is a package of map interfaces to handle common map-like operations.

Lex contains a single Lex struct that exposes methods to perform traversal and mutation operations
for a collection of Go interfaces. The Lex struct can be extended to handle
the acceptance and selection of interface specific types. To extend the Lex an interface
can be defined that calls the exposed Lex methods.

Package map comes with all Go primative types as interfaces out of the box.

Each map interface comes with a constructor function that takes zero to n arguments of the
map interface type.

The map interfaces to not expose the underlying interface map to prevent a dirty reference.
This pattern should be adopted when wrapping the Lex struct.

The package is built around the Go API reference documentation. Please consider using `godoc`
to build custom integrations. If you are using Go 1.12 or earlier, godoc should be included. All
Go 1.13 users will need to grab this package using the `go get` flow.

## Installing

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or project's Go module dependencies.

```go get github.com/gellel/map```

To update the SDK use `go get -u` to retrieve the latest version of the SDK.

```go get -u github.com/gellel/map```

## Dependencies

The SDK includes a vendor folder containing the runtime dependencies of the SDK. The metadata of the SDK's dependencies can be found in the Go module file go.mod.

## Go Modules

If you are using Go modules, your go get will default to the latest tagged release version of the SDK. To get a specific release version of the SDK use `@<tag>` in your `go get` command.

```go get github.com/gelle/map@<version>```

To get the latest SDK repository change use @latest.

## License

This SDK is distributed under the Apache License, Version 2.0, see LICENSE for more information.

## Snippets

Lex exports all primative Go types as interfaces. 

```Go
package main

import (
    "fmt"

    "github.com/gellel/map"
)

var (
    b    lex.Byter        // map[interface{}]byte
    c64  lex.Complexer64  // map[interface{}]complex64
    c128 lex.Complexer128 // map[interface{}]complex128
    f32  lex.Floater32    // map[interface{}]float32
    f64  lex.Floater64    // map[interface{}]float64
    i    lex.Inter        // map[interface{}]interface{}
    i8   lex.Inter8       // map[interface{}]int8
    i16  lex.Inter16      // map[interface{}]int16
    i32  lex.Inter32      // map[interface{}]int32
    i64  lex.Inter64      // map[interface{}]int64
    r    lex.Runer        // map[interface{}]rune
    s    *lex.Lex         // map[interface{}]interface{}
    u    lex.UInter       // map[interface{}]uint
    u8   lex.UInter8      // map[interface{}]uint8
    u16  lex.UInter16     // map[interface{}]uint16
    u32  lex.UInter32     // map[interface{}]uint32
    u64  lex.UInter64     // map[interface{}]uint64
    v    lex.Interfacer   // map[interface{}]interface{}
)

func main() {
    var (
        s = lex.NewStringer("a", "b", "c", "go!")
    )
    s.Bounds(0)          // true
    fmt.Println(s.Pop()) // "go!"
}
```

Each interface is intended to handle a unique Go lang primative type.

A Lex interface implements all methods of lex.Lex.

```Go

import (
    "github.com/gellel/lex"
)

func main() {

    var (
        numbers = lex.NewInter(6, 1, 2, 3)
    )
    numbers.Sort().Each(func(i int, n int) {
        fmt.Println(i, n) // (0, 1), (1, 2), (2, 3), (3, 6)
    })
}
```

## Extending

Lex supports interface extension by wrapping the Lex in an struct and exposing a corresponding interface.

This is the pattern implemented by this package and is used for the provided interface types.

```Go
package food 

import (
    "github.com/gellel/map"
)

// Food is a struct that describes food.
type Food struct {
    Name string
}

// Fooder is an interface that contains a collection of Food.
type Fooder interface {
    Append(Food) Fooder
    Prepend(Food) Fooder
}

// fooder is a struct that interfaces with lex.Lex.
type fooder struct { s *lex.Lex }

// Add adds a Food struct to the map.
func (f *fooder) Add(food Food) Fooder {
    f.l.Add(food...)
    return f
}

// Del deletes a Food struct from the map.
func (f *fooder) Del(i interface{}) Fooder { 
    f.l.Del(i)
    return f
}
```