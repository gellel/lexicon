# lexicon
Golang map structure. Extends built-in map data structure. Aims to simplify, map pointer interfacing, offering common object-like commands such as add, delete, peek and more. Requires the slice package for keys and values methods.

Get it:

```
go get github.com/gellel/lexicon
```

Import it:

```
import (
	"github.com/gellel/lexicon"
)
```

## Usage

Creating a basic lexicon pointer.

```go
package main

import (
	"fmt"

	"github.com/gellel/lexicon"
)

func main() {

    lexicon := lexicon.New()

    lexicon.Add("int", 1)

    lexicon.Add("string", "a")

    lexicon.Add("bool", true).Add("other", interface{})
}
```

Creating a lexicon wrapper to accept specific data.

```go
package main

import (
    "github.com/gellel/slice"
)

type T struct{}

type Types struct {
    lexicon *lexicon.Lexicon
}

func (pointer *Types) Add(key string, t T) {
    pointer.lexicon.Add(key, t)
}
```

Using a built-in string lexicon.

```go
package main

import (
    "github.com/gellel/lexicon"
)

func main() {

   lexicon := lexicon.String()

   lexicon.Add("a", "hello").Fetch("a")
}
```

## License

[MIT](https://github.com/gellel/slice/blob/master/LICENSE)
